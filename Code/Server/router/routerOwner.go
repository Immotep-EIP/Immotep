package router

import (
	"github.com/gin-gonic/gin"
	"immotep/backend/controllers"
	_ "immotep/backend/docs" // mandatory import for swagger doc
	"immotep/backend/router/middlewares"
)

func registerOwnerRoutes(owner *gin.RouterGroup) {
	owner.Use(middlewares.AuthorizeOwner())

	owner.GET("/dashboard/", controllers.GetOwnerDashboard)

	properties := owner.Group("/properties/")
	{
		properties.POST("/", controllers.CreateProperty)
		properties.GET("/", controllers.GetPropertiesByOwner)

		propertyId := properties.Group("/:property_id/")
		{
			propertyId.Use(middlewares.CheckPropertyOwnerOwnership("property_id"))
			propertyId.GET("/", controllers.GetProperty)
			propertyId.PUT("/", controllers.UpdateProperty)
			propertyId.PUT("/archive/", controllers.ArchiveProperty)
			propertyId.GET("/picture/", controllers.GetPropertyPicture)
			propertyId.PUT("/picture/", controllers.UpdatePropertyPicture)

			// TODO: move to lease routes
			propertyId.POST("/send-invite/", controllers.InviteTenant)
			propertyId.DELETE("/cancel-invite/", middlewares.CheckLeaseInvite("property_id"), controllers.CancelInvite)

			propertyId.GET("/damages/", controllers.GetDamagesByProperty)

			reports := propertyId.Group("/inventory-reports/")
			{
				reports.GET("/", controllers.GetAllInventoryReportsByProperty)
				reports.GET("/:report_id/",
					middlewares.CheckInventoryReportPropertyOwnership("report_id"),
					controllers.GetInventoryReport)
			}

			registerOwnerInventoryRoutes(propertyId)

			leases := propertyId.Group("/leases/")
			registerOwnerLeaseRoutes(leases)
		}
	}
}

func registerOwnerLeaseRoutes(leases *gin.RouterGroup) {
	leases.GET("/", controllers.GetAllLeasesByProperty)

	leaseId := leases.Group("/:lease_id/")
	{
		leaseId.Use(middlewares.CheckLeasePropertyOwnership("property_id", "lease_id"))
		leaseId.GET("/", controllers.GetLease)
		leaseId.PUT("/end/", controllers.EndLease)

		damages := leaseId.Group("/damages/")
		{
			damages.GET("/", controllers.GetDamagesByLease)

			damageId := damages.Group("/:damage_id/")
			{
				damageId.Use(middlewares.CheckDamageLeaseOwnership("damage_id"))
				damageId.GET("/", controllers.GetDamage)
				damageId.PUT("/", controllers.UpdateDamageOwner)
				damageId.PUT("/fix/", controllers.FixDamage)
			}
		}

		docs := leaseId.Group("/docs/")
		{
			docs.POST("/", controllers.UploadDocument)
			docs.GET("/", controllers.GetAllDocumentsByLease)

			docId := docs.Group("/:doc_id/")
			{
				docId.Use(middlewares.CheckDocumentLeaseOwnership("doc_id"))
				docId.GET("/", controllers.GetDocument)
				docId.DELETE("/", controllers.DeleteDocument)
			}
		}

		reports := leaseId.Group("/inventory-reports/")
		{
			reports.POST("/", controllers.CreateInventoryReport)
			reports.GET("/", controllers.GetInventoryReportsByLease)
			reports.GET("/:report_id/",
				middlewares.CheckInventoryReportLeaseOwnership("report_id"),
				controllers.GetInventoryReport)

			// AI
			reports.POST("/summarize/", controllers.GenerateSummary)
			reports.POST("/compare/:old_report_id/",
				middlewares.CheckInventoryReportLeaseOwnership("old_report_id"),
				controllers.GenerateComparison)
		}
	}
}

func registerOwnerInventoryRoutes(propertyId *gin.RouterGroup) {
	propertyId.GET("/inventory/", controllers.GetPropertyInventory)

	// TODO: move to inventory group
	rooms := propertyId.Group("/rooms/")
	{
		rooms.POST("/", controllers.CreateRoom)
		rooms.GET("/", controllers.GetRoomsByProperty)

		roomId := rooms.Group("/:room_id/")
		{
			roomId.Use(middlewares.CheckRoomPropertyOwnership("room_id"))
			roomId.GET("/", controllers.GetRoom)
			roomId.DELETE("/", controllers.DeleteRoom)
			roomId.PUT("/archive/", controllers.ArchiveRoom)

			furnitures := roomId.Group("/furnitures/")
			{
				furnitures.POST("/", controllers.CreateFurniture)
				furnitures.GET("/", controllers.GetFurnituresByRoom)

				furnitureId := furnitures.Group("/:furniture_id/")
				{
					furnitureId.Use(middlewares.CheckFurnitureRoomOwnership("room_id", "furniture_id"))
					furnitureId.GET("/", controllers.GetFurniture)
					furnitureId.DELETE("/", controllers.DeleteFurniture)
					furnitureId.PUT("/archive/", controllers.ArchiveFurniture)
				}
			}
		}
	}
}
