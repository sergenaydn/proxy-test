package handlers

import (
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type App struct {
	Router *gin.Engine
	DB     *clientv3.Client
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.indexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/posts", a.CreatePostHandler()).Methods("POST")
	a.Router.HandleFunc("/api/posts", a.GetPostsHandler()).Methods("GET")
	a.Router.HandleFunc("/api/posts/{id}", a.DeletePostHandler()).Methods("DELETE")
	a.Router.HandleFunc("/api/posts/{id}", a.GetPostByIdHandler()).Methods("GET")

}
