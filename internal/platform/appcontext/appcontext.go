package appcontext

import (
	"database/sql"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/datasources/repositories"
	"github.com/tapiaw38/cardon-tour-be/internal/adapters/storeservice"
)

type Context struct {
	Repositories *repositories.Repositories
	StoreService storeservice.StoreService
}

type Factory func() *Context

type Options func(*Context)

func NewFactory(db *sql.DB, storeService storeservice.StoreService) Factory {
	return func() *Context {
		return &Context{
			Repositories: repositories.CreateRepositories(db),
			StoreService: storeService,
		}
	}
}
