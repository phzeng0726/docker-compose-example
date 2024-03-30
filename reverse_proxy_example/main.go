package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// ReverseProxyHandler 生成並返回處理反向代理請求的處理函數
func ReverseProxyHandler(target *url.URL) http.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(target)

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		r.Host = target.Host
		w.Header().Set("X-Ben", "Rad")
		proxy.ServeHTTP(w, r)
	}
}

// parseURL 將給定的 URL 字符串解析為 *url.URL 對象
func parseURL(rawURL string) *url.URL {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}
	return parsedURL
}

func main() {
	// 定義反向代理目標
	targets := map[string]*url.URL{
		"/":      parseURL("http://google.com"),
		"/api1/": parseURL("http://172.25.0.3:5000"), // Python
		"/api2/": parseURL("http://172.25.0.4:5001"), // Go
	}

	// 設置反向代理處理函數
	for path, target := range targets {
		http.HandleFunc(path, ReverseProxyHandler(target))
	}

	// 啟動服務器
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
