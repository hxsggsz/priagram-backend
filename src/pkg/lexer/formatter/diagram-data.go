package formatter

import (
	"priagram/src/pkg/id"
)

type DiagramData struct {
	Id        string     `json:"id"`
	Type      string     `json:"type"`
	Position  Position   `json:"position"`
	Data      Data       `json:"data"`
	Relations []Relation `json:"relations"`
}

type Data struct {
	ModelName    string         `json:"modelName"`
	ModelContent []ModelContent `json:"modelContent"`
}

func newData(modelName string, modelContent []ModelContent) Data {
	return Data{
		ModelName: modelName, ModelContent: modelContent,
	}
}

func NewDiagramData(modelName string, diagramType string, position Position, modelContents []ModelContent, relations []Relation) DiagramData {
	diagramId := id.FmtId(modelName, diagramType)

	data := newData(modelName, modelContents)

	return DiagramData{
		Id: diagramId, Type: diagramType, Position: position, Data: data, Relations: relations,
	}
}
