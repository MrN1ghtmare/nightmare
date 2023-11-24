package nightmareCli

import (
	"fmt"
	"log"
	"nightmare/internal/domain/entity"
	"nightmare/internal/domain/port"
	"nightmare/internal/domain/service"

	"github.com/spf13/cobra"
)

var Port = NewPortController(service.NewPortService())

type PortController struct {
	srvc port.PortService
}

func (pc *PortController) Create(cmd *cobra.Command, args []string) {
	p := entity.NewPort(&args[0])

	if err := pc.srvc.Create(p); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("port %s created!\n", *p.Name)
}

func NewPortController(srvc port.PortService) *PortController {
	return &PortController{srvc}
}
