package postgresql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/muhrifqii/curium_go_fiber/domain"
	"github.com/muhrifqii/curium_go_fiber/internal/sqler"
	"go.uber.org/zap"
)

type OrganizationRepository struct {
	db *sqler.SqlxWrapper
}

func NewOrganizationRepository(db *sqlx.DB, zap *zap.Logger) *OrganizationRepository {
	return &OrganizationRepository{
		db: sqler.NewSqlxWrapper(db, zap),
	}
}

func (r *OrganizationRepository) GetByID(c context.Context, id string) (domain.Organization, error) {
	var org domain.Organization
	err := r.db.Get(&org, "SELECT * FROM organizations WHERE identifier = $1", id)
	return org, err
}

func (r *OrganizationRepository) CreateOrganization(c context.Context, org *domain.Organization) error {
	_, err := r.db.NamedExec("INSERT INTO organizations (a_name) VALUES (:a_name)", org)
	return err
}

func (r *OrganizationRepository) UpdateOrganization(c context.Context, org *domain.Organization) error {
	_, err := r.db.NamedExec("UPDATE organizations SET name = :name WHERE identifier = :identifier", org)
	return err
}
