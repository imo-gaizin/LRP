package lrp

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "へろーわーるどやで！！！")
        w.Write(
            []byte(`
              <html>
                <head>
                  <title>チャット</title>
                </head>
                <body>
                  チャットしましょう!
                </body>
              </html>
            `)
        )
    })
}
