package model

type Category struct {
	ID           string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	ProductID    string `gorm:"type:uuid;not null;references:ID" json:"productID" binding:"required"`
	CategoryName string `gorm:"type:varchar(255);not null" json:"categoryname" binding:"required,alphanum"`
}
