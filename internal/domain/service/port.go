package service

import (
	"fmt"
	"nightmare/internal/domain/entity"
)

type PortService struct {
	*sourceFileService
}

func (ps *PortService) Create(p *entity.Port) error {
	if err := ps.validateName(p.Name); err != nil {
		return err
	}

	p.FileName = ps.genFileName(p.Name)

	targetDir := "./internal/domain/port"
	content := fmt.Sprintf(
		"package port\n\ntype %s struct {\n}\n\nfunc New%s() *%s {\n\treturn %s{}\n}\n",
		*p.Name,
		*p.Name,
		*p.Name,
		*p.Name,
	)

	if err := ps.createSourceFile(&targetDir, p.FileName, &content); err != nil {
		return err
	}

	return nil
}

func NewPortService() *PortService {
	return &PortService{newSourceFileService("port")}
}
