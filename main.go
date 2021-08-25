package main
import (
    "github.com/kidsana/goservice"
    "log"
    "net/http"
)
// init main func
func main() {
    routes := NewRouter()
    log.Fatal(
        // start on port 3000 by default
        http.ListenAndServe(":3000", routes),
    )
}
