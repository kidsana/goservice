package main
import (
    "github.com/kidsana/goservice/routes"
    "log"
    "net/http"
)
// init main func
func main() {
    route := NewRouter()
    log.Fatal(
        // start on port 3000 by default
        http.ListenAndServe(":3000", route),
    )
}
