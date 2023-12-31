package router

import (
	"context"

	"github.com/letsvote/api/internal/controller/file"
	"github.com/letsvote/api/internal/controller/user"
	"github.com/letsvote/api/internal/handler/authenticated"
	"github.com/letsvote/api/internal/handler/public"
)

// New creates and returns a new Router instance
func New(
	ctx context.Context,
	userCtrl user.Controller,
	fileCtrl file.Controller,
) Router {
	return Router{
		ctx:                  ctx,
		corsOrigins:          nil,
		authenticatedHandler: authenticated.New(userCtrl, fileCtrl),
		publicHandler:        public.New(userCtrl),
	}
}
