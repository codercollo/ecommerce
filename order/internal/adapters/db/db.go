package db

import (
	"fmt"

	"github.com/codercollo/ecommerce/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Order represents the GORM database table schema for orders.
type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
}

// OrderItem represents the GORM database table schema for individual order line items.
type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

// Adapter satisfies ports.DBPort by bridging domain entities with GORM/MySQL storage.
type Adapter struct {
	db *gorm.DB
}

// NewAdapter initializes the MySQL database connection and executes automatic schema migrations.
func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	// Open connection to MySQL database.
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	// Automigrate schema tables
	err := db.AutoMigrate(&Order{}, OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

// Get fetches an order from the database by ID and converts it to a domain.Order entity.
func (a Adapter) Get(id string) (domain.Order, error) {
	var orderEntity Order

	// Preload nested OrderItems relation to ensure associated items are populated
	res := a.db.First(&orderEntity, id)

	// Map database order items to domain order items
	var orderItems []domain.OrderItem
	for _, orderItem := range orderEntity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	// Map database GORM entity to core domain entity
	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt.UnixNano(),
	}

	return order, res.Error
}

// Save persists a domain.Order entity into MySQL and assigns the generated primary ID back to the domain struct.
func (a Adapter) Save(order *domain.Order) error {
	var orderItems []OrderItem

	// Map domain items to GORM database entities.
	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	// Build GBRM database model
	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}

	// Insert record and associated items in a single database transaction
	res := a.db.Create(&orderModel)
	if res.Error == nil {
		order.ID = int64(orderModel.ID)
	}

	return res.Error
}
