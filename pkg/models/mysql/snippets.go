// Package mysql
// Date:  2021/10/4 17:00
// Desc:
package mysql

import (
	"MoonFoxBox/pkg/models"
	"database/sql"
)

// SnippetModel Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert This will insert a new snippet into the database.
func (sm *SnippetModel) Insert(title, content, expires string) (int, error) {
	// Write the SQL statement we want to execute. I've split it over two lines
	// for readability (which is why it's surrounded with backquotes instead
	// of normal double quotes).
	stmt := `INSERT INTO snippets(title, content, created, expires)
VALUES (?, ?, UTC_TIMESTAMP,DATE_ADD(UTC_TIMESTAMP,INTERVAL ? DAY ))`

	result, err := sm.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

// Get This will return a specific snippet based on its id.
func (sm *SnippetModel) Get(id int) (*models.Snippets, error) {
	return nil, nil
}

// Latest  This will return the 10 most recently created snippets.
func (sm *SnippetModel) Latest() ([]*models.Snippets, error) {
	return nil, nil
}
