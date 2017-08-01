package main

import (
	"log"
	"net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "
			<html>
        <head>
          <title>チャット</title>
        </head>
        <body>
          チャットしましょう!
        </body>
      </html>
		")
}
