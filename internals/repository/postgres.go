package repository

import (
	"context"
	"shortlink/internals/domain/entity"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresDbPool(dbUrl string) (dbPool *pgxpool.Pool, err error) {
	dbPool, err = pgxpool.Connect(context.Background(), dbUrl)

	if err != nil {
		return nil, err
	}

	return dbPool, nil
}

type PostgresRepository struct {
	dbPool *pgxpool.Pool
}

func NewPostgresRepository(dbPool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{dbPool: dbPool}
}

func (repo *PostgresRepository) Add(link *entity.Link) error {
	query := `
		INSERT INTO 
			links (id, short_url, origin_url) 
		VALUES 
			($1, $2, $3)
		;
	`
	row := repo.dbPool.QueryRow(context.Background(), query, link.ID, link.ShortURL, link.OriginalURL)
	err := row.Scan()

	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepository) FindShortURLByOriginalURL(OriginalURL string) (string, error) {
	var ShortURL string
	query := `
		SELECT 
			short_url 
		FROM 
			links
		WHERE
			origin_url = $1
		;
	`
	row := repo.dbPool.QueryRow(context.Background(), query, OriginalURL)
	err := row.Scan(&ShortURL)

	if err != nil {
		return "", nil
	}
	return ShortURL, nil
}

func (repo *PostgresRepository) FindOriginalURLByShortURL(ShortURL string) (string, error) {
	var OriginalURL string
	query := `
		SELECT 
			origin_url 
		FROM 
			links
		WHERE
			short_url = $1
		;
	`
	row := repo.dbPool.QueryRow(context.Background(), query, ShortURL)
	err := row.Scan(&OriginalURL)

	if err != nil {
		return "", nil
	}
	return OriginalURL, nil
}
