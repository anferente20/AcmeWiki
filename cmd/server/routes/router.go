package routes

import (
	"os"

	"github.com/anferente20/AcmeWiki/cmd/server/handler"
	"github.com/anferente20/AcmeWiki/cmd/server/middleware"
	"github.com/anferente20/AcmeWiki/internal/pages"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
}

func NewRouter(eng *gin.Engine) Router {
	return &router{eng: eng}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildPagesRoutes()
}

func (r *router) setGroup() {
	token := os.Getenv("TOKEN")
	au := middleware.NewAuthenticator(token)
	r.rg = r.eng.Group("/api/v1")
	r.rg.Use(au.Authenticate())
}

func (r *router) buildPagesRoutes() {
	repo := pages.NewPageRepository()
	service := pages.NewPageService(repo)
	_ = handler.NewPageHandler(service)
}
