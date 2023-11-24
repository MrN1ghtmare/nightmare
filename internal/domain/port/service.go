package port

import "nightmare/internal/domain/entity"

type EntityService interface {
	Create(e *entity.Entity) error
}

type ServiceService interface {
	Create(e *entity.Service) error
}

type PortService interface {
	Create(e *entity.Port) error
}

type ProjectService interface {
	Create(p *entity.Project) error
}
