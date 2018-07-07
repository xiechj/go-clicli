package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandler()*httprouter.Router{
	router:=httprouter.New()

	router.POST("/user",CreateUser)
	router.POST("/login",Login)
	router.GET("/posts",AllPosts)

	return router
}
func main(){
	r:=RegisterHandler()
	http.ListenAndServe(":4000",r)
}