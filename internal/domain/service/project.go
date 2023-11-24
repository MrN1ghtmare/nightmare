package service

import (
	"nightmare/internal/domain/entity"
	"os"
	"os/exec"
)

type ProjectService struct {
	*sourceFileService
}

func (ps *ProjectService) createDefaultDirs() error {
	const defaultPerm = 0755
	defaultDirs := []string{
		"build",
		"cmd",
		"configs",
		"deploy",
		"docs",
		"examples",
		"internal",
		"pkg",
		"test",
		"internal/controller",
		"internal/domain",
		"internal/domain/entity",
		"internal/domain/port",
		"internal/domain/service",
	}

	for _, dir := range defaultDirs {
		if err := os.MkdirAll(dir, defaultPerm); err != nil {
			return err
		}
	}

	return nil
}

func (ps *ProjectService) initGoProject(packageName *string) error {
	return exec.Command("go", "mod", "init", *packageName).Run()
}

func (ps *ProjectService) createNightmareConfigFile() error {
	_, err := os.Create("./configs/nightmare.yml")
	return err
}

func (ps *ProjectService) Create(p *entity.Project) error {
	if err := ps.initGoProject(p.Name); err != nil {
		return err
	}

	if err := ps.createDefaultDirs(); err != nil {
		return err
	}

	if err := ps.createNightmareConfigFile(); err != nil {
		return err
	}

	return nil
}

func NewProjectService() *ProjectService {
	return &ProjectService{newSourceFileService("service")}
}
