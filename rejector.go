package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var UAprefix = StrList("USERAGENT_PREFIX", "GoogleHC/")
var HealthPath = StrList("HEALTH_PATH", "")

func StrList(envvar, defaultValue string) []string {
	strs := os.Getenv(envvar)
	if "" == strs {
		strs = defaultValue
	}
	list := strings.Split(strs, ",")
	o := 0
	for _, s := range list {
		if "" != s {
			list[o] = s
			o++
		}
	}
	return list[:o]
}

func IsHealthCheck(req *http.Request) bool {
	if "GET" != req.Method {
		return false
	}
	for _, path := range HealthPath {
		if path == req.RequestURI {
			return true
		}
	}
	ua := req.UserAgent()
	for _, pref := range UAprefix {
		if strings.HasPrefix(ua, pref) {
			return true
		}
	}
	return false
}

func Handle(rw http.ResponseWriter, req *http.Request) {
	if IsHealthCheck(req) {
		rw.WriteHeader(200)
		return
	}
	rw.WriteHeader(403)
}

func main() {
	srv := &http.Server{
		Addr:              ":8000",
		Handler:           http.HandlerFunc(Handle),
		ReadHeaderTimeout: 15*time.Second,
		MaxHeaderBytes:    16*1024,
		IdleTimeout:       1*time.Second,
		ReadTimeout:       30*time.Second,
		WriteTimeout:      20*time.Second,
		// Disable h2_bundle:
		TLSNextProto:      make(
			map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	err := srv.ListenAndServe()
	if nil != err && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "Could not listed at %s: %v\n", srv.Addr, err)
	}
}
