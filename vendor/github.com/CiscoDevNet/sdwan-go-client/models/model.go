package models

//Model is the interface for all the models which we create in the SDWAN
type Model interface {
	ToMap() (map[string]interface{}, error)
}
