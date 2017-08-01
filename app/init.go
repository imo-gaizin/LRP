package lrp

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(`
              <html>
                <head>
                  <title>へろーわーるど</title>
                </head>
                <body>
                  へろーわーるどやで！！！
                </body>
              </html>
        `))
    })
}
