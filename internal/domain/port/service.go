package port

import "nightmare/internal/domain/entity"

type EntityService interface {
	Create(e *entity.Entity) error
}
