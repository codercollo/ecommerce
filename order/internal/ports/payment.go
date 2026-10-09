package ports

import "github.com/codercollo/ecommerce/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
