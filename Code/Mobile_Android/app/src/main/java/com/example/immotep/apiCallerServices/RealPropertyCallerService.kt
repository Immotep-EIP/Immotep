package com.example.immotep.apiCallerServices

import androidx.navigation.NavController
import com.example.immotep.apiClient.ApiService
import java.time.OffsetDateTime

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
    fun toProperty() = Property(
        id = this.id,
        image = "",
        address = this.address,
        tenant = this.tenant,
        available = this.status == "available",
        startDate = if (this.start_date != null) OffsetDateTime.parse(this.start_date) else null,
        endDate = if (this.end_date != null) OffsetDateTime.parse(this.end_date) else null
    )

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
        available = this.status == "available",
        startDate = if (this.start_date != null) OffsetDateTime.parse(this.start_date) else null,
        endDate = if (this.end_date != null) OffsetDateTime.parse(this.end_date) else null,
        area = this.area_sqm.toInt(),
        rent = this.rental_price_per_month,
        deposit = this.deposit_price,
        documents = arrayOf()
    )
}


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
) {
    fun toProperty() = Property(
        id = id,
        image = picture ?: "",
        address = address,
        tenant = null,
        available = true,
        startDate = null,
        endDate = null
    )
}

//custom properties class

interface IProperty {
    val id: String
    val image: String
    val address: String
    val tenant: String?
    val available: Boolean
    val startDate: OffsetDateTime?
    val endDate: OffsetDateTime?
}

data class Property(
    override val id: String = "",
    override val image: String = "",
    override val address: String = "",
    override val tenant: String? = null,
    override val available: Boolean = true,
    override val startDate: OffsetDateTime? = null,
    override val endDate: OffsetDateTime? = null
) : IProperty


data class Document(
    val id: String,
    val name: String,
    val data: String,
    val created_at: String
)

interface IDetailedProperty : IProperty {
    val area : Int
    val rent : Int
    val deposit : Int
    val documents : Array<Document>
    val zipCode : String
    val city : String
    val country : String
    val name : String
    val appartementNumber : String?
}

data class DetailedProperty(
    override val id : String = "",
    override val image : String = "",
    override val address : String = "",
    override val tenant : String? = null,
    override val available : Boolean = true,
    override val startDate : OffsetDateTime? = null,
    override val endDate : OffsetDateTime? = null,
    override val appartementNumber : String? = "",
    override val area : Int = 0,
    override val rent : Int = 0,
    override val deposit : Int = 0,
    override val documents : Array<Document> = arrayOf(),
    override val zipCode : String = "",
    override val city : String = "",
    override val country : String = "",
    override val name : String = ""
) : IDetailedProperty {
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

    fun toProperty() : Property {
        return Property(
            this.id,
            this.image,
            this.address,
            this.tenant,
            this.available,
            this.startDate,
            this.endDate
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

    suspend fun getPropertiesAsProperties(onError: () -> Unit): Array<Property> {
        try {
            val properties = apiService.getProperties(getBearerToken())
            return properties.map { it.toProperty() }.toTypedArray()
        } catch (e: Exception) {
            onError()
            throw e
        }
    }

    suspend fun addProperty(property: AddPropertyInput, onError: () -> Unit): Property {
        try {
            return apiService.addProperty(getBearerToken(), property).toProperty()
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