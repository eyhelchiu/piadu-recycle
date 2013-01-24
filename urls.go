package main

import (
	"net/http"
)

var (
	handlers = map[string] func(http.ResponseWriter, *http.Request) {
		"/": indexHandler,
		"/append": recycleAppendHandler,
		"/search": recycleSearchHandler,
		"/wap": mobileHandler,
		"/wap/geo": mobileGeoHandler,
	}
)
