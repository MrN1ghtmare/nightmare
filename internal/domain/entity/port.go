package entity

type Port struct {
	Name     *string
	FileName *string
}

func NewPort(name *string) *Port {
	return &Port{Name: name}
}
