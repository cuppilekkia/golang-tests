package app

import (
	"microservices-course/mvc-project/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
