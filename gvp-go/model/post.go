package model

import (
	"github.com/satori/uuid"
	"gorm.io/gorm"
	"time"
)

type Post struct{
	ID uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserId uint `json:"user_id" gorm:"not null"`
	CategoryId uint `json:"category_id" gorm:"not null"`
	Category *Category
	Title string `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg string `json:"head_img"`
	Content string `json:"content" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
}

func (base *Post) BeforeCreate(tx *gorm.DB) error {
    uuid := uuid.NewV4().String()
    tx.Statement.SetColumn("ID", uuid)
    return nil
}