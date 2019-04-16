package modifier

/**
Чтобы добавить новый модификатор - надо прописать его в krakend.json,
в отдельном файле (см. black_list.go для примера) сделать структуру модификатора
и реализовать метод Execute (здесь логика работы модификатора) и добавить создание в фабрику modifier.factory.go
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/devopsfaith/krakend/logging"
	"github.com/google/martian/parse"
	"io/ioutil"
	"github.com/devopsfaith/krakend-martian/modifier/collection"
	"net/http"
	"strings"
)

func init() {
	parse.Register("collection.modifier", collectionModifierFromJSON)
}

// обязательно слайс, чтобы выполнялся по порядку
type modificationList struct {
	modifiers []modifiersByPath
	Logger    logging.Logger
}

type modifiersByPath struct {
	Path       string
	Collection []collection.ModifierInterface
}

type collectionModifierJSON struct {
	Scope            []parse.ModifierType `json:"scope"`
	ModificationList []modificationJSON   `json:"modifications"`
}

type modificationJSON struct {
	ModifierList []modifierJSON
	Path         string
}

type modifierJSON struct {
	Type string
	Data interface{}
}

func (m *modificationList) SetLogger(l logging.Logger) {
	m.Logger = l

	for _, cm := range m.modifiers {
		for _, mm := range cm.Collection {
			if v, ok := mm.(LoggerAdder); ok {
				v.SetLogger(m.Logger)
			}
		}
	}
}

func (m *modificationList) addModifier(path string, modifier collection.ModifierInterface) {
	for key, mbp := range m.modifiers {
		if mbp.Path == path {
			m.modifiers[key].Collection = append(m.modifiers[key].Collection, modifier)
			return
		}
	}

	modifiersByPath := modifiersByPath{}
	modifiersByPath.Path = path
	modifiersByPath.Collection = append(modifiersByPath.Collection, modifier)
	m.modifiers = append(m.modifiers, modifiersByPath)
}

func (m modificationList) ModifyRequest(req *http.Request) error {
	var (
		body map[string]interface{}
	)

	rh := req.Header.Get(ContentHeader)
	rh = strings.ToLower(rh)
	rh = strings.Split(rh, ";")[0]

	if rh == HeaderApplicationJSONValue {
		bodyBytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			m.Logger.Error(err)
			return err
		}

		err = json.Unmarshal(bodyBytes, &body)
		if err == nil && body != nil {
			modifyBody(m, body)

			if err = req.Body.Close(); err != nil {
				m.Logger.Error(err)
			}
			js, err := json.Marshal(body)
			if err != nil {
				m.Logger.Error(err)
				return err
			}
			req.Body = ioutil.NopCloser(bytes.NewReader(js))
		}
	}
	return nil
}

func (m modificationList) ModifyResponse(res *http.Response) error {
	var (
		body map[string]interface{}
	)

	rh := res.Header.Get(ContentHeader)
	rh = strings.ToLower(rh)

	if rh == HeaderApplicationJSONValue {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			m.Logger.Error(err)
			return err
		}

		err = json.Unmarshal(bodyBytes, &body)
		if err == nil && body != nil {
			modifyBody(m, body)

			if err := res.Body.Close(); err != nil {
				m.Logger.Error(err)
			}
			js, err := json.Marshal(body)
			if err != nil {
				m.Logger.Error(err)
				return err
			}
			res.Body = ioutil.NopCloser(bytes.NewReader(js))
		}
	}
	return nil
}

func modifyBody(ml modificationList, body map[string]interface{}) {
	for _, modifiersByPath := range ml.modifiers {
		itemCollection := getCollectionByPath(modifiersByPath.Path, body)
		for _, modifier := range modifiersByPath.Collection {
			modifier.Execute(itemCollection)
		}
	}
}

func getCollectionByPath(path string, body map[string]interface{}) []map[string]interface{} {
	var itemCollection []map[string]interface{}

	if path == "" {
		itemCollection = append(itemCollection, body)
	} else {
		for _, item := range getRecursiveCollection(strings.Split(path, "."), body) {
			itemCollection = append(itemCollection, item.(map[string]interface{}))
		}
	}

	return itemCollection
}

func getRecursiveCollection(path []string, item interface{}) []interface{} {
	var itemCollection []interface{}

	if len(path) == 0 {
		switch item.(type) {
		case []interface{}:
			itemCollection = append(itemCollection, item.([]interface{})...)
		case map[string]interface{}:
			itemCollection = append(itemCollection, item)
		}
	} else {
		pathKey := path[0]
		switch item.(type) {
		case []interface{}:
			for _, sliceItem := range item.([]interface{}) {
				itemCollection = append(itemCollection, getRecursiveCollection(path, sliceItem)...)
			}
		case map[string]interface{}:
			item := item.(map[string]interface{})
			if _, ok := item[pathKey]; ok {
				i := item[pathKey]
				itemCollection = getRecursiveCollection(path[1:], i)
			}
		}
	}
	return itemCollection
}

func collectionModifierFromJSON(b []byte) (*parse.Result, error) {
	msg := &collectionModifierJSON{}
	ml := &modificationList{}

	if err := json.Unmarshal(b, msg); err != nil {
		return nil, err
	}

	for _, mod := range msg.ModificationList {
		for _, modifierConfig := range mod.ModifierList {
			modifier, err := collection.CreateModifier(modifierConfig.Type, modifierConfig.Data)
			if err != nil {
				return nil, fmt.Errorf(CollectionModifierWarning, err, mod.Path)
			}

			if v, ok := modifier.(LoggerAdder); ok {
				v.SetLogger(ml.Logger)
			}
			ml.addModifier(mod.Path, modifier)
		}
	}

	r, err := parse.NewResult(ml, msg.Scope)
	if err != nil {
		return nil, err
	}

	return r, nil
}
