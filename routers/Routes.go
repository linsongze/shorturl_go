package routers

import "github.com/buaazp/fasthttprouter"

func InitRoutes(router *fasthttprouter.Router)  {
	router.GET("/s/:sc",redirect)
	router.GET("/",Index)
	router.GET("/to/short",Short)
}
