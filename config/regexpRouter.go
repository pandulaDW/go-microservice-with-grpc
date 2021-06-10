package config

import (
	"net/http"
	"regexp"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpRouter struct {
	routes []*route
}

func (router *RegexpRouter) Handler(pattern *regexp.Regexp, handler http.Handler) {
	router.routes = append(router.routes, &route{pattern, handler})
}

func (router *RegexpRouter) ServeHTTP(rw http.ResponseWriter, res *http.Request) {
	for _, r := range router.routes {
		if r.pattern.MatchString(res.URL.Path) {
			r.handler.ServeHTTP(rw, res)
			return
		}
	}
	// no pattern matched, send 404 response
	http.NotFound(rw, res)
}

func (router *RegexpRouter) HandleFunc(pattern *regexp.Regexp, handler http.HandlerFunc) {
	newRoute := route{pattern: pattern, handler: handler}
	router.routes = append(router.routes, &newRoute)
}
