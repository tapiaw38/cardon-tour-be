package appcontext

import (
	"database/sql"

	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories"
)

type Context struct {
	Repositories *repositories.Repositories
}

type Factory func() *Context

type Options func(*Context)

func NewFactory(db *sql.DB) func() *Context {
	return func() *Context {
		return &Context{
			Repositories: repositories.CreateRepositories(db),
		}
	}
}
