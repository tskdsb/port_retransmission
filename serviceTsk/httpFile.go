package serviceTsk

import (
	"net/http"
)

func HttpFileServer(pattern, dir string) {
	http.Handle(pattern, http.FileServer(http.Dir(dir)))
}

func PrefixHttpFileServer(pattern, prefix, dir string) {
	http.Handle(pattern, http.StripPrefix(prefix, http.FileServer(http.Dir(dir))))
}
