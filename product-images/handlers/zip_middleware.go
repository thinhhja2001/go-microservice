package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipHandler struct {
}

func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			gw := gzip.NewWriter(rw)
			gw.Write([]byte("deptrai"))
			println(gw.)
			return
		}
		next.ServeHTTP(rw, r)
	})
}

type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

func NewWrappedResponseWriter(rw *http.ResponseWriter) *WrappedResponseWriter {
	gw:= gzip.NewWriter(rw)
	return &WrappedResponseWriter{rw: rw,gw: gw}
}

