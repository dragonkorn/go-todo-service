package database

import (
	"service/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DB ...
type DB struct {
	Connection        *gorm.DB
	ShareDBConnection *gorm.DB
	config            *config.Configuration
}

func (b *DB) connectDatabase(host string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(host), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

//NewServerDB ...
func NewServerDB(c *config.Configuration) *DB {
	g := &DB{
		config: c,
	}

	g.Connection = g.connectDatabase(c.DbConnection)

	if c.IsDbDebug {
		g.Connection = g.Connection.Debug()
	}

	return g
}
