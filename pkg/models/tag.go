package models

type Tag struct {
	//gorm.Model
	Id        int `gorm:"primary_key" json:"id"`
	Name      string
	Questions []Question `gorm:"many2many:taggings;"`
}
