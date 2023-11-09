package model

import "time"

type Candiate struct {
	ID        int    `gorm:"primary key" json:"candidate_id"`
	Name      string `gorm:"not null" json:"candidate_name"`
	Email     string `gorm:"not null;uniqueIndex" json:"candidate_email"`
	UserName  string `gorm:"not null" json:"candidate_username"`
	Password  string ` json:"candidate_password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Products struct {
	ProductID   int    `gorm:"primary key;"`
	ID          int    `gorm:"foreignKey"`
	ProductName string `gorm:"not null"`
}

type Delivery struct {
	ProductID int `gorm:"foreignKey"`
	ID        int `gorm:"foreignKey"`
	Status    string
}

type Users struct {
	ID        int    `gorm:"primarykey;autoIncreament"`
	Email     string `gorm:"not null;uniqueIndex"`
	Name      string `gorm:"not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
}
