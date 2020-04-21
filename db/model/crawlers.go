package model

import (
	"github.com/huandu/go-sqlbuilder"
	"github.com/pborman/uuid"
)

type crawlerManager struct {
	*manager
}

func newCrawlerManager(conn SQLConnection, core *Managers) *crawlerManager {
	tableStruct := sqlbuilder.NewStruct(new(Crawler))
	tableStruct.Flavor = sqlbuilder.PostgreSQL

	return &crawlerManager{
		manager: &manager{
			core:        core,
			connection:  conn,
			tableStruct: tableStruct,
			tableName:   "crawlers",
		},
	}
}

type CrawlerManager interface {
	GetCrawlers() ([]*Crawler, error)
	CreateCrawler(name, certificate string) (int, uuid.UUID, error)
	GetCrawler(id int) (*Crawler, error)
	DeleteCrawler(id int) error
}

func (m crawlerManager) GetCrawlers() ([]*Crawler, error) {
	sb := m.tableStruct.SelectFrom(m.tableName)
	sql, args := sb.Build()

	rows, err := m.connection.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	var items []*Crawler

	defer rows.Close()
	for rows.Next() {
		var item Crawler
		if err := rows.Scan(m.tableStruct.Addr(&item)...); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}

func (m crawlerManager) CreateCrawler(name, certificate string) (int, uuid.UUID, error) {
	reference := uuid.NewRandom()

	ib := sqlbuilder.NewInsertBuilder().InsertInto(m.tableName)
	ib = ib.Cols("reference", "name", "certificate").Values(reference, name, certificate)

	builder := sqlbuilder.Build("$? returning id;", ib)
	sql, args := builder.Build()

	rows, err := m.connection.Query(sql, args...)
	if err != nil {
		return 0, uuid.NIL, err
	}

	var id int
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, uuid.NIL, err
		}
	}

	return id, reference, nil
}

func (m crawlerManager) GetCrawler(id int) (*Crawler, error) {
	sb := m.tableStruct.SelectFrom(m.tableName)
	sb = sb.Where(sb.Equal("id", id))
	sql, args := sb.Build()

	rows, err := m.connection.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	var crawler Crawler
	if err := rows.Scan(m.tableStruct.Addr(&crawler)...); err != nil {
		return nil, err
	}

	return &crawler, nil
}

func (m crawlerManager) DeleteCrawler(id int) error {
	sb := m.tableStruct.DeleteFrom(m.tableName)
	sb = sb.Where(sb.Equal("id", id))
	sql, args := sb.Build()

	_, err := m.connection.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}
