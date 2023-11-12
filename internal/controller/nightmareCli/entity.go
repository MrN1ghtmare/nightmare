package nightmareCli

import (
	"fmt"
	"log"
	"nightmare/internal/domain/entity"
	"nightmare/internal/domain/port"
	"nightmare/internal/domain/service"

	"github.com/spf13/cobra"
)

var Entity = NewEntityController(service.NewEntityService())

type EntityController struct {
	srvc port.EntityService
}

func (ec *EntityController) Create(cmd *cobra.Command, args []string) {
	e := entity.NewEntity(&args[0])

	if err := ec.srvc.Create(e); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("entity %s created!\n", *e.Name)
}

func NewEntityController(srvc port.EntityService) *EntityController {
	return &EntityController{srvc}
}
