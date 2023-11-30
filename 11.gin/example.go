package main

import "zenkie.cn/learn-gin/router"

func main() {
	r := router.Router()
	// router.Run() // listen and serve on 0.0.0.0:8080
	r.Run("localhost:8080")
}
