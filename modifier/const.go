package modifier

const (
	ContentHeader              = "Content-Type"
	HeaderApplicationJSONValue = "application/json"

	//TODO вынести из этого пакета зависимость от логера..
	//Logger
	CollectionModifierWarning = "collectionModifierFromJSON.createModifier: %s, path: %s"
	CollectionModifierDebug   = "create \"%s\" modifier for path \"%s\": OK"

	QueryParamsModifierWarning = "queryParamsModifierFromJSON.createModifier: %s, type: %s"
	QueryParamsModifierDebug   = "create queryParams \"%s\" modifier: OK"

	PostParamsModifierWarning = "postParamsModifierFromJSON.createModifier: %s, type: %s"
)
