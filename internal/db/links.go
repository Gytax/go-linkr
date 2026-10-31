package db

import (
	"database/sql"
	"fmt"
)

type Link struct {
	ID     int64  `db:"id"`
	Slug   string `db:"slug"`
	Target string `db:"target"`
}

func (d *DB) FindLinkBySlug(slug string) (Link, error) {
	if d == nil || d.db == nil {
		return Link{}, fmt.Errorf("Database is not available")
	}

	var link Link

	err := d.db.QueryRowx(`
    SELECT id, slug, target
    FROM links
    WHERE slug = ?
  `, slug).Scan(&link.ID, &link.Slug, &link.Target)

	if err != nil {
		if err == sql.ErrNoRows {
			return Link{}, fmt.Errorf("No link found for slug %s", slug)
		}
		return Link{}, fmt.Errorf("Error fetching link with slug %s: %w", slug, err)
	}

	return link, nil
}
