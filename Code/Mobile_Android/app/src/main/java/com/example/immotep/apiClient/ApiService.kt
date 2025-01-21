package com.example.immotep.apiClient

import retrofit2.http.Body
import retrofit2.http.Field
import retrofit2.http.FormUrlEncoded
import retrofit2.http.GET
import retrofit2.http.Header
import retrofit2.http.POST
import retrofit2.http.Path
import java.util.Vector

//enums classes

enum class Cleanliness {
    not_set,
    dirty,
    medium,
    clean,
}

enum class State {
    not_set,
    broken,
    needsRepair,
    bad,
    medium,
    good,
    new
}

enum class InventoryLocationsTypes {
    room,
    furniture
}

//input and output classes

data class LoginResponse(
    val access_token: String,
    val refresh_token: String,
    val token_type: String,
    val expires_in: Int,
    val properties: Map<String, Any>,
)

data class RegistrationInput(
    val email: String,
    val password: String,
    val firstName: String,
    val lastName: String,
)

data class RegistrationResponse(
    val id: String,
    val email: String,
    val firstname: String,
    val lastname: String,
    val role: String,
    val created_at: String,
    val updated_at: String,
)

data class ProfileResponse(
    val id: String,
    val email: String,
    val firstname: String,
    val lastname: String,
    val role: String,
    val created_at: String,
    val updated_at: String,
)

data class AddPropertyInput(
    val name: String = "",
    val address: String = "",
    val city: String = "",
    val postal_code: String ="",
    val country: String = "",
    val area_sqm: Double = 0.0,
    val rental_price_per_month: Int = 0,
    val deposit_price: Int = 0,
)

data class AddPropertyResponse(
    val id: String,
    val owner_id: String,
    val name: String,
    val address: String,
    val city: String,
    val postal_code: String,
    val country: String,
    val area_sqm: Double,
    val rental_price_per_month: Int,
    val deposit_price: Int,
    val picture: String?,
    val created_at: String,
)

data class GetPropertyResponse(
    val id: String,
    val owner_id: String,
    val name: String,
    val address: String,
    val city: String,
    val postal_code: String,
    val country: String,
    val area_sqm: Double,
    val rental_price_per_month: Int,
    val deposit_price: Int,
    val created_at: String,
    val status: String,
    val nb_damage: Int,
    val tenant: String,
    val start_date: String?,
    val end_date: String?
)

data class FurnitureOutput(
    val id: String,
    val property_id: String,
    val room_id: String,
    val name: String,
    val quantity: Int
)

data class FurnitureInput(
    val name: String,
    val quantity: Int
)

data class InventoryReportFurniture(
    val id: String,
    val cleanliness: Cleanliness,
    val note: String,
    val pictures: Vector<String>,
    val state: State,
)

data class InventoryReportRoom(
    val id: String,
    val cleanliness: Cleanliness,
    val state: State,
    val note: String,
    val pictures: Vector<String>,
    val furnitures: Vector<InventoryReportFurniture>
)

data class InventoryReportInput(
    val type: String,
    val rooms: Vector<InventoryReportRoom>
)

data class AddRoomInput(
    val name : String,
)

data class RoomOutput(
    val id : String,
    val name : String,
    val property_id : String,
)

//ai input data classes

data class SummarizeInput(
    val id : String,
    val pictures : Vector<String>,
    val type : InventoryLocationsTypes
)

//ai output data classes

data class SummarizeOutput(
    val cleanliness: Cleanliness,
    val note: String,
    val state: State
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

    //property functions
    @GET("${API_PREFIX}/owner/properties")
    suspend fun getProperties(@Header("Authorization") authHeader : String): Array<GetPropertyResponse>

    @GET("${API_PREFIX}/owner/properties/{propertyId}")
    suspend fun getProperty(@Header("Authorization") authHeader : String, @Path("propertyId") propertyId: String): GetPropertyResponse

    @POST("${API_PREFIX}/owner/properties")
    suspend fun addProperty(@Header("Authorization") authHeader : String, @Body addPropertyInput: AddPropertyInput) : AddPropertyResponse


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

    //ia functions
    @GET("${API_PREFIX}/owner/properties/{propertyId}/inventory-reports/summarize/")
    suspend fun aiSummarize(
        @Header("Authorization") authHeader : String,
        @Path("propertyId") propertyId: String,
        @Body summarizeInput: SummarizeInput
    ) : SummarizeOutput
}
