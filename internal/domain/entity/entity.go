package entity

type Entity struct {
	Name     *string
	FileName *string
}

func NewEntity(name *string) *Entity {
	return &Entity{Name: name}
}
