package app

import (
	"github.com/arun6783/bookstore_users-api/controllers/ping"
	"github.com/arun6783/bookstore_users-api/controllers/users"
)

func mapURls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/users/:user_id", users.Get)

	router.GET("internal/users/search", users.Search)

}
