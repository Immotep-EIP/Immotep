package com.example.immotep.apiCallerServices

import androidx.compose.ui.graphics.ImageBitmap
import androidx.navigation.NavController
import com.example.immotep.apiClient.ApiService
import com.example.immotep.apiClient.CreateOrUpdateResponse
import com.example.immotep.utils.DateFormatter
import retrofit2.HttpException
import java.time.OffsetDateTime

//enum classes

enum class PropertyStatus {
    unavailable,
    available,
    invite_sent
}

fun stringToPropertyState(str : String) : PropertyStatus {
    return when(str) {
        "available" -> PropertyStatus.available
        "invite sent" -> PropertyStatus.invite_sent
        else -> PropertyStatus.unavailable
    }
}

enum class DamageStatus {
    PENDING,
    PLANNED,
    AWAITING_OWNER_CONFIRMATION,
    AWAITING_TENANT_CONFIRMATION,
    FIXED
}

fun stringToDamageStatus(str : String) : DamageStatus {
    return when(str) {
        "pending" -> DamageStatus.PENDING
        "planned" -> DamageStatus.PLANNED
        "awaiting_owner_confirmation" -> DamageStatus.AWAITING_OWNER_CONFIRMATION
        "awaiting_tenant_confirmation" -> DamageStatus.AWAITING_TENANT_CONFIRMATION
        else -> DamageStatus.FIXED
    }
}

enum class DamagePriority {
    LOW,
    MEDIUM,
    HIGH,
    URGENT
}

fun stringToDamagePriority(str : String) : DamagePriority {
    return when (str) {
        "urgent" -> DamagePriority.URGENT
        "medium" -> DamagePriority.MEDIUM
        "high" -> DamagePriority.HIGH
        else -> DamagePriority.LOW
    }
}

//input and output api classes


//input api classes

data class AddPropertyInput(
    val name: String = "",
    val address: String = "",
    val city: String = "",
    val postal_code: String ="",
    val country: String = "",
    val area_sqm: Double = 0.0,
    val rental_price_per_month: Int = 0,
    val deposit_price: Int = 0,
    val apartment_number: String = ""
) {
    fun toDetailedProperty(
        id : String,
        currentLease : LeaseDetailedProperty? = null,
        currentInvite : InviteDetailedProperty? = null
    ) : DetailedProperty {
        return DetailedProperty(
            id = id,
            address = address,
            status  = PropertyStatus.available,
            appartementNumber = apartment_number,
            area = area_sqm.toInt(),
            rent= rental_price_per_month,
            deposit = deposit_price,
            zipCode = postal_code,
            city = city,
            country = country,
            name = name,
            lease = currentLease,
            invite = currentInvite
        )
    }
}

data class DocumentInput(
    val name: String = "",
    val data: String = ""
) {
    fun toDocument(id : String) : Document {
        return Document(
            id = id,
            name = name,
            data = data,
            created_at = DateFormatter.currentDateAsOffsetDateTimeString()
        )
    }
}


data class UpdatePropertyPictureInput(
    val data : String
)

//output api classes

data class InvitePropertyResponse(
    val end_date: String,
    val start_date: String,
    val tenant_email: String
) {
    fun toInviteDetailedProperty() = InviteDetailedProperty(
        startDate = OffsetDateTime.parse(this.start_date),
        endDate = OffsetDateTime.parse(this.end_date),
        tenantEmail = this.tenant_email
    )
}

data class PropertyPictureResponse(
    val id: String,
    val created_at: String,
    val data: String,
)

data class LeasePropertyResponse(
    val active: Boolean,
    val end_date: String,
    val id: String,
    val start_date: String,
    val tenant_email: String,
    val tenant_name: String
) {
    fun toLeaseDetailedProperty() = LeaseDetailedProperty(
        id = this.id,
        startDate = OffsetDateTime.parse(this.start_date),
        endDate = OffsetDateTime.parse(this.end_date),
        tenantEmail = this.tenant_email,
        tenantName = this.tenant_name
    )
}


data class GetPropertyResponse(
    val id: String,
    val apartment_number: String?,
    val archived: Boolean,
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
    val picture_id: String?,
    val invite: InvitePropertyResponse?,
    val lease: LeasePropertyResponse?
) {
    fun toDetailedProperty() : DetailedProperty  {
        val statusAsPropertyStatus = stringToPropertyState(this.status)
        var currentLease : LeaseDetailedProperty? = null
        var currentInvite : InviteDetailedProperty? = null
        if (this.lease != null && this.lease.active && statusAsPropertyStatus == PropertyStatus.unavailable) {
            currentLease = this.lease.toLeaseDetailedProperty()
        }
        if (this.invite != null && statusAsPropertyStatus == PropertyStatus.invite_sent) {
            currentInvite = this.invite.toInviteDetailedProperty()
        }
        return DetailedProperty(
            id = this.id,
            name = this.name,
            zipCode = this.postal_code,
            city = this.city,
            country = this.country,
            address = this.address,
            appartementNumber = this.apartment_number,
            status = stringToPropertyState(this.status),
            area = this.area_sqm.toInt(),
            rent = this.rental_price_per_month,
            deposit = this.deposit_price,
            lease = currentLease,
            invite = currentInvite
        )
    }
}

//custom properties class

data class Document(
    val id: String,
    val name: String,
    val data: String,
    val created_at: String
)

data class InviteDetailedProperty(
    val startDate: OffsetDateTime,
    val endDate: OffsetDateTime,
    val tenantEmail: String
)

data class LeaseDetailedProperty(
    val id: String,
    val startDate: OffsetDateTime,
    val endDate: OffsetDateTime,
    val tenantEmail: String,
    val tenantName: String
)

data class DetailedProperty(
     val id : String = "",
     val address : String = "",
     val status: PropertyStatus = PropertyStatus.unavailable,
     val appartementNumber : String? = "",
     val area : Int = 0,
     val rent : Int = 0,
     val deposit : Int = 0,
     val zipCode : String = "",
     val city : String = "",
     val country : String = "",
     val name : String = "",
     val picture: ImageBitmap? = null,
     val invite : InviteDetailedProperty? = null,
     val lease : LeaseDetailedProperty? = null,

) {
    fun toAddPropertyInput() : AddPropertyInput {
        return AddPropertyInput(
            address = this.address,
            area_sqm = this.area.toDouble(),
            deposit_price = this.deposit,
            rental_price_per_month = this.rent,
            city = this.city,
            name = this.name,
            country = this.country,
            postal_code = this.zipCode,
            apartment_number = this.appartementNumber ?: ""
        )
    }

}

data class ArchivePropertyInput(
    val archive: Boolean
)


class RealPropertyCallerService (
    apiService: ApiService,
    navController: NavController,
) : ApiCallerService(apiService, navController) {

    suspend fun getPropertiesAsDetailedProperties():  Array<DetailedProperty> {
        return changeRetrofitExceptionByApiCallerException(logoutOnUnauthorized = true) {
            val properties = apiService.getProperties(getBearerToken())
            properties.map { it.toDetailedProperty() }.toTypedArray()
        }
    }

    suspend fun getPropertyPicture(propertyId: String): String? = changeRetrofitExceptionByApiCallerException {
        val response = apiService.getPropertyPicture(getBearerToken(), propertyId)

        when {
            response.code() == 204 -> null
            response.isSuccessful -> response.body()?.data
            else -> throw HttpException(response)
        }
    }

    suspend fun updatePropertyPicture(propertyId: String, propertyPicture: String): CreateOrUpdateResponse = changeRetrofitExceptionByApiCallerException {
        apiService.updatePropertyPicture(getBearerToken(), propertyId, UpdatePropertyPictureInput(data = propertyPicture))
    }

    suspend fun addProperty(property: AddPropertyInput): CreateOrUpdateResponse {
        return changeRetrofitExceptionByApiCallerException {
            apiService.addProperty(getBearerToken(), property)
        }
    }

    suspend fun archiveProperty(propertyId: String) {
        return changeRetrofitExceptionByApiCallerException {
            apiService.archiveProperty(getBearerToken(), propertyId, ArchivePropertyInput(true))
        }
    }

    suspend fun getPropertyWithDetails(propertyId: String? = null): DetailedProperty = changeRetrofitExceptionByApiCallerException {
        if (this.isOwner()) {
            if (propertyId == null) throw Exception("Property id is null")
            apiService.getProperty(getBearerToken(), propertyId).toDetailedProperty()
        } else {
            apiService.getPropertyTenant(
                getBearerToken(),
                "current"
            ).toDetailedProperty()
        }
    }

    suspend fun updateProperty(
        property: AddPropertyInput,
        propertyId: String,
    ): CreateOrUpdateResponse = changeRetrofitExceptionByApiCallerException {
        apiService.updateProperty(getBearerToken(), property, propertyId)
    }

    suspend fun getPropertyDocuments(propertyId: String, leaseId : String): Array<Document> = changeRetrofitExceptionByApiCallerException {
        if (this.isOwner()) {
            apiService.getPropertyDocuments(getBearerToken(), propertyId, leaseId)
        } else {
            apiService.getPropertyDocumentsTenant(getBearerToken(), "current")
        }
    }

    suspend fun uploadDocument(
        propertyId: String,
        leaseId : String,
        document: DocumentInput
    ): CreateOrUpdateResponse = changeRetrofitExceptionByApiCallerException {
        if (this.isOwner()) {
            apiService.uploadDocument(getBearerToken(), propertyId, leaseId, document)
        } else {
            apiService.uploadDocumentTenant(getBearerToken(), "current", document)
        }
    }

}