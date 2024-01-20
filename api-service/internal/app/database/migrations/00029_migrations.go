package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up, Down)
}

var queryString = `
CREATE TABLE IF NOT EXISTS users ( 
	id serial PRIMARY KEY NOT NULL, 
	email TEXT NOT NULL, 
	password TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW());

CREATE TABLE IF NOT EXISTS sessions (
	id serial PRIMARY KEY NOT NULL,
	user_id INTEGER, 
	refresh_token TEXT,
	expired_at TIMESTAMP,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
	CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE);

CREATE TABLE IF NOT EXISTS projects (
	id serial PRIMARY KEY NOT NULL,
	title TEXT,
	description TEXT,
	user_id INTEGER,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
	CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);	

CREATE TABLE IF NOT EXISTS tasks (
	id serial PRIMARY KEY NOT NULL,
	title TEXT, 
	description TEXT,
	project_id INTEGER,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
	CONSTRAINT fk_project FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
	);	
	

CREATE TABLE IF NOT EXISTS comments (
	id serial PRIMARY KEY NOT NULL,
	title TEXT, 
	description TEXT,
	task_id INTEGER,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW(),
	CONSTRAINT fk_task FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
	);
`

func Up(tx *sql.Tx) error {
	_, err := tx.Exec(queryString)
	if err != nil {
		return err
	}
	return nil
}

func Down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users;")
	if err != nil {
		return err
	}
	return nil
}
