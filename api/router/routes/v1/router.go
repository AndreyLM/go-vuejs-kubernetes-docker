package v1

import (
	"net/http"

	"github.com/learning/project/api/jwt"

	"github.com/go-xorm/xorm"
	"github.com/learning/project/api/router/routes"
)

// Middleware - sub route middleware
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")
		if _, err := jwt.ParseToken(token); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// GetRoutes - get v1 routes
func GetRoutes(db *xorm.Engine) (SubRoute map[string]routes.SubRoutePackage) {
	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": routes.SubRoutePackage{
			Routes: routes.Routes{
				routes.Route{
					Name:        "V1HalrthToute",
					Method:      "GET",
					Pattern:     "/health",
					HandlerFunc: Health(db),
				},
			},
			Middleware: Middleware,
		},
	}
	return SubRoute
}
