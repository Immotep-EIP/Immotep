package router

import (
	"github.com/gin-gonic/gin"
	"immotep/backend/controllers"
	_ "immotep/backend/docs" // mandatory import for swagger doc
	"immotep/backend/router/middlewares"
)

func registerTenantRoutes(tenant *gin.RouterGroup) {
	tenant.Use(middlewares.AuthorizeTenant())

	tenant.POST("/invite/:id/", controllers.AcceptInvite)

	leases := tenant.Group("/leases/")
	{
		leases.GET("/", controllers.GetAllLeasesByTenant)

		leaseId := leases.Group("/:lease_id/")
		{
			leaseId.Use(middlewares.CheckLeaseTenantOwnership("lease_id"))
			leaseId.GET("/", controllers.GetLease)

			property := leaseId.Group("/property/")
			{
				property.Use(middlewares.GetPropertyByLease())
				property.GET("/", controllers.GetProperty)
				property.GET("/picture/", controllers.GetPropertyPicture)
				property.GET("/inventory/", controllers.GetPropertyInventory)
			}

			docs := leaseId.Group("/docs/")
			{
				docs.POST("/", controllers.UploadDocument)
				docs.GET("/", controllers.GetAllDocumentsByLease)

				docId := docs.Group("/:doc_id/")
				{
					docId.Use(middlewares.CheckDocumentLeaseOwnership("doc_id"))
					docId.GET("/", controllers.GetDocument)
				}
			}
		}
	}
}
