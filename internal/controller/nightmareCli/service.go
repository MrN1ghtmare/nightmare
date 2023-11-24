package nightmareCli

import (
	"fmt"
	"log"
	"nightmare/internal/domain/entity"
	"nightmare/internal/domain/port"
	"nightmare/internal/domain/service"

	"github.com/spf13/cobra"
)

var Service = NewServiceController(service.NewServiceService())

type ServiceController struct {
	srvc port.ServiceService
}

func (sc *ServiceController) Create(cmd *cobra.Command, args []string) {
	s := entity.NewService(&args[0])

	if err := sc.srvc.Create(s); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("service %s created!\n", *s.Name)
}

func NewServiceController(srvc port.ServiceService) *ServiceController {
	return &ServiceController{srvc}
}
