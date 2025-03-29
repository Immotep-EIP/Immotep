package router

import (
	"github.com/gin-gonic/gin"
	"immotep/backend/controllers"
	_ "immotep/backend/docs" // mandatory import for swagger doc
	"immotep/backend/router/middlewares"
)

func registerOwnerRoutes(owner *gin.RouterGroup) {
	owner.Use(middlewares.AuthorizeOwner())

	properties := owner.Group("/properties/")
	{
		properties.POST("/", controllers.CreateProperty)
		properties.GET("/", controllers.GetAllPropertiesByOwner)
		properties.GET("/archived/", controllers.GetArchivedPropertiesByOwner)

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

			registerOwnerInventoryRoutes(propertyId)

			invReports := propertyId.Group("/inventory-reports/")
			registerOwnerInvReportRoutes(invReports)

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

		docs := leaseId.Group("/docs/")
		{
			docs.POST("/", controllers.UploadDocument)
			docs.GET("/", controllers.GetAllDocumentsByLease)

			docId := docs.Group("/:doc_id/")
			{
				docId.Use(middlewares.CheckDocumentLeaseOwnership("doc_id"))
				docId.GET("/", controllers.GetDocument)
				// document.DELETE("/:doc_id/", controllers.DeleteDocument)
			}
		}
	}
}

func registerOwnerInventoryRoutes(propertyId *gin.RouterGroup) {
	propertyId.GET("/inventory/", controllers.GetPropertyInventory)

	// TODO: move to inventory group
	rooms := propertyId.Group("/rooms/")
	{
		rooms.POST("/", controllers.CreateRoom)
		rooms.GET("/", controllers.GetAllRoomsByProperty)
		rooms.GET("/archived/", controllers.GetArchivedRoomsByProperty)

		roomId := rooms.Group("/:room_id/")
		{
			roomId.Use(middlewares.CheckRoomPropertyOwnership("property_id", "room_id"))
			roomId.GET("/", controllers.GetRoom)
			roomId.PUT("/archive/", controllers.ArchiveRoom)

			furnitures := roomId.Group("/furnitures/")
			{
				furnitures.POST("/", controllers.CreateFurniture)
				furnitures.GET("/", controllers.GetAllFurnituresByRoom)
				furnitures.GET("/archived/", controllers.GetArchivedFurnituresByRoom)

				furnitureId := furnitures.Group("/:furniture_id/")
				{
					furnitureId.Use(middlewares.CheckFurnitureRoomOwnership("room_id", "furniture_id"))
					furnitureId.GET("/", controllers.GetFurniture)
					furnitureId.PUT("/archive/", controllers.ArchiveFurniture)
				}
			}
		}
	}
}

func registerOwnerInvReportRoutes(invReports *gin.RouterGroup) {
	invReports.POST("/", controllers.CreateInventoryReport)
	invReports.GET("/", controllers.GetAllInventoryReportsByProperty)

	invReports.GET("/:report_id/",
		middlewares.CheckInventoryReportPropertyOwnership("property_id", "report_id"),
		controllers.GetInventoryReport)

	invReports.POST("/summarize/", controllers.GenerateSummary)
	invReports.POST("/compare/:old_report_id/",
		middlewares.CheckInventoryReportPropertyOwnership("property_id", "old_report_id"),
		controllers.GenerateComparison)
}
