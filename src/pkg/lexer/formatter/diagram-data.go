package formatter

import "strings"

type DiagramData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Data Data   `json:"data"`
}

type Data struct {
	ModelName    string         `json:"modelName"`
	ModelContent []ModelContent `json:"modelContent"`
}

type Diagram struct {
	Data      []DiagramData `json:"data"`
	Relations []Relation    `json:"relations"`
}

func newData(modelName string, modelContent []ModelContent) Data {
	return Data{
		ModelName: modelName, ModelContent: modelContent,
	}
}

func NewDiagramData(modelName string, diagramType string, modelContents []ModelContent) DiagramData {

	data := newData(strings.ToLower(modelName), modelContents)

	return DiagramData{
		Id: strings.ToLower(modelName), Type: diagramType, Data: data,
	}
}

func NewDiagram(data []DiagramData, relations []Relation) Diagram {
	return Diagram{Data: data, Relations: relations}
}
