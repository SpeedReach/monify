package media

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"monify/lib/media"
)

type Infra struct {
	db         *sql.DB
	objStorage media.Storage
	config     Config
	logger     *zap.Logger
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
		objStorage: NewS3MediaStorage(config),
	}, nil
}
