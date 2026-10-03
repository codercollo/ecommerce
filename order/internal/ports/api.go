package ports

import "github.com/codercollo/ecommerce/order/internal/application/domain"

// APIPort defines the inbound application business logic contract.
type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
