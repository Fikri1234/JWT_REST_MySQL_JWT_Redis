package router

import (
	"JWT_REST_MUX_MySQL_Redis/configuration"
	"JWT_REST_MUX_MySQL_Redis/model"
	"JWT_REST_MUX_MySQL_Redis/service"
	"JWT_REST_MUX_MySQL_Redis/util"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter ...
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	model.AppRoutes = append(model.AppRoutes, service.UserRoutes)
	model.AppRoutes = append(model.AppRoutes, service.UserDetailRoutes)

	for _, route := range model.AppRoutes {
		// create subroute
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		// loop through each sub route
		for _, r := range route.SubRoutes {
			var handler http.Handler
			handler = r.HandlerFunc

			if r.Protected {
				handler = util.VerifyInterceptorHTTP(r.HandlerFunc)
			}

			// attach sub route
			routePrefix.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method).
				Name(r.Name)
		}
	}

	router.Use(configuration.CORS)

	return router
}
