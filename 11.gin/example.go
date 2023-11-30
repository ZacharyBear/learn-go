package main

import "zenkie.cn/learn-gin/router"

// todo rename it with "main"
func start() {
	r := router.Router()
	// router.Run() // listen and serve on 0.0.0.0:8080
	r.Run("localhost:8080")
}
