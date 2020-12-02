package router

import (
	"github.com/gin-gonic/gin"
	branchController "xpertise-go/branch/controller"
	portalController "xpertise-go/portal/controller"
	"xpertise-go/user/auth"
	userController "xpertise-go/user/controller"
)

// SetupRouter contains all the api that will be used.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// adminV1 := r.Group("api/v1/admin")

	branchV1 := r.Group("api/v1/branch")
	{
		branchV1.POST("/create", branchController.CreateAComment)
	}

	portalV1 := r.Group("api/v1/portal")
	{
		portalV1.POST("/document/create", portalController.CreateADocument)
		portalV1.GET("/document/query", portalController.QueryDocumentByID)
	}
	userV1 := r.Group("api/v1/user")
	{
		userV1.DELETE("/delete/:id", userController.DeleteAStudentByID)
		userV1.PUT("/update/age", userController.UpdateAStudentByAge)
		userV1.GET("/query/all", userController.QueryAllStudents)
		userV1.GET("/query/id", userController.QueryStudentByID)
		userV1.GET("/query/age", userController.QueryStudentsByAge)
		userV1.POST("/register", userController.Register)
		userV1.POST("/login", userController.Login)

		userV1.POST("/password/reset",userController.ResetPassword)
		userV1.POST("/folder/create", auth.JwtAuth(),userController.CreateAFolder)
		userV1.POST("/folder/add",auth.JwtAuth(),userController.AddToMyFolder)
		userV1.POST("/account_info/reset",userController.ResetAccountInfo)
	}
	return r
}
