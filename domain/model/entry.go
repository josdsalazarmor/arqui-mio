package model

import (

	"github.com/jinzhu/gorm"
)

type Entry struct {
	gorm.Model
	ID        int        	`gorm:"primary_key" json:"id"`
	UserId        int       `gorm:"not null" json:"userId"  binding:"required"`
	Title     string     	`gorm:"not null" json:"title" binding:"required"`
	Link     string     	`gorm:"not null" json:"link" binding:"required"`
	ChatID     string    	`gorm:"not null" json:"chatID" binding:"required"`

	// When you get an Entry Model with gorm, you can get associated tags using junction table automatically.
	// When you write or update an Entry Model, you have to include　tags id. (API Requests don't have tags id.)
	Tags []*Tag `gorm:"many2many:entries_tags;association_autoupdate:false;association_autocreate:false;" json:"tags"`
}
