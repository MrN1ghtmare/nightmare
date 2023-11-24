package service

import (
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
	content := "package port\n"

	if err := ps.createSourceFile(&targetDir, p.FileName, &content); err != nil {
		return err
	}

	return nil
}

func NewPortService() *PortService {
	return &PortService{newSourceFileService("port")}
}
