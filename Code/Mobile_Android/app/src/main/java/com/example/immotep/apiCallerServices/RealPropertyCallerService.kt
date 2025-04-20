package com.example.immotep.apiCallerServices

import androidx.navigation.NavController
import com.example.immotep.apiClient.ApiService
import com.example.immotep.apiClient.CreateOrUpdateResponse
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
    fun toDetailedProperty(id : String) : DetailedProperty {
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
            name = name
        )
    }
}

//output api classes

data class InvitePropertyResponse(
    val end_date: String,
    val start_date: String,
    val tenant_email: String
)

data class LeasePropertyResponse(
    val active: Boolean,
    val end_date: String,
    val id: String,
    val start_date: String,
    val tenant_email: String,
    val tenant_name: String
)


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
    val tenant: String,
    val picture_id: String?,
    val start_date: String?,
    val end_date: String?,
    val invite: InvitePropertyResponse,
    val lease: LeasePropertyResponse
) {

    fun toDetailedProperty() = DetailedProperty(
        id = this.id, 
        image = "", 
        name = this.name,
        zipCode = this.postal_code,
        city = this.city,
        country = this.country,
        address = this.address,
        appartementNumber = this.apartment_number,
        tenant = this.tenant,
        status = stringToPropertyState(this.status),
        startDate = if (this.start_date != null) OffsetDateTime.parse(this.start_date) else null,
        endDate = if (this.end_date != null) OffsetDateTime.parse(this.end_date) else null,
        area = this.area_sqm.toInt(),
        rent = this.rental_price_per_month,
        deposit = this.deposit_price,
        currentLeaseId = lease.id,
        documents = arrayOf()
    )
}

//custom properties class

data class Document(
    val id: String,
    val name: String,
    val data: String,
    val created_at: String
)

data class DetailedProperty(
     val id : String = "",
     val image : String = "",
     val address : String = "",
     val currentLeaseId: String = "",
     val tenant : String? = null,
     val status: PropertyStatus = PropertyStatus.unavailable,
     val startDate : OffsetDateTime? = null,
     val endDate : OffsetDateTime? = null,
     val appartementNumber : String? = "",
     val area : Int = 0,
     val rent : Int = 0,
     val deposit : Int = 0,
     val documents : Array<Document> = arrayOf(),
     val zipCode : String = "",
     val city : String = "",
     val country : String = "",
     val name : String = ""
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
        return changeRetrofitExceptionByApiCallerException {
            val properties = apiService.getProperties(getBearerToken())
            properties.map { it.toDetailedProperty() }.toTypedArray()
        }
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

    suspend fun getPropertyWithDetails(propertyId: String): DetailedProperty = changeRetrofitExceptionByApiCallerException {
        apiService.getProperty(getBearerToken(), propertyId).toDetailedProperty()
    }

    suspend fun updateProperty(
        property: AddPropertyInput,
        propertyId: String,
    ): CreateOrUpdateResponse = changeRetrofitExceptionByApiCallerException {
        apiService.updateProperty(getBearerToken(), property, propertyId)
    }

    suspend fun getPropertyDocuments(propertyId: String, leaseId : String): Array<Document> = changeRetrofitExceptionByApiCallerException {
        apiService.getPropertyDocuments(getBearerToken(), propertyId, leaseId)
    }


}