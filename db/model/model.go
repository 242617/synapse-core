package model

import (
	"database/sql"

	"github.com/huandu/go-sqlbuilder"
)

type Queries interface {
	CrawlerManager
}

type Managers struct {
	*crawlerManager
}

type manager struct {
	core       *Managers
	connection SQLConnection

	tableName   string
	tableStruct *sqlbuilder.Struct
}

type SQLConnection interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func NewManager(conn SQLConnection) *Managers {
	core := new(Managers)
	core.crawlerManager = newCrawlerManager(conn, core)
	return core
}
