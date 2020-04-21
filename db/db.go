package db

import (
	"database/sql"

	"github.com/huandu/go-sqlbuilder"

	"github.com/242617/synapse-core/config"
	"github.com/242617/synapse-core/db/model"
	"github.com/242617/synapse-core/log"
)

type DB interface {
	model.Queries
	Connect(string) error
	Close() error
}

type db struct {
	*model.Managers

	log        log.Logger
	cfg        config.DBConfig
	connection *sql.DB
}

func New(config config.DBConfig, logger log.Logger) DB {
	return &db{
		cfg: config,
		log: logger,
	}
}

func (db *db) Connect(connection string) error {
	if db.connection != nil {
		db.connection.Close()
	}

	sqlDB, err := sql.Open("postgres", connection)
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(150)
	sqlbuilder.DefaultFlavor = sqlbuilder.PostgreSQL

	if err := sqlDB.Ping(); err != nil {
		return err
	}

	db.connection = sqlDB
	db.Managers = model.NewManager(db.connection)

	return nil
}

func (db *db) Close() error {
	return db.connection.Close()
}
