package quote

type Author struct {
	value string
}

func (vo Author) Value() string {
	return vo.value
}

func NewAuthorFromString(value string) Author {
	return Author{value: value}
}
