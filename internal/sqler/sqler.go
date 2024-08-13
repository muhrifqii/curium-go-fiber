package sqler

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type SqlxWrapper struct {
	*sqlx.DB
	log *zap.Logger
}

func NewSqlxWrapper(db *sqlx.DB, log *zap.Logger) *SqlxWrapper {
	return &SqlxWrapper{
		DB:  db,
		log: log,
	}
}

func (s *SqlxWrapper) Beginx() (*sqlx.Tx, error) {
	tx, err := s.DB.Beginx()
	if err != nil {
		s.log.Error("failed to begin transaction", zap.Error(err))
	}
	return tx, err
}

func (s *SqlxWrapper) Get(dest interface{}, query string, args ...interface{}) error {
	err := s.DB.Get(dest, query, args...)
	if err != nil {
		s.log.Error("failed to execute Get wit dest query", zap.Error(err), zap.String("query", query))
	}
	return err
}

func (s *SqlxWrapper) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := s.DB.Exec(query, args...)
	if err != nil {
		s.log.Error("failed to execute Exec query", zap.Error(err), zap.String("query", query))
	}
	return result, err
}

func (s *SqlxWrapper) MustExec(query string, args ...interface{}) sql.Result {
	result, err := s.DB.Exec(query, args...)
	if err != nil {
		s.log.Error("failed to execute MustExec query", zap.Error(err), zap.String("query", query))
	}
	return result
}

func (s *SqlxWrapper) NamedExec(query string, args interface{}) (sql.Result, error) {
	result, err := s.DB.NamedExec(query, args)
	if err != nil {
		s.log.Error("failed to execute NamedExec query", zap.Error(err), zap.String("query", query))
	}
	return result, err
}

func (s *SqlxWrapper) NamedQuery(query string, args interface{}) (*sqlx.Rows, error) {
	rows, err := s.DB.NamedQuery(query, args)
	if err != nil {
		s.log.Error("failed to execute NamedQuery query", zap.Error(err), zap.String("query", query))
	}
	return rows, err
}
