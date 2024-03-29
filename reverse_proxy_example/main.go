package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// 建立目標 URL
	targetURL1, err := url.Parse("http://127.0.0.1:5000")
	if err != nil {
		log.Fatal("Error parsing target URL1:", err)
	}

	targetURL2, err := url.Parse("http://127.0.0.1:5001")
	if err != nil {
		log.Fatal("Error parsing target URL2:", err)
	}

	// 建立反向代理
	proxy1 := httputil.NewSingleHostReverseProxy(targetURL1)
	proxy2 := httputil.NewSingleHostReverseProxy(targetURL2)

	// 設定代理的路徑
	router.PathPrefix("/").Handler(proxy1)
	router.PathPrefix("/v2/").Handler(proxy2)

	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
