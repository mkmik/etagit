package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	data     = map[string]string{}
	counters = map[string]int{}
)

func handle(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		req.ParseForm()
		data[req.Form.Get("etag")] = req.Form.Get("text")

		// http redirect works in chrome but not in firefox, let's make it explicit
		fmt.Fprintf(w, `<html><a href="/">back</a></html>`)
		return
	}

	var welcome string
	etag := req.Header.Get("If-None-Match")

	if etag == "" {
		etag = etagGen.Generate()
	} else {
		welcome = fmt.Sprintf(`<p>you visited %d times before (your session ID is %q)</p>`, counters[etag], etag)
	}

	counters[etag]++
	w.Header().Set("ETag", etag)

	fmt.Fprintf(w, `<html>
        <body>
          <p>Write something to be stored in the session:</p>
          <form method="POST" action="/process">
            <input type="text" name="text" value="%s"><br>
            <input type="hidden" name="etag" value="%s"><br>
            <input type="submit" value="submit">
            %s
          </form>
        </body>
      </html>`,
		data[etag], etag, welcome)
}

func run() error {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", handle)
	return http.ListenAndServe(":8080", nil)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
