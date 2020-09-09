package connector

type ConnectorInterface interface {
	Find(params map[string]interface{}, result interface{}) (success bool)
	Create(items interface{})
	Delete(items interface{}, result interface{})
	// Update(params map[string]interface{}, result interface{})
	// UpdateOne() interface{}
}
