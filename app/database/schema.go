package database

var migration = []string{
	`
	create table if not exists users(
		uid	UUID PRIMARY KEY, 
		email VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL,
		username VARCHAR(50),
		avatar VARCHAR(255),
		provider VARCHAR(50),
		created_at TIMESTAMP,
		updated_at TIMESTAMP
	)
	`,
}
