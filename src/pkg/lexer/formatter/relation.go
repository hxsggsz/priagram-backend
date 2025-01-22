package formatter

type Relation struct {
	Id     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}

func NewRelation(id string, source string, target string) Relation {
	return Relation{
		Id: id, Source: source, Target: target,
	}
}
