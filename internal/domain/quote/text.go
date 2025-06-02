package quote

type Text struct {
	value string
}

func (vo Text) Value() string {
	return vo.value
}

func NewtextFromString(value string) Text {
	return Text{value: value}
}
