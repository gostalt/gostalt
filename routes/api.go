package routes

import (
	"net/http"

	"github.com/gostalt/framework/route"
	mw "github.com/gostalt/framework/route/middleware"

	"gostalt/app/http/handler/api"
)

var API = route.Collection(
	route.Get("/", http.HandlerFunc(api.Welcome)),
).Prefix("api").Middleware(mw.AddJSONContentTypeHeader)
