package proxyServer

import (
	"io"
	"net/http"

	"github.com/neutralusername/systemge/configs"
	"github.com/neutralusername/systemge/httpServer"
)

func New() (*httpServer.HTTPServer, error) {
	return httpServer.New(
		"httpServer",
		&configs.HTTPServer{
			TcpListenerConfig: &configs.TcpListener{
				Port:        8080,
				Domain:      "localhost",
				TlsCertPath: "../main/MyCertificate.crt",
				TlsKeyPath:  "../main/MyKey.key",
			},
		},
		nil,
		httpServer.HandlerFuncs{
			"/": func(w http.ResponseWriter, r *http.Request) {
				// retrieve query parameter "url"
				url := r.URL.Query().Get("url")

				// get response from this url
				response, err := http.Get(url)
				if err != nil {
					println(err.Error(), url)
					http.Error(w, "Failed to fetch URL", http.StatusBadRequest)
					return
				}
				defer response.Body.Close()

				// copy status code from the fetched response
				w.WriteHeader(response.StatusCode)

				// copy headers from the fetched response
				for key, values := range response.Header {
					for _, value := range values {
						w.Header().Add(key, value)
					}
				}

				// copy body from the fetched response
				_, err = io.Copy(w, response.Body)
				if err != nil {
					http.Error(w, "Failed to read response body", http.StatusInternalServerError)
				}
			},
		},
	)
}
