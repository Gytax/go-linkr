package db

const schemaSQL = `
CREATE TABLE IF NOT EXISTS links (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  slug TEXT NOT NULL,
  target TEXT NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_links_slug_unique ON links(slug);
DROP INDEX IF EXISTS idx_links_slug;
`

func (d *DB) initializeSchema() error {
	_, err := d.db.Exec(schemaSQL)
	return err
}
