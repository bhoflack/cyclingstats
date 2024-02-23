package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/bhoflack/cyclingstats/pkg/model"
)

type Client struct {
	db *sql.DB
}

func NewClient() (*Client, error) {
	db, err := sql.Open("sqlite3", "file:cyclingstats.db")
	if err != nil {
		return nil, err
	}

	if err := initSchema(db); err != nil {
		return nil, err

	}
	return &Client{db: db}, nil
}

func initSchema(db *sql.DB) error {

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS cyclists (
			entryid INTEGER PRIMARY KEY,
			label text,
			id text
		);
		`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS pages (
			riderid INTEGER,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
			page text,
			PRIMARY KEY (riderid, timestamp)
		)`)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AddCyclist(cyclist model.Cyclist) error {
	_, err := c.db.Exec("INSERT INTO cyclists (entryid, label, id) VALUES (?,?,?)", cyclist.EntryId, cyclist.Label, cyclist.Id)
	return err
}

func (c *Client) ListCyclists() ([]model.Cyclist, error) {
	rows, err := c.db.Query("SELECT entryid, label, id FROM cyclists")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cyclists []model.Cyclist
	for rows.Next() {
		var c model.Cyclist
		if err := rows.Scan(&c.EntryId, &c.Label, &c.Id); err != nil {
			return nil, err
		}
		cyclists = append(cyclists, c)
	}
	return cyclists, nil
}

func (c *Client) AddPage(cyclist model.Cyclist, page string) error {
	_, err := c.db.Exec("INSERT INTO pages (riderid, page) VALUES (?,?)", cyclist.Id, page)
	return err
}

func (c *Client) GetPage(cyclist model.Cyclist) (string, error) {
	var page string
	err := c.db.QueryRow("SELECT page FROM pages WHERE riderid = ?", cyclist.Id).Scan(&page)
	if err != nil {
		return "", err
	}
	return page, nil
}

func (c *Client) GetPages() (map[string]string, error) {
	rows, err := c.db.Query("SELECT riderid, page FROM pages order by riderid, timestamp desc")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var previousRiderID string
	var riderID string
	pages := make(map[string]string)
	for rows.Next() {
		var page string
		if err := rows.Scan(&riderID, &page); err != nil {
			return nil, err
		}

		if riderID != previousRiderID {
			pages[riderID] = page
		}
	}
	return pages, nil
}

func (c *Client) Close() error {
	return c.db.Close()
}
