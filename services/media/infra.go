package media

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Infra struct {
	db         *sql.DB
	objStorage S3ImageStorage
	config     Config
}

func Setup(config Config) (Infra, error) {
	db, err := sql.Open("pgx", config.PostgresURI)
	if err != nil {
		return Infra{}, err
	}
	if err = db.Ping(); err != nil {
		return Infra{}, err
	}

	return Infra{
		db:         db,
		config:     config,
		objStorage: NewS3ImageStorage(config),
	}, nil
}
