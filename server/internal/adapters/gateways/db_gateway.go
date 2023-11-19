package gateways

import "gorm.io/gorm"

type DB interface {
	Connect() *gorm.DB
	Begin() *gorm.DB
}

type DBGateway struct {
	DB DB
}

func NewDBGateway() DB {
	return &DBGateway{}
}

func (g *DBGateway) Connect() *gorm.DB {
	return g.DB.Connect()
}

func (g *DBGateway) Begin() *gorm.DB {
	return g.DB.Begin()
}
