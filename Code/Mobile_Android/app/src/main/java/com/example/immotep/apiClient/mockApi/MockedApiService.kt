package com.example.immotep.apiClient.mockApi

import com.example.immotep.apiCallerServices.AddPropertyInput
import com.example.immotep.apiCallerServices.AddPropertyResponse
import com.example.immotep.apiCallerServices.AiCallInput
import com.example.immotep.apiCallerServices.AiCallOutput
import com.example.immotep.apiCallerServices.ArchivePropertyInput
import com.example.immotep.apiCallerServices.FurnitureInput
import com.example.immotep.apiCallerServices.FurnitureOutput
import com.example.immotep.apiCallerServices.GetPropertyResponse
import com.example.immotep.apiCallerServices.InventoryReportInput
import com.example.immotep.apiCallerServices.InviteInput
import com.example.immotep.apiCallerServices.InviteOutput
import com.example.immotep.apiCallerServices.ProfileResponse
import com.example.immotep.apiCallerServices.RoomOutput
import com.example.immotep.apiClient.API_PREFIX
import com.example.immotep.apiClient.AddRoomInput
import com.example.immotep.apiClient.ApiService
import com.example.immotep.authService.LoginResponse
import com.example.immotep.authService.RegistrationInput
import com.example.immotep.authService.RegistrationResponse
import com.example.immotep.inventory.Cleanliness
import com.example.immotep.inventory.InventoryReportOutput
import com.example.immotep.inventory.State
import kotlinx.coroutines.flow.flowOf
import retrofit2.http.Body
import retrofit2.http.DELETE
import retrofit2.http.Field
import retrofit2.http.FormUrlEncoded
import retrofit2.http.GET
import retrofit2.http.Header
import retrofit2.http.POST
import retrofit2.http.PUT
import retrofit2.http.Path

val fakeProperty = GetPropertyResponse(
    id = "test",
    name = "test",
    address = "test",
    city = "test",
    country = "test",
    postal_code = "68100",
    created_at = "test",
    apartment_number = "test",
    archived = false,
    area_sqm = 100.0,
    owner_id = "test",
    rental_price_per_month = 1000,
    deposit_price = 1000,
    status = "available",
    nb_damage = 0,
    tenant = "",
    start_date = null,
    end_date = null
)

val fakeRoom = RoomOutput(
    id = "test",
    name = "test",
    property_id = "test",
)

val fakeFurniture = FurnitureOutput(
    id = "test",
    name = "test",
    room_id = "test",
    property_id = "test",
    quantity = 1
)

val fakeInventoryReport = InventoryReportOutput(
    date = "test",
    id = "test",
    property_id = "test",
    rooms = arrayOf(),
    type = "test",
)

val fakeAiCallOutput = AiCallOutput(
    cleanliness = Cleanliness.clean,
    note = "Test",
    state = State.good
)

class MockedApiService : ApiService {
    // Example: Simulate a successful login

    override suspend fun login(grantType: String, username: String, password: String): LoginResponse {
        if (username == "error@gmail.com" || password == "testError") {
            throw Exception("Unknown user,401")
        }
        return LoginResponse(
            access_token = "test",
            refresh_token = "test",
            token_type = "access",
            expires_in = 100000,
            properties = mapOf("test" to "test")
        )
    }

    override suspend fun refreshToken(grantType: String, refreshToken: String): LoginResponse {
        return LoginResponse(
            access_token = "test",
            refresh_token = "test",
            token_type = "access",
            expires_in = 100000,
            properties = mapOf("test" to "test")
        )
    }

    override suspend fun register(registrationInput: RegistrationInput): RegistrationResponse {
        return RegistrationResponse(
            id = "test",
            email = "test",
            firstname = "test",
            lastname = "test",
            role = "test",
            created_at = "test",
            updated_at = "test"
        )
    }

    //profile functions
    override suspend fun getProfile(authHeader : String): ProfileResponse {
        return ProfileResponse(
            id = "test",
            email = "test",
            firstname = "test",
            lastname = "test",
            role = "test",
            created_at = "test",
            updated_at = "test"
        )
    }

    //property functions
    override suspend fun getProperties(authHeader : String): Array<GetPropertyResponse> {
        return arrayOf(fakeProperty)
    }

    override suspend fun getProperty(authHeader : String, propertyId: String): GetPropertyResponse {
        return fakeProperty
    }

    override suspend fun addProperty(authHeader : String, addPropertyInput: AddPropertyInput) : AddPropertyResponse {
        return AddPropertyResponse(
            id = "test",
            owner_id = "test",
            name = "property2",
            address = "test",
            city = "testcity",
            postal_code = "68100",
            country = "testCountry",
            area_sqm =54.0,
            rental_price_per_month = 750,
            deposit_price = 750,
            picture = null,
            created_at = "test",
        )
    }

    override suspend fun updateProperty(
        authHeader : String,
        addPropertyInput: AddPropertyInput,
        propertyId: String
    ) : GetPropertyResponse {
        return fakeProperty
    }

    override suspend fun archiveProperty(authHeader : String, propertyId: String, archive: ArchivePropertyInput) {
        return
    }

    //rooms functions
    override suspend fun getAllRooms(
        authHeader : String,
        propertyId: String,
    ) : Array<RoomOutput> {
        return arrayOf(fakeRoom)
    }

    override suspend fun addRoom(
        authHeader : String,
        propertyId: String,
        room: AddRoomInput
    ) : RoomOutput {
        return fakeRoom
    }

    //furnitures functions
    override suspend fun getAllFurnitures(
        authHeader : String,
        propertyId: String,
        roomId: String,
    ) : Array<FurnitureOutput> {
        return arrayOf(fakeFurniture)
    }

    override suspend fun addFurniture(
        authHeader : String,
        propertyId: String,
        roomId: String,
        furniture: FurnitureInput
    ) : FurnitureOutput {
        return fakeFurniture
    }

    //inventory report functions

    override suspend fun inventoryReport(
        authHeader : String,
        propertyId: String,
        inventoryReportInput: InventoryReportInput
    ) {

    }

    override suspend fun getAllInventoryReports(
        authHeader : String,
        propertyId: String,
    ) : Array<InventoryReportOutput> {
        return arrayOf(fakeInventoryReport)
    }

    override suspend fun getInventoryReportByIdOrLatest(
        authHeader : String,
        propertyId: String,
        reportId: String,
    ) : InventoryReportOutput {
        return fakeInventoryReport
    }

    //ia functions
    override suspend fun aiSummarize(
        authHeader : String,
        propertyId: String,
        summarizeInput: AiCallInput
    ) : AiCallOutput {
        return fakeAiCallOutput
    }

    override suspend fun aiCompare(
        authHeader : String,
        propertyId: String,
        oldReportId: String,
        summarizeInput: AiCallInput
    ) : AiCallOutput {
        return fakeAiCallOutput
    }

    //tenant functions
    override suspend fun inviteTenant(
        authHeader : String,
        propertyId: String,
        invite: InviteInput
    ) : InviteOutput {
        return InviteOutput(
            id = "test",
            property_id = "test",
            tenant_email = "test",
            start_date = "test",
            end_date = "test",
            created_at = "test"
        )
    }

}
