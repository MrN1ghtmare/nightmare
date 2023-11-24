package nightmareCli

import (
	"fmt"
	"log"
	"nightmare/internal/domain/entity"
	"nightmare/internal/domain/port"
	"nightmare/internal/domain/service"

	"github.com/spf13/cobra"
)

var Project = NewProjectController(service.NewProjectService(), service.NewPortService())

type ProjectController struct {
	projectSrvc port.ProjectService
	portSrvc    port.PortService
}

func (pc *ProjectController) createDefaultPorts() error {
	defaultPorts := []string{"Service", "Repository", "Controller"}

	for _, defaultPort := range defaultPorts {
		p := entity.NewPort(&defaultPort)
		if err := pc.portSrvc.Create(p); err != nil {
			return err
		}
	}

	return nil
}

func (pc *ProjectController) Create(cmd *cobra.Command, args []string) {
	p := entity.NewProject(&args[0])

	if err := pc.projectSrvc.Create(p); err != nil {
		log.Fatal(err)
	}

	if err := pc.createDefaultPorts(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("project %s created!\n", *p.Name)
}

func NewProjectController(projectSrvc port.ProjectService, portSrvc port.PortService) *ProjectController {
	return &ProjectController{projectSrvc: projectSrvc, portSrvc: portSrvc}
}
