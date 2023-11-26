package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/takeuchi-shogo/ticket-api/config"
	"github.com/takeuchi-shogo/ticket-api/internal/adapters/gateways"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type DB struct {
	Connection *bun.DB
}

func NewDB(c config.ServerConfig) gateways.DB {
	return &DB{
		Connection: connection(
			c.DBHost,
			c.DBUserName,
			c.DBPassword,
			c.DBName,
		),
	}
}

func connection(host, username, password, dbName string) *bun.DB {
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold:             time.Nanosecond,
	// 		LogLevel:                  logger.Info,
	// 		IgnoreRecordNotFoundError: true,
	// 		Colorful:                  true,
	// 	},
	// )

	count := 0

	// conn, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbName)), &gorm.Config{
	// 	Logger: newLogger,
	// })
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbName)
	sqldb, err := dbOpen(dsn)

	if err != nil {
		for {
			if err == nil {
				break
			}
			fmt.Print(".\n")
			time.Sleep(time.Second)
			count++
			// connection wait 10 seconds for database starting...
			if count > 10 {
				fmt.Print("database connection failed\n")
				panic(err.Error())
			}
			sqldb, err = dbOpen(dsn)
			// conn, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbName)), &gorm.Config{})
		}
	}

	// sqldb.SetMaxIdleConns(config.MaxIdleConns)
	// sqldb.SetMaxOpenConns(config.MaxOpenConns)

	db := bun.NewDB(sqldb, mysqldialect.New())

	log.Print("database connection success\n")

	return db
}

func dbOpen(dsn string) (*sql.DB, error) {
	return sql.Open("mysql", dsn)
}

func (db *DB) Connect() *bun.DB {
	return db.Connection
}

// Start a Transaction
func (db *DB) Transaction() (bun.Tx, error) {
	ctx := context.Background()
	return db.Connection.BeginTx(ctx, &sql.TxOptions{})
}
