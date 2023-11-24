package service

import (
	"fmt"
	"nightmare/internal/domain/entity"
)

type EntityService struct {
	*sourceFileService
}

func (es *EntityService) Create(e *entity.Entity) error {
	if err := es.validateName(e.Name); err != nil {
		return err
	}

	e.FileName = es.genFileName(e.Name)

	targetDir := "./internal/domain/entity"
	content := fmt.Sprintf(
		"package entity\n\ntype %s struct {\n}\n\nfunc New%s() %s {\n\treturn %s{}\n}\n",
		*e.Name,
		*e.Name,
		*e.Name,
		*e.Name,
	)

	if err := es.createSourceFile(&targetDir, e.FileName, &content); err != nil {
		return err
	}

	return nil
}

func NewEntityService() *EntityService {
	return &EntityService{newSourceFileService("entity")}
}
