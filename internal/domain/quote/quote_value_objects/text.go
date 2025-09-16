package quote_value_objects

type Text struct {
	value string
}

func (vo Text) Value() string {
	return vo.value
}

func NewTextFromString(value string) Text {
	return Text{value: value}
}
