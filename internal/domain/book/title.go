package book

type Title struct {
	value string
}

func (t Title) Value() string {
	return t.value
}

func NewTitleFromString(value string) Title {
	return Title{value: value}
}
