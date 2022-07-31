package api

import (
	"io/ioutil"
	"net/http"
	"os"
)

var URLtoProxy = os.Getenv("URL_TO_PROXY")

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	var url = URLtoProxy + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		url += "?" + r.URL.RawQuery
	}
	req, err := http.NewRequest(r.Method, URLtoProxy+r.URL.Path, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(resp.StatusCode)
	resBody, err := ioutil.ReadAll(resp.Body)
	w.Write(resBody)
}
