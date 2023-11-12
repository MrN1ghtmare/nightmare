package service

import (
	"fmt"
	"nightmare/internal/domain/entity"
)

type ServiceService struct {
	*sourceFileService
}

func (ss *ServiceService) Create(s *entity.Service) error {
	if err := ss.validateName(s.Name); err != nil {
		return err
	}

	ss.generateElementFileName(s.Name)

	targetDir := "./internal/domain/service"
	content := fmt.Sprintf(
		"package service\n\ntype %s struct {\n}\n\nfunc New%s() %s {\n\treturn %s{}\n}\n",
		*s.Name,
		*s.Name,
		*s.Name,
		*s.Name,
	)

	if err := ss.createSourceFile(&targetDir, s.FileName, &content); err != nil {
		return err
	}

	return nil
}

func NewServiceService() *ServiceService {
	return &ServiceService{newSourceFileService("service")}
}
