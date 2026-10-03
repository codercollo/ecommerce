package api

import (
	"github.com/codercollo/ecommerce/order/internal/application/domain"
	"github.com/codercollo/ecommerce/order/internal/ports"
)

// Application handles core order business logic and port interactions.
type Application struct {
	db ports.DBPort
}

// NewApplication constructs a new Application instance with required dependencies.
func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

// PlaceOrder executes the order creation and persistence workflow.
func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	// Persist the order via the database port
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
