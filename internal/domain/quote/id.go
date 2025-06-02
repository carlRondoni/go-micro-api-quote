package quote

import "github.com/google/uuid"

type Id struct {
	value uuid.UUID
}

func (vo Id) Value() uuid.UUID {
	return vo.value
}

func (vo Id) String() string {
	return vo.value.String()
}

func NewIdFromUuid(value uuid.UUID) Id {
	return Id{value: value}
}

func NewIdFromString(value string) (Id, error) {
	uuid, err := uuid.Parse(value)
	if err != nil {
		return Id{}, err
	}

	return Id{value: uuid}, nil
}
