package postgresql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/muhrifqii/curium_go_fiber/domain"
	"github.com/muhrifqii/curium_go_fiber/internal/sqler"
	"go.uber.org/zap"
)

type OAuthClientRepository struct {
	db *sqler.SqlxWrapper
}

func NewOauthClientRepository(db *sqlx.DB, zap *zap.Logger) *OAuthClientRepository {
	return &OAuthClientRepository{
		db: sqler.NewSqlxWrapper(db, zap),
	}
}

func (r *OAuthClientRepository) GetByID(c context.Context, clientID string) (domain.OAuthClient, error) {
	var client domain.OAuthClient
	err := r.db.Get(&client, "SELECT * FROM clients WHERE client_id = $1", clientID)
	return client, err
}

func (r *OAuthClientRepository) Exists(c context.Context, clientID, clientSecret string) (bool, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM clients WHERE client_id = $1 AND client_secret = $2", clientID, clientSecret)
	return count > 0, err
}

func (r *OAuthClientRepository) ExistsByID(c context.Context, clientID string) (bool, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM clients WHERE client_id = $1", clientID)
	return count > 0, err
}

func (r *OAuthClientRepository) CreateClient(c context.Context, client *domain.OAuthClient) error {
	query := `
		INSERT INTO clients (
			client_id,
			client_secret,
			client_name,
			grant_types,
			scope,
			redirect_uris,
			post_logout_redirect_uris
		)
		VALUES (
			:client_id,
			:client_secret,
			:client_name,
			:grant_types,
			:scope,
			:redirect_uris,
			:post_logout_redirect_uris
		)
	`
	_, err := r.db.NamedExec(query, client)
	return err
}

func (r *OAuthClientRepository) UpdateClient(c context.Context, client *domain.OAuthClient) error {
	query := `
		UPDATE clients
		SET
			client_secret = :client_secret,
			client_name = :client_name,
			scope = :scope,
			grant_types = :grant_types,
			redirect_uris = :redirect_uris,
			post_logout_redirect_uris = :post_logout_redirect_uris,
			updated_at = CURRENT_TIMESTAMP
		WHERE client_id = :client_id
	`
	_, err := r.db.NamedExec(query, client)
	return err
}
