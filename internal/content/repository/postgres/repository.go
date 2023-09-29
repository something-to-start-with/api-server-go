package postgres

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/something-to-start-with/api-server-go/internal/content"
	"github.com/something-to-start-with/api-server-go/internal/models"
)

var _ content.Repository = (*Repository)(nil)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SelectAll() ([]*models.Content, error) {
	entities := make([]contentEntity, 0)
	query := `SELECT id, data FROM content`
	err := r.db.Select(&entities, query)
	if err != nil {
		return nil, errors.Wrap(err, "Error selecting from content")
	}
	return toModels(entities), nil
}

func (r *Repository) Save(c *models.Content) (*models.Content, error) {
	var id int64
	err := r.db.Get(&id, `INSERT INTO content (data) VALUES ($1) RETURNING id`, c.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Error inserting to content")
	}
	c.ID = id
	return c, nil
}

func (r *Repository) Update(id int, c *models.Content) error {
	res, err := r.db.Exec(`UPDATE content SET data = $1 WHERE id = $2`, c.Body, id)
	if err != nil {
		return errors.Wrapf(err, "Error updating content id %d", id)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.Wrapf(err, "Error updating content id %d", id)
	}
	if rowsAffected == 0 {
		return content.ErrNotFound
	}
	return nil
}

func (r *Repository) SelectByID(id int) (*models.Content, error) {
	var entity contentEntity
	query := `SELECT id, data FROM content WHERE id = $1`
	err := r.db.Get(&entity, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, content.ErrNotFound
		} else {
			return nil, errors.Wrapf(err, "Error selecting content by id %d", id)
		}
	}
	return entity.toModel(), nil
}

func (r *Repository) DeleteByID(id int) error {
	res, err := r.db.Exec(`DELETE FROM content WHERE id = $1`, id)
	if err != nil {
		return errors.Wrapf(err, "Error deleting content by id %d", id)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.Wrapf(err, "Error deleting content by id %d", id)
	}
	if rowsAffected == 0 {
		return content.ErrNotFound
	}
	return nil
}
