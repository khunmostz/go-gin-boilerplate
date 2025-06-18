package domain

type Bar struct {
	ID          string `json:"id" bson:"_id" gorm:"primaryKey;column:bar_id;type:string"`
	Name        string `json:"name" bson:"name;type:string" gorm:"column:name;type:string"`
	Description string `json:"description" bson:"description;type:string" gorm:"column:description;type:string"`
	Status      string `json:"status" bson:"status;type:string" gorm:"column:status;type:string"`
}
