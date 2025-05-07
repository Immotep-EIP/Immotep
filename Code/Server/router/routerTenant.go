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
				property.GET("/inventory/", controllers.GetPropertyInventory)
			}

			damages := leaseId.Group("/damages/")
			{
				damages.POST("/", controllers.CreateDamage)
				damages.GET("/", controllers.GetDamagesByLease)

				damageId := damages.Group("/:damage_id/")
				{
					damageId.Use(middlewares.CheckDamageLeaseOwnership("damage_id"))
					damageId.GET("/", controllers.GetDamage)
					damageId.PUT("/", controllers.UpdateDamageTenant)
					damageId.POST("/pictures/", controllers.AddPicturesToDamage)
					damageId.PUT("/fix/", controllers.FixDamage)
				}
			}

			docs := leaseId.Group("/docs/")
			{
				docs.POST("/", controllers.UploadLeaseDocument)
				docs.GET("/", controllers.GetAllDocumentsByLease)

				// docId := docs.Group("/:doc_name/")
				// {
				// 	docId.GET("/", controllers.GetDocument)
				// }
			}

			reports := leaseId.Group("/inventory-reports/")
			{
				reports.GET("/", controllers.GetInventoryReportsByLease)
				reports.GET("/:report_id/",
					middlewares.CheckInventoryReportLeaseOwnership("report_id"),
					controllers.GetInventoryReport)
			}
		}
	}
}
