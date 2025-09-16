package author_value_objects

type Name struct {
	value string
}

func (n Name) Value() string {
	return n.value
}

func NewNameFromString(value string) Name {
	return Name{value: value}
}
