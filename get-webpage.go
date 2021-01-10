package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type webpage struct {
	url string
	body []byte
	err error
}

func (w *webpage) get() {
	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}
	defer resp.Body.Close()

	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
		return
	}
}

func (w *webpage) isOk() bool {
	return w.err == nil
}

func main() {
	w := &webpage{url: "https://www.amazon.com.au"}
	w.get()
	if w.isOk() {
		fmt.Printf("URL: %s returns bytes: %d", w.url, len(w.body))
	} else {
		fmt.Printf("URL: %s returns error: %s and bytes: %d", w.url, w.err, len(w.body))
	}
}
