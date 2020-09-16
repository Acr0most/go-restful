package connector

type ConnectorInterface interface {
	Find(params map[string]interface{}, result interface{}) (err error)
	Create(items interface{}) (err error)
	Delete(items interface{}, result interface{}) (err error)
	Patch(items interface{}, result interface{}, model interface{}) (err error)
}
