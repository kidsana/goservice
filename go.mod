module github.com/kidsana/goservice

go 1.17

require github.com/gorilla/mux v1.8.0 // indirect
require github.com/kidsana/goservice

//replace github.com/kidsana/goservice/handlers => ./handlers
//replace github.com/kidsana/goservice/routes => ./routes
//replace github.com/kidsana/goservice/structs => ./structs
