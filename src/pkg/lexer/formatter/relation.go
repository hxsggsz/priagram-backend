package formatter

import "fmt"

type Relation struct {
	Id           string `json:"id"`
	Source       string `json:"source"`
	Target       string `json:"target"`
	SourceHandle string `json:"sourceHandle"`
	TargetHandle string `json:"targetHandle"`
	Animated     bool   `json:"animated"`
}

func NewRelation(id string, source string, target string) Relation {
	sourceHandle := fmt.Sprintf("%s-source", source)
	targetHandle := fmt.Sprintf("%s-target", target)
	return Relation{
		Id: id, Source: source, SourceHandle: sourceHandle, Target: target, TargetHandle: targetHandle, Animated: true,
	}
}
