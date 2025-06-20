package domain

// Foo represents a foo item in the system
// @Description Foo item information with unique identifier and name
type Foo struct {
	ID   string `json:"id" bson:"_id" gorm:"primaryKey;column:user_id;type:string" example:"507f1f77bcf86cd799439011" swaggertype:"string" description:"Unique identifier for the foo item"`
	Name string `json:"name" bson:"name;type:string" gorm:"column:name;type:string" example:"Sample Foo" validate:"required" description:"Name of the foo item (required)"`
}
