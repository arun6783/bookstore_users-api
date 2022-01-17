package app

import (
	"github.com/arun6783/bookstore_users-api/controllers/ping"
	"github.com/arun6783/bookstore_users-api/controllers/users"
)

func mapURls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)

	router.GET("/users/:user_id", users.GetUser)

	router.GET("/users/search", users.SearchUser)

}
