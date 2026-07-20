package di

import (
	"github.com/murashi19/koda-b8-backend1/internal/handlers"
	"github.com/murashi19/koda-b8-backend1/internal/models"
	"github.com/murashi19/koda-b8-backend1/internal/repo"
	"github.com/murashi19/koda-b8-backend1/internal/service"
)

type Container struct {
	userData    *[]models.User
	userRepo    *repo.UserRepo
	userService *service.UserService
	authHandler *handlers.AuthHandler
}

func (c *Container) initDeps() {
	c.userRepo = repo.NewUserRepo(c.userData)
	c.userService = service.NewUserService(c.userRepo)
	c.authHandler = handlers.NewAuthHandler(c.userService)
}

func (c *Container) AuthHandler() *handlers.AuthHandler {
	return c.authHandler
}

func NewContainer() *Container {
	container := &Container{userData: &[]models.User{}}
	container.initDeps()
	return container
}
