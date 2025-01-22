package formatter

type ModelContent struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewModelContent(id string, name string, modelType string) ModelContent {
	return ModelContent{
		Id: id, Name: name, Type: modelType,
	}
}
