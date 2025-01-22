package formatter

type DiagramData struct {
	ModelName    string         `json:"modelName"`
	DiagramType  string         `json:"diagramType"`
	ModelContent []ModelContent `json:"modelContent"`
	Relations    []Relation     `json:"relations"`
}

func NewDiagramData(modelName string, diagramType string, modelContents []ModelContent, relations []Relation) DiagramData {
	return DiagramData{
		ModelName: modelName, DiagramType: diagramType, ModelContent: modelContents, Relations: relations,
	}
}
