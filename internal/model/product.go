package model

type Product struct {
	ID       string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	Name     string `gorm:"type:varchar(255);not null;unique" json:"name" binding:"required,alphanum"`
	Quantity int    `gorm:"type:int;not null;" json:"quantity" binding:"required"`
	Price    int    `gorm:"type:int;not null;" json:"price" binding:"required"`
}
