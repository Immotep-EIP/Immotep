package com.example.immotep.apiClient

import androidx.annotation.Nullable
import com.example.immotep.apiCallerServices.AddPropertyInput
import com.example.immotep.apiCallerServices.AddRoomInput
import com.example.immotep.apiCallerServices.AiCallInput
import com.example.immotep.apiCallerServices.AiCallOutput
import com.example.immotep.apiCallerServices.ArchivePropertyInput
import com.example.immotep.apiCallerServices.CreatedInventoryReport
import com.example.immotep.apiCallerServices.Document
import com.example.immotep.apiCallerServices.DocumentInput
import com.example.immotep.apiCallerServices.FurnitureInput
import com.example.immotep.apiCallerServices.FurnitureOutput
import com.example.immotep.apiCallerServices.GetPropertyResponse
import com.example.immotep.apiCallerServices.InventoryReportInput
import com.example.immotep.apiCallerServices.InviteInput
import com.example.immotep.apiCallerServices.InviteOutput
import com.example.immotep.apiCallerServices.ProfileResponse
import com.example.immotep.apiCallerServices.ProfileUpdateInput
import com.example.immotep.apiCallerServices.PropertyPictureResponse
import com.example.immotep.apiCallerServices.RoomOutput
import com.example.immotep.apiCallerServices.UpdatePropertyPictureInput
import com.example.immotep.authService.LoginResponse
import com.example.immotep.authService.RegistrationInput
import com.example.immotep.authService.RegistrationResponse
import com.example.immotep.inventory.InventoryReportOutput
import okhttp3.Response
import okhttp3.ResponseBody
import retrofit2.http.Body
import retrofit2.http.DELETE
import retrofit2.http.Field
import retrofit2.http.FormUrlEncoded
import retrofit2.http.GET
import retrofit2.http.Header
import retrofit2.http.POST
import retrofit2.http.PUT
import retrofit2.http.Path


data class CreateOrUpdateResponse(
    val id : String
)

const val API_PREFIX = "/v1"

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

    @GET("${API_PREFIX}/owner/properties/{propertyId}/picture/")
    suspend fun getPropertyPicture(@Header("Authorization") authHeader : String, @Path("propertyId") propertyId: String): retrofit2.Response<PropertyPictureResponse>

    @PUT("${API_PREFIX}/owner/properties/{propertyId}/picture/")
    suspend fun updatePropertyPicture(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Body picture: UpdatePropertyPictureInput
    ): CreateOrUpdateResponse

    @POST("${API_PREFIX}/owner/properties")
    suspend fun addProperty(@Header("Authorization") authHeader : String, @Body addPropertyInput: AddPropertyInput) : CreateOrUpdateResponse

    @PUT("${API_PREFIX}/owner/properties/{propertyId}")
    suspend fun updateProperty(
        @Header("Authorization") authHeader : String,
        @Body addPropertyInput: AddPropertyInput,
        @Path("propertyId") propertyId: String
    ) : CreateOrUpdateResponse

    @PUT("${API_PREFIX}/owner/properties/{propertyId}/archive")
    suspend fun archiveProperty(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Body archive : ArchivePropertyInput
    ) : CreateOrUpdateResponse


    @GET("${API_PREFIX}/owner/properties/{propertyId}/leases/{leaseId}/docs/")
    suspend fun getPropertyDocuments(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Path("leaseId") leaseId: String
    ): Array<Document>

    @POST("${API_PREFIX}/owner/properties/{propertyId}/leases/{leaseId}/docs/")
    suspend fun uploadDocument(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Path("leaseId") leaseId: String,
        @Body document: DocumentInput
    ): CreateOrUpdateResponse

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
    ) : CreateOrUpdateResponse

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
    ) : CreateOrUpdateResponse

    //inventory report functions
    @POST("${API_PREFIX}/owner/properties/{propertyId}/leases/{leaseId}/inventory-reports/")
    suspend fun inventoryReport(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Path("leaseId") leaseId: String,
        @Body inventoryReportInput: InventoryReportInput
    ) : CreatedInventoryReport

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
    @POST("${API_PREFIX}/owner/properties/{propertyId}/leases/{leaseId}/inventory-reports/summarize/")
    suspend fun aiSummarize(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Path("leaseId") leaseId: String,
        @Body summarizeInput: AiCallInput
    ) : AiCallOutput

    @POST("${API_PREFIX}/owner/properties/{propertyId}/leases/{leaseId}/inventory-reports/compare/{oldReportId}/")
    suspend fun aiCompare(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Path("leaseId") leaseId: String,
        @Path("oldReportId") oldReportId: String,
        @Body summarizeInput: AiCallInput
    ) : AiCallOutput

    //invitation tenant functions
    @DELETE("${API_PREFIX}/owner/properties/{propertyId}/cancel-invite")
    suspend fun cancelTenantInvitation(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
    ) : retrofit2.Response<Unit>

    @POST("${API_PREFIX}/owner/properties/{propertyId}/send-invite")
    suspend fun inviteTenant(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Body invite: InviteInput
    ) : CreateOrUpdateResponse

}
