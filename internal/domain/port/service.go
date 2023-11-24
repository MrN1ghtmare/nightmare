package port

import "nightmare/internal/domain/entity"

type EntityService interface {
	Create(e *entity.Entity) error
}

type ServiceService interface {
	Create(e *entity.Service) error
}
