package api

import (
	"github.com/codercollo/ecommerce/order/internal/application/core/domain"
	"github.com/codercollo/ecommerce/order/internal/ports"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Application handles core order business logic and port interactions.
type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

// NewApplication constructs a new Application instance with required dependencies.
func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

// PlaceOrder executes the order creation and persistence workflow.
func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	// Persist the order via the database port
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}

	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		st, _ := status.FromError(paymentErr)
		fieldErr := &errdetails.BadRequest_FieldViolation{
			Field:       "payment",
			Description: st.Message(),
		}

		badReq := &errdetails.BadRequest{}
		badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
		orderStatus := status.New(codes.InvalidArgument, "order creation failed")

		statusWithDetails, _ := orderStatus.WithDetails(badReq)
		return domain.Order{}, statusWithDetails.Err()
	}
	return order, nil
}
