package database

import (
	"fmt"
	"go-cin/model"
	"go-cin/pkg/config"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgDb struct {
	config config.Config
}

func NewPgDb(config config.Config) (*PgDb, error) {
	return &PgDb{config: config}, nil
}

func (d *PgDb) ConnectDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", d.config.Database.DbHost, d.config.Database.DbPort, d.config.Database.DbUser, d.config.Database.DbPass,
		d.config.Database.DbName)
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  dsn,
				PreferSimpleProtocol: true,
			},
		))
	if err != nil {
		return nil, fmt.Errorf("gorm open error: %w", err)
	}
	err = AutoMigrate(db)
	if err != nil {
		return nil, fmt.Errorf("db automigrate error: %w", err)
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Album{},
	)
}

func (d *PgDb) Ping() error {
	db, err := d.ConnectDb()
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
