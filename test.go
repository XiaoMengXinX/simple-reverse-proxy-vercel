package main

import (
	"net/http"

	"github.com/XiaoMengXinX/simple-reverse-proxy-vercel/api"
)

func main() {
	http.HandleFunc("/", api.ProxyHandler)
	http.ListenAndServe(":8000", nil)
}
