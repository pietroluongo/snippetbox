package postgres

import (
	"database/sql"

	"pietroluongo.com/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `
	INSERT INTO snippets (title, content, created, expires)
	VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + ($3 || ' days')::INTERVAL)
	RETURNING ID;
	`

	result := m.DB.QueryRow(stmt, title, content, expires)

	if result.Err() != nil {
		return 0, result.Err()
	}

	var id int64

	err := result.Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
