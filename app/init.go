package lrp

import (
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
                  へろーわーるどやで 02 ！！！
                </body>
              </html>
        `))
    })
}
