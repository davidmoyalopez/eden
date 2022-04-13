package app

import (
	"eden/controllers/lands"
	"eden/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/lands/:land_id", lands.Get)
	router.GET("/lands/search", lands.Search)
	router.POST("/lands", lands.Create)
}
