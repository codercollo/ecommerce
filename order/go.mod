module github.com/codercollo/ecommerce/order

go 1.25.5

require (
	github.com/codercollo/ecommerce/golang/order v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.84.0
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.31.2
)

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/go-sql-driver/mysql v1.10.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.57.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.40.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260706201446-f0a921348800 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
)

replace github.com/codercollo/ecommerce/golang/order => ../golang/order

replace github.com/codercollo/ecommerce/golang/payment => ../golang/payment

replace github.com/codercollo/ecommerce/golang/shipping => ../golang/shipping
