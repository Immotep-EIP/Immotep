package com.example.immotep.apiCallerServices

import androidx.navigation.NavController
import com.example.immotep.apiClient.ApiService
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
)

//output api classes

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
    val start_date: String?,
    val end_date: String?
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

    suspend fun getPropertiesAsDetailedProperties(onError: () -> Unit):  Array<DetailedProperty> {
        try {
            val properties = apiService.getProperties(getBearerToken())
            return properties.map { it.toDetailedProperty() }.toTypedArray()
        } catch (e: Exception) {
            onError()
            throw e
        }
    }

    suspend fun addProperty(property: AddPropertyInput, onError: () -> Unit): DetailedProperty {
        try {
            return apiService.addProperty(getBearerToken(), property).toDetailedProperty()
        } catch (e: Exception) {
            onError()
            throw e
        }
    }

    suspend fun archiveProperty(propertyId: String, onError: () -> Unit) {
        try {
            apiService.archiveProperty(getBearerToken(), propertyId, ArchivePropertyInput(true))
        } catch (e: Exception) {
            onError()
            throw e
        }
    }

    suspend fun getPropertyWithDetails(propertyId: String, onError: () -> Unit): DetailedProperty {
        try {
            val propertyWithDetails = apiService.getProperty(getBearerToken(), propertyId).toDetailedProperty()
            return propertyWithDetails
        } catch (e: Exception) {
            onError()
            throw e
        }
    }

    suspend fun updateProperty(
        property: AddPropertyInput,
        propertyId: String,
        onError: () -> Unit
    ): DetailedProperty {
        try {
            return apiService.updateProperty(getBearerToken(), property, propertyId)
                .toDetailedProperty()
        } catch (e: Exception) {
            onError()
            throw e
        }
    }

    suspend fun getPropertyDocuments(propertyId: String, onError: () -> Unit): Array<Document> {
        try {
            val documents = apiService.getPropertyDocuments(getBearerToken(), propertyId)
            return documents
        } catch (e: Exception) {
            onError()
            throw e
        }
    }


}