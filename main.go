package main
import (
    "handlers"
    "routes"
    "log"
    "net/http"
)
// init main func
func main() {
    router := NewRouter()
    log.Fatal(
        // start on port 3000 by default
        http.ListenAndServe(":3000", router),
    )
}