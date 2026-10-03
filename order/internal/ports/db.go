package ports

import "github.com/codercollo/ecommerce/order/internal/application/domain"

// DBPort defines the outbound database storage contract.
type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
