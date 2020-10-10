package v0

// User Model struct
type User struct {
	ID   string
	Name string
}

// Product Model struct
type Product struct {
	ID     string
	Price  uint
	UserID uint
	User   User
}

// ModelSchemaList v0 Model Structs
var ModelSchemaList = []interface{}{User{}, Product{}}
