package com.example.immotep.apiClient

import com.example.immotep.apiCallerServices.AddPropertyInput
import com.example.immotep.apiCallerServices.AiCallInput
import com.example.immotep.apiCallerServices.AiCallOutput
import com.example.immotep.apiCallerServices.ArchivePropertyInput
import com.example.immotep.apiCallerServices.Document
import com.example.immotep.apiCallerServices.FurnitureInput
import com.example.immotep.apiCallerServices.FurnitureOutput
import com.example.immotep.apiCallerServices.GetPropertyResponse
import com.example.immotep.apiCallerServices.InventoryReportInput
import com.example.immotep.apiCallerServices.InviteInput
import com.example.immotep.apiCallerServices.InviteOutput
import com.example.immotep.apiCallerServices.ProfileResponse
import com.example.immotep.apiCallerServices.ProfileUpdateInput
import com.example.immotep.apiCallerServices.RoomOutput
import com.example.immotep.authService.LoginResponse
import com.example.immotep.authService.RegistrationInput
import com.example.immotep.authService.RegistrationResponse
import com.example.immotep.inventory.InventoryReportOutput
import retrofit2.http.Body
import retrofit2.http.Field
import retrofit2.http.FormUrlEncoded
import retrofit2.http.GET
import retrofit2.http.Header
import retrofit2.http.POST
import retrofit2.http.PUT
import retrofit2.http.Path



data class AddRoomInput(
    val name : String,
)

const val API_PREFIX = "/api/v1"

interface ApiService {

    //Login functions
    @FormUrlEncoded
    @POST("${API_PREFIX}/auth/token")
    suspend fun login(
        @Field("grant_type") grantType: String = "password",
        @Field("username") username: String,
        @Field("password") password: String,
    ): LoginResponse

    @FormUrlEncoded
    @POST("${API_PREFIX}/auth/token")
    suspend fun refreshToken(
        @Field("grant_type") grantType: String = "refresh_token",
        @Field("refresh_token") refreshToken: String,
    ): LoginResponse

    @POST("${API_PREFIX}/auth/register")
    suspend fun register(
        @Body registrationInput: RegistrationInput,
    ): RegistrationResponse

    //profile functions
    @GET("${API_PREFIX}/profile")
    suspend fun getProfile(@Header("Authorization") authHeader : String): ProfileResponse

    @PUT("${API_PREFIX}/profile")
    suspend fun updateProfile(@Header("Authorization") authHeader : String, @Body profileUpdateInput: ProfileUpdateInput): ProfileResponse

    //property functions
    @GET("${API_PREFIX}/owner/properties")
    suspend fun getProperties(@Header("Authorization") authHeader : String): Array<GetPropertyResponse>

    @GET("${API_PREFIX}/owner/properties/{propertyId}")
    suspend fun getProperty(@Header("Authorization") authHeader : String, @Path("propertyId") propertyId: String): GetPropertyResponse

    @GET("${API_PREFIX}/owner/properties/{propertyId}/documents")
    suspend fun getPropertyDocuments(@Header("Authorization") authHeader : String, @Path("propertyId") propertyId: String): Array<Document>

    @POST("${API_PREFIX}/owner/properties")
    suspend fun addProperty(@Header("Authorization") authHeader : String, @Body addPropertyInput: AddPropertyInput) : GetPropertyResponse

    @PUT("${API_PREFIX}/owner/properties/{propertyId}")
    suspend fun updateProperty(
        @Header("Authorization") authHeader : String,
        @Body addPropertyInput: AddPropertyInput,
        @Path("propertyId") propertyId: String
    ) : GetPropertyResponse

    @PUT("${API_PREFIX}/owner/properties/{propertyId}/archive")
    suspend fun archiveProperty(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Body archive : ArchivePropertyInput
    )

    //rooms functions
    @GET("${API_PREFIX}/owner/properties/{propertyId}/rooms")
    suspend fun getAllRooms(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
    ) : Array<RoomOutput>

    @POST("${API_PREFIX}/owner/properties/{propertyId}/rooms")
    suspend fun addRoom(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Body room: AddRoomInput
    ) : RoomOutput

    //furnitures functions
    @GET("${API_PREFIX}/owner/properties/{propertyId}/rooms/{roomId}/furnitures")
    suspend fun getAllFurnitures(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Path("roomId") roomId: String,
    ) : Array<FurnitureOutput>

    @POST("${API_PREFIX}/owner/properties/{propertyId}/rooms/{roomId}/furnitures")
    suspend fun addFurniture(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Path("roomId") roomId: String,
        @Body furniture: FurnitureInput
    ) : FurnitureOutput

    //inventory report functions
    @POST("${API_PREFIX}/owner/properties/{propertyId}/inventory-reports")
    suspend fun inventoryReport(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Body inventoryReportInput: InventoryReportInput
    )

    @GET("${API_PREFIX}/owner/properties/{propertyId}/inventory-reports")
    suspend fun getAllInventoryReports(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
    ) : Array<InventoryReportOutput>

    @GET("${API_PREFIX}/owner/properties/{propertyId}/inventory-reports/{report_id}")
    suspend fun getInventoryReportByIdOrLatest(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Path("report_id") reportId: String,
    ) : InventoryReportOutput

    //ia functions
    @POST("${API_PREFIX}/owner/properties/{propertyId}/inventory-reports/summarize/")
    suspend fun aiSummarize(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Body summarizeInput: AiCallInput
    ) : AiCallOutput

    @POST("${API_PREFIX}/owner/properties/{propertyId}/inventory-reports/compare/{old_report_id}")
    suspend fun aiCompare(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Path("old_report_id") oldReportId: String,
        @Body summarizeInput: AiCallInput
    ) : AiCallOutput

    //tenant functions
    @POST("${API_PREFIX}/owner/properties/{propertyId}/send-invite")
    suspend fun inviteTenant(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Body invite: InviteInput
    ) : InviteOutput
}
