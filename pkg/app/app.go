package app

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/starnuik/avito_pvz/pkg/gen"
	"github.com/starnuik/avito_pvz/pkg/handler"
	"github.com/starnuik/avito_pvz/pkg/middleware"
	"github.com/starnuik/avito_pvz/pkg/password"
	"github.com/starnuik/avito_pvz/pkg/repository"
	"github.com/starnuik/avito_pvz/pkg/token"
	"github.com/starnuik/avito_pvz/pkg/usecase"
)

type App interface {
	http.Handler
	Close() error

	Run()
}

type app struct {
	close func() error
	*gin.Engine
}

func New() (App, error) {
	hasher := password.NewHasher()
	gen := gen.New()

	// TODO context, conn string
	repo, err := repository.New(context.TODO(), "postgres://postgres:postgres@localhost:5432/pvz")
	if err != nil {
		return nil, err
	}

	usecase := usecase.New(repo, hasher, gen)

	tokenParser := token.NewParser()
	handler := handler.New(usecase, tokenParser)

	router := gin.Default()

	router.GET("/ping", handler.GetPing)

	// Auth
	router.POST("/dummyLogin", handler.PostDummyLogin)
	router.POST("/register", handler.PostRegister)
	router.POST("/login", handler.PostLogin)

	authGroup := router.Group("", middleware.AuthCheck(tokenParser))

	// Create
	authGroup.POST("/pvz", handler.PostPvz)
	authGroup.POST("/rececptions", handler.PostReceptions)
	authGroup.POST("/products", handler.PostProducts)

	// Read
	authGroup.GET("/pvz", handler.GetPvz)

	// Update / Delete
	authGroup.POST("/pvz/:id/close_last_reception", handler.PostCloseLastReception)
	authGroup.POST("/pvz/:id/delete_last_product", handler.PostDeleteLastProduct)

	return &app{
		Engine: router,
		close:  func() error { return repo.Close(context.TODO()) },
	}, nil
}

func (app *app) Run() {
	// TODO graceful shutdown
	app.Engine.Run(":8080")
}

func (app *app) Close() error {
	return app.close()
}
