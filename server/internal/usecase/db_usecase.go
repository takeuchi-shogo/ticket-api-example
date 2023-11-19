package usecase

import "gorm.io/gorm"

type DBUsecase interface {
	Connect() *gorm.DB
	Begin() *gorm.DB
}
