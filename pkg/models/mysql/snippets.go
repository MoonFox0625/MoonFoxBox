// Package mysql
// Date:  2021/10/4 17:00
// Desc:
package mysql

import (
	"MoonFoxBox/pkg/models"
	"database/sql"
	"time"
)

// SnippetModel Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert This will insert a new snippet into the database.
func (sm *SnippetModel) Insert(title, content string, expires time.Time) (int, error) {
	return 0, nil
}

// Get This will return a specific snippet based on its id.
func (sm *SnippetModel) Get(id int) (*models.Snippets, error) {
	return nil, nil
}

// Latest  This will return the 10 most recently created snippets.
func (sm *SnippetModel) Latest() ([]*models.Snippets, error) {
	return nil, nil
}
