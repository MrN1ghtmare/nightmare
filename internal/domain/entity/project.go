package entity

type Project struct {
	Name *string
}

func NewProject(name *string) *Project {
	return &Project{Name: name}
}
