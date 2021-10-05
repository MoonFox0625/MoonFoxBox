// Package mysql
// Date:  2021/10/4 17:00
// Desc:
package mysql

import (
	"MoonFoxBox/pkg/models"
	"database/sql"
	"errors"
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
func (sm *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires
FROM snippets
WHERE expires > UTC_TIMESTAMP()
  AND id = ?;`

	var s = &models.Snippet{}
	// row := sm.DB.QueryRow(stmt, id)
	// err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	// 上面的代码可以简写成下面
	err := sm.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// Latest  This will return the 10 most recently created snippets.
func (sm *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires
FROM snippets
WHERE UTC_TIMESTAMP() < expires
ORDER BY created DESC
LIMIT 10;`

	rows, err := sm.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var snippets []*models.Snippet

	for rows.Next() {
		var s = &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
