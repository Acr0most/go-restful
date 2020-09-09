package connector

type ConnectorInterface interface {
	Find(params map[string]interface{}, result interface{}) (success bool)
	Create(items interface{})
	Delete(items interface{}, result interface{})
	Patch(items interface{}, result interface{}, model interface{})
}
