package main

import (
	"net/http"
	"rest-api/packages/routes"
)



func main() {
	routes.InitRoutes()
	http.ListenAndServe(":8000",nil)
}


