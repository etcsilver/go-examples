package ws

import (
	"github.com/etcsilver/go-examples/domain"
	"github.com/etcsilver/go-examples/pkg/application"
)

type webservices struct {
	app *application.Application
}

//NewWorkflowService...
func NewWebServices(app *application.Application) domain.WebServices {
	return &webservices{app: app}
}
