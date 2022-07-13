package router

import (
	"github.com/etcsilver/go-examples/api"
	"github.com/etcsilver/go-examples/domain"
	"github.com/etcsilver/go-examples/pkg/ws"

	"github.com/etcsilver/go-examples/pkg/application"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Get(app *application.Application) *gin.Engine {
	log.Info().Msg("Definiendo metodos...")
	if app.CfgApp.Ambiente == "produccion" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	ws := ws.NewWebServices(app)
	service := domain.NewWorkflowService(app, ws)
	handler := api.NewHandler(service)

	r := gin.New()
	//Example handler
	r.POST("/hello", func(c *gin.Context) { handler.Hello(c) })
	return r
}
