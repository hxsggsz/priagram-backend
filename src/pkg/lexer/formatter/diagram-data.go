package formatter

type DiagramData struct {
	ModelName    string `json:"modelName"`
	DiagramType  string `json:"diagramType"`
	ModelContent []ModelContent
}

func NewDiagramData(modelName string, diagramType string, modelContents []ModelContent) DiagramData {
	return DiagramData{
		ModelName: modelName, DiagramType: diagramType, ModelContent: modelContents,
	}
}
