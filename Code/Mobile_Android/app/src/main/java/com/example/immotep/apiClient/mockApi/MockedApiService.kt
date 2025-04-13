package com.example.immotep.apiClient.mockApi

import androidx.annotation.Nullable
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
import okhttp3.Response
import okhttp3.ResponseBody


class MockedApiService : ApiService {
    // Example: Simulate a successful login

    override suspend fun login(grantType: String, username: String, password: String): LoginResponse {
        if (username == "error@gmail.com" || password == "testError") {
            throw Exception("Unknown user,401")
        }
        return fakeLoginResponse
    }

    override suspend fun refreshToken(grantType: String, refreshToken: String): LoginResponse {
        return fakeLoginResponse
    }

    override suspend fun register(registrationInput: RegistrationInput): RegistrationResponse {
        return fakeRegistrationResponse
    }

    //profile functions
    override suspend fun getProfile(authHeader : String): ProfileResponse {
        return fakeProfileResponse
    }

    //property functions
    override suspend fun getProperties(authHeader : String): Array<GetPropertyResponse> {
        return arrayOf(parisFakeProperty, marseilleFakeProperty, lyonFakeProperty, emptyFakeProperty)
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
        return arrayOf(fakeDocument)
    }
    override suspend fun updateProperty(
        authHeader : String,
        addPropertyInput: AddPropertyInput,
        propertyId: String
    ) : GetPropertyResponse {
        val modifiedFakeProperty = GetPropertyResponse(
            id = parisFakeProperty.id,
            name = addPropertyInput.name,
            address = addPropertyInput.address,
            city = addPropertyInput.city,
            postal_code = addPropertyInput.postal_code,
            country = addPropertyInput.country,
            area_sqm = addPropertyInput.area_sqm,
            deposit_price = addPropertyInput.deposit_price,
            rental_price_per_month = addPropertyInput.rental_price_per_month,
            status = parisFakeProperty.status,
            created_at = parisFakeProperty.created_at,
            apartment_number = addPropertyInput.apartment_number,
            archived = parisFakeProperty.archived,
            owner_id = parisFakeProperty.owner_id,
            nb_damage = parisFakeProperty.nb_damage,
            tenant = parisFakeProperty.tenant,
            start_date = parisFakeProperty.start_date,
            end_date = parisFakeProperty.end_date
        )
        return modifiedFakeProperty
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

    //invite tenant functions
    override suspend fun inviteTenant(
        authHeader : String,
        propertyId: String,
        invite: InviteInput
    ) : InviteOutput {
        return fakeInviteOutput
    }
    //tenant functions
    override suspend fun cancelTenantInvitation(authHeader: String, propertyId: String): retrofit2.Response<Unit> {
        return retrofit2.Response.success(Unit)
    }

    override suspend fun updateProfile(authHeader : String, profileUpdateInput: ProfileUpdateInput) : ProfileResponse {
        if (profileUpdateInput.email == "error@gmail.com") {
            throw Exception("Unknown user,401")
        }
        return ProfileResponse(
            id = "test123",
            email = profileUpdateInput.email,
            firstname = profileUpdateInput.firstname,
            lastname = profileUpdateInput.lastname,
            role = "test",
            created_at = baseDateStr,
            updated_at = baseDateStr
        )
    }

}
