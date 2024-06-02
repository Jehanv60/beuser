package app

import (
	"github.com/Jehanv60/controller"
	"github.com/Jehanv60/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(penggunaController controller.PenggunaController) *httprouter.Router {
	router := httprouter.New()
	//pengguna
	router.GET("/api/pengguna", penggunaController.FindAll)
	router.GET("/api/pengguna/:penggunaId", penggunaController.FindById)
	router.PUT("/api/pengguna/:penggunaId", penggunaController.Update)
	router.DELETE("/api/pengguna/:penggunaId", penggunaController.Delete)
	router.POST("/api/pengguna", penggunaController.Create)
	router.POST("/api/login", penggunaController.LoginAuth)
	router.POST("/api/logout", controller.Logout)
	router.PanicHandler = exception.ErrorHandler
	return router
}
