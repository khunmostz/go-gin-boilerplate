package domain

// Bar represents a bar item in the system
// @Description Bar item information with comprehensive details including name, description, and status
type Bar struct {
	ID          string `json:"id" bson:"_id" gorm:"primaryKey;column:id;type:string" example:"507f1f77bcf86cd799439011" swaggertype:"string" description:"Unique identifier for the bar item"`
	Name        string `json:"name" bson:"name;type:string" gorm:"column:name;type:string" example:"Sample Bar" validate:"required" description:"Name of the bar item (required)"`
	Description string `json:"description" bson:"description;type:string" gorm:"column:description;type:string" example:"This is a sample bar description" description:"Detailed description of the bar item"`
	Status      string `json:"status" bson:"status;type:string" gorm:"column:status;type:string" example:"active" enums:"active,inactive" description:"Current status of the bar item"`
}
