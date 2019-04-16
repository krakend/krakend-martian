package collection

import (
	"strings"
)

type Move map[string]interface{}

func (m *Move) Execute(body []map[string]interface{}) {
	if nil != m {
		for _, item := range body {
			for originalItemPath, newItemPath := range *m {
				moveItem(originalItemPath, newItemPath.(string), item)
			}
		}
	}
}

func moveItem(originalItemPath, newItemPath string, item interface{}) {
	if originalItemPath == "" || newItemPath == "" {
		return
	}
	originalItem, itemExist := findItemRecursive(strings.Split(originalItemPath, "."), item)
	if !itemExist {
		return
	}
	createItemRecursive(strings.Split(newItemPath, "."), originalItem, item)
}

func findItemRecursive(path []string, item interface{}) (interface{}, bool) {
	var foundedItem interface{}
	var itemExist bool

	if len(path) == 0 {
		foundedItem = item
		itemExist = true
	} else {
		pathKey := path[0]
		switch item.(type) {
		case map[string]interface{}:
			item := item.(map[string]interface{})
			if _, ok := item[pathKey]; ok {
				i := item[pathKey]
				foundedItem, itemExist = findItemRecursive(path[1:], i)
				if len(path[1:]) == 0 && itemExist {
					delete(item, pathKey)
				}
			} else {
				itemExist = false
			}
		default:
			return nil, false
		}
	}
	return foundedItem, itemExist
}

func createItemRecursive(path []string, originalItem interface{}, item interface{}) {
	if len(path) == 0 {
		return
	}

	pathKey := path[0]
	if len(path) == 1 {
		item.(map[string]interface{})[pathKey] = originalItem
	} else {
		if _, ok := item.(map[string]interface{})[pathKey]; ok {
			createItemRecursive(path[1:], originalItem, item.(map[string]interface{})[pathKey])
		} else {
			if len(path) > 1 {
				item.(map[string]interface{})[pathKey] = map[string]interface{}{}
				createItemRecursive(path[1:], originalItem, item.(map[string]interface{})[pathKey])
			}
		}
	}
}
