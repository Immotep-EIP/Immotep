package fr.keyz.apiClient.mockApi

import fr.keyz.apiCallerServices.AddPropertyInput
import fr.keyz.apiCallerServices.AddRoomInput
import fr.keyz.apiCallerServices.AiCallInput
import fr.keyz.apiCallerServices.AiCallOutput
import fr.keyz.apiCallerServices.ArchivePropertyInput
import fr.keyz.apiCallerServices.CreatedInventoryReport
import fr.keyz.apiCallerServices.DamageInput
import fr.keyz.apiCallerServices.DamageOutput
import fr.keyz.apiCallerServices.Document
import fr.keyz.apiCallerServices.DocumentInput
import fr.keyz.apiCallerServices.FurnitureInput
import fr.keyz.apiCallerServices.FurnitureOutput
import fr.keyz.apiCallerServices.GetDashBoardOutput
import fr.keyz.apiCallerServices.GetPropertyResponse
import fr.keyz.apiCallerServices.InventoryReportInput
import fr.keyz.apiCallerServices.InviteInput
import fr.keyz.apiCallerServices.ProfileResponse
import fr.keyz.apiCallerServices.ProfileUpdateInput
import fr.keyz.apiCallerServices.PropertyPictureResponse
import fr.keyz.apiCallerServices.RoomOutput
import fr.keyz.apiCallerServices.UpdatePropertyPictureInput
import fr.keyz.apiClient.ApiService
import fr.keyz.apiClient.ArchiveInput
import fr.keyz.apiClient.CreateOrUpdateResponse
import fr.keyz.authService.LoginResponse
import fr.keyz.authService.RegistrationInput
import fr.keyz.authService.RegistrationResponse
import fr.keyz.inventory.InventoryReportOutput


class MockedApiService : ApiService {
    // Example: Simulate a successful login

    override suspend fun login(grantType: String, username: String, password: String): LoginResponse {
        if (username == "error@gmail.com" || password == "testError") {
            throw Exception("Unknown user,401")
        }
        if (username == "tenant@gmail.com") {
            return fakeLoginResponse.copy(
                access_token = "tenantAccessToken",
            )
        }
        return fakeLoginResponse
    }

    override suspend fun loggin(grantType: String, username: String, password: String): String {
        TODO("Not yet implemented")
    }

    override suspend fun refreshToken(grantType: String, refreshToken: String): LoginResponse {
        return fakeLoginResponse
    }

    override suspend fun register(registrationInput: RegistrationInput): RegistrationResponse {
        return fakeRegistrationResponse
    }

    //profile functions
    override suspend fun getProfile(authHeader : String): ProfileResponse {
        if (authHeader == "Bearer tenantAccessToken") {
            return fakeProfileResponse.copy(
                role = "tenant"
            )
        }
        return fakeProfileResponse
    }

    //property functions
    override suspend fun getProperties(authHeader : String): Array<GetPropertyResponse> {
        return arrayOf(parisFakeProperty, marseilleFakeProperty, lyonFakeProperty, emptyFakeProperty)
    }

    override suspend fun getProperty(authHeader : String, propertyId: String): GetPropertyResponse {
        return parisFakeProperty
    }

    override suspend fun addProperty(authHeader : String, addPropertyInput: AddPropertyInput) : CreateOrUpdateResponse {
        return CreateOrUpdateResponse(
            id = parisFakeProperty.id
        )
    }

    override suspend fun getPropertyDocuments(
        authHeader: String,
        propertyId: String,
        leaseId: String
    ): Array<Document> {
        return arrayOf(fakeDocument)
    }

    override suspend fun getPropertyDocumentsTenant(
        authHeader: String,
        leaseId: String
    ): Array<Document> {
        return arrayOf(fakeDocument)
    }

    override suspend fun updateProperty(
        authHeader : String,
        addPropertyInput: AddPropertyInput,
        propertyId: String
    ) : CreateOrUpdateResponse {
        return CreateOrUpdateResponse(parisFakeProperty.id)
    }

    override suspend fun archiveProperty(authHeader : String, propertyId: String, archive: ArchivePropertyInput) : CreateOrUpdateResponse {
        return CreateOrUpdateResponse(
            id = propertyId
        )
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
    ) : CreateOrUpdateResponse {
        return CreateOrUpdateResponse(
            id = fakeRoom.id
        )
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
    ) : CreateOrUpdateResponse {
        return CreateOrUpdateResponse(
            id = "newFurnitureId"
        )
    }

    //inventory report functions

    override suspend fun inventoryReport(
        authHeader : String,
        propertyId: String,
        leaseId: String,
        inventoryReportInput: InventoryReportInput
    )  : CreatedInventoryReport {
        return CreatedInventoryReport(
            date = baseDateStr,
            errors = arrayOf(),
            id = "newInventoryReport",
            lease_id = leaseId,
            pdf_data = fakeDocument.data,
            pdf_name = fakeDocument.name,
            property_id = propertyId,
            type = inventoryReportInput.type
        )
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
        leaseId: String,
        summarizeInput: AiCallInput
    ) : AiCallOutput {
        return fakeAiCallOutput
    }

    override suspend fun aiCompare(
        authHeader : String,
        propertyId: String,
        leaseId: String,
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
    ) : CreateOrUpdateResponse {
        return CreateOrUpdateResponse(
            id = fakeInviteOutput.id
        )
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

    override suspend fun uploadDocument(
        authHeader : String,
        propertyId: String,
        leaseId: String,
        document: DocumentInput
    ): CreateOrUpdateResponse {
        return CreateOrUpdateResponse(
            id = "newDocumentId"
        )
    }

    override suspend fun uploadDocumentTenant(
        authHeader : String,
        leaseId: String,
        document: DocumentInput
    ): CreateOrUpdateResponse {
        return CreateOrUpdateResponse(
            id = "newDocumentId"
        )
    }

    override suspend fun getPropertyPicture(
        authHeader: String,
        propertyId: String
    ): retrofit2.Response<PropertyPictureResponse> {
        return retrofit2.Response.success(
            PropertyPictureResponse(
            id = "pictureId",
            created_at = baseDateStr,
            data = "",
        )
        )
    }

    override suspend fun updatePropertyPicture(
        authHeader: String,
        propertyId: String,
        picture: UpdatePropertyPictureInput
    ): CreateOrUpdateResponse {
        return CreateOrUpdateResponse(
            id = "pictureId"
        )
    }

    override suspend fun getPropertyTenant(authHeader: String, leaseId: String): GetPropertyResponse {
        return parisFakeProperty
    }

    override suspend fun getPropertyDamages(
        authHeader : String,
        propertyId: String,
        leaseId: String
    ): Array<DamageOutput> {
        return fakeDamagesArray
    }

    override suspend fun getPropertyDamagesTenant(
        authHeader : String,
        leaseId: String
    ): Array<DamageOutput> {
        return fakeDamagesArray
    }

    override suspend fun getAllRoomsTenant(authHeader: String, leaseId: String): Array<RoomOutput> {
        return arrayOf(fakeRoom)
    }

    override suspend fun addDamage(
        authHeader: String,
        leaseId: String,
        damage: DamageInput
    ): CreateOrUpdateResponse {
        return CreateOrUpdateResponse(id = "newDamage")
    }

    override suspend fun archiveRoom(
        authHeader: String,
        propertyId: String,
        roomId: String,
        archive: ArchiveInput
    ): CreateOrUpdateResponse {
        return CreateOrUpdateResponse(roomId)
    }

    override suspend fun getDashboard(authHeader: String, lang: String): GetDashBoardOutput {
        return fakeGetDashBoardOutput
    }

    override suspend fun getPropertyDamage(
        authHeader: String,
        propertyId: String,
        leaseId: String,
        damageId: String
    ): DamageOutput {
        return fakeDamageOutput
    }

    override suspend fun getPropertyDamageTenant(
        authHeader: String,
        leaseId: String,
        damageId: String
    ): DamageOutput {
        return fakeDamageOutput
    }
}
