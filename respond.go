package mbot

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type customdata struct {
	name string
	vals []string
}

var d customdata

func create(name string) string {
	d = customdata{name: name}
	PersistAdd("test", name)
	return "ok"
}

func appendToData(msg string) string {
	d.vals = append(d.vals, msg)
	return strings.Join(d.vals, ", ")
}

func openUrl(urlToOpen string) string {
	val, err := PersistGet("test")
	if err != nil {
		panic(err)
	}

	values := url.Values{"url": {"http://example.com"}}
	resp, _ := http.PostForm("http://localhost:8080", values)
	fmt.Printf("%v ", resp)

	return val
}

func get(msg string) string {
	if strings.HasPrefix(msg, "create") {
		return create(msg)
	}
	if strings.HasPrefix(msg, "open") {
		return openUrl(msg)
	}
	return appendToData(msg)
}
