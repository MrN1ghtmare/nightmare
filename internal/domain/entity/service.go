package entity

type Service struct {
	Name     *string
	FileName *string
}

func NewService(name *string) *Service {
	return &Service{Name: name}
}
