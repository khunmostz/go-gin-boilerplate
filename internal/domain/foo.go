package domain

// Foo represents a foo item
// @Description Foo item information
type Foo struct {
	ID   string `json:"id" bson:"_id" gorm:"primaryKey;column:user_id;type:string" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name string `json:"name" bson:"name;type:string" gorm:"column:name;type:string" example:"Sample Foo"`
}
