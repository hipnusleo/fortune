package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// BaseModel is a base model
type BaseModel struct {
	ID         int
	CreatedOn  int
	ModifiedOn int
	DeletedOn  int
}

// Setup database(mysql)
func Setup() {

}
