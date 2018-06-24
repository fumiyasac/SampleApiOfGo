package middleware

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// ReverseProxyHandler ... リバースプロキシを作成する
func ReverseProxyHandler(req *http.Request, res http.ResponseWriter) {

	// MEMO: リモートAPIのOriginURLを記載する
	remoteURL, error := url.Parse("https://www.example.com/api")
	if error != nil {
		log.Fatal(error)
	}
	proxy := httputil.NewSingleHostReverseProxy(remoteURL)
	proxy.ServeHTTP(res, req)
}
