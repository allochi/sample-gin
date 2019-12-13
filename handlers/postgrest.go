package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// PostgrestURL URL for PostgREST - should move to configuration
const PostgrestURL = "http://localhost:3000"

// PostgrestProxy a proxy to postgrest
func PostgrestProxy(ctx *gin.Context) {
	director := func(r *http.Request) {
		uri := ctx.Request.URL.RequestURI()[5:]
		url, _ := url.Parse(PostgrestURL + uri)
		r.URL = url
		// r.Header.Add("Authorization", authorizationToken)
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
