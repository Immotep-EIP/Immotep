package com.example.immotep.apiClient.mockApi

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
import com.example.immotep.apiClient.AddRoomInput
import com.example.immotep.apiClient.ApiService
import com.example.immotep.authService.LoginResponse
import com.example.immotep.authService.RegistrationInput
import com.example.immotep.authService.RegistrationResponse
import com.example.immotep.inventory.Cleanliness
import com.example.immotep.inventory.InventoryReportOutput
import com.example.immotep.inventory.State




val parisFakeProperty = GetPropertyResponse(
    id = "parisFakeProperty",
    name = "parisFake",
    address = "19 rue de la paix",
    city = "Paris",
    country = "France",
    postal_code = "75000",
    created_at = "2025-03-09T13:52:54.823Z",
    apartment_number = "",
    archived = false,
    area_sqm = 45.0,
    owner_id = "test",
    rental_price_per_month = 2000,
    deposit_price = 2000,
    status = "Busy",
    nb_damage = 0,
    tenant = "test@gmail.com",
    start_date = "2025-03-09T13:52:54.823Z",
    end_date = "2026-03-09T13:52:54.823Z",
)

val marseilleFakeProperty = GetPropertyResponse(
    id = "marsFakeProperty",
    name = "marsFake",
    address = "1 rue de la companie des indes",
    city = "Marseille",
    country = "France",
    postal_code = "13000",
    created_at = "2025-03-09T13:52:54.823Z",
    apartment_number = "10",
    archived = false,
    area_sqm = 100.0,
    owner_id = "test",
    rental_price_per_month = 1000,
    deposit_price = 1000,
    status = "Busy",
    nb_damage = 0,
    tenant = "crashbandicoot@gmail.com",
    start_date = "2025-03-09T13:52:54.823Z",
    end_date = "2026-03-09T13:52:54.823Z"
)

val lyonFakeProperty = GetPropertyResponse(
    id = "lyonFakeProperty",
    name = "lyonFake",
    address = "30 rue de la source",
    city = "Lyon",
    country = "France",
    postal_code = "69000",
    created_at = "2025-03-09T13:52:54.823Z",
    apartment_number = "3",
    archived = false,
    area_sqm = 100.0,
    owner_id = "test",
    rental_price_per_month = 1000,
    deposit_price = 1000,
    status = "Busy",
    nb_damage = 0,
    tenant = "tomnook@gmail.com",
    start_date = "2025-03-09T13:52:54.823Z",
    end_date = "2026-03-09T13:52:54.823Z"
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
            id = "test123",
            email = "robin.denni@epitech.eu",
            firstname = "Test",
            lastname = "User",
            role = "test",
            created_at = "2025-03-09T13:52:54Z+0000",
            updated_at = "2025-03-09T13:52:54Z+0000"
        )
    }

    //property functions
    override suspend fun getProperties(authHeader : String): Array<GetPropertyResponse> {
        return arrayOf(parisFakeProperty, marseilleFakeProperty, lyonFakeProperty)
    }

    override suspend fun getProperty(authHeader : String, propertyId: String): GetPropertyResponse {
        return parisFakeProperty
    }

    override suspend fun addProperty(authHeader : String, addPropertyInput: AddPropertyInput) : GetPropertyResponse {
        return parisFakeProperty
    }

    override suspend fun getPropertyDocuments(
        authHeader: String,
        propertyId: String
    ): Array<Document> {
        TODO("Not yet implemented")
        return arrayOf()
    }
    override suspend fun updateProperty(
        authHeader : String,
        addPropertyInput: AddPropertyInput,
        propertyId: String
    ) : GetPropertyResponse {
        return parisFakeProperty
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

    override suspend fun updateProfile(authHeader : String, profileUpdateInput: ProfileUpdateInput) : ProfileResponse {
        return ProfileResponse(
            id = "test123",
            email = profileUpdateInput.email,
            firstname = profileUpdateInput.firstname,
            lastname = profileUpdateInput.lastname,
            role = "test",
            created_at = "2025-03-09T13:52:54.823Z",
            updated_at = "2025-03-09T13:52:54.823Z"
        )
    }

}
