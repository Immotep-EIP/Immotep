package com.example.immotep.realProperty.details

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiClient.AddPropertyInput
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.ApiService
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import com.example.immotep.realProperty.IProperty
import com.example.immotep.realProperty.Property
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import java.time.LocalDateTime
import java.time.OffsetDateTime
import java.util.Date

interface IDetailedProperty : IProperty {
    val area : Int
    val rent : Int
    val deposit : Int
    val documents : Array<String>
    val zipCode : String
    val city : String
    val country : String
    val name : String
}

data class DetailedProperty(
    override val id : String = "",
    override val image : String = "",
    override val address : String = "",
    override val tenant : String? = null,
    override val available : Boolean = true,
    override val startDate : OffsetDateTime? = null,
    override val endDate : OffsetDateTime? = null,
    override val area : Int = 0,
    override val rent : Int = 0,
    override val deposit : Int = 0,
    override val documents : Array<String> = arrayOf(),
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
            postal_code = this.zipCode
        )
    }
}

fun IDetailedProperty.toProperty() : Property {
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


class RealPropertyDetailsViewModel(private val navController: NavController) : ViewModel() {
    private var _property = MutableStateFlow(DetailedProperty())
    val property: StateFlow<DetailedProperty> = _property.asStateFlow()

    fun loadProperty(propertyId: String) {
        viewModelScope.launch {
            var bearerToken = ""
            val authService = AuthService(navController.context.dataStore)
            try {
                bearerToken = authService.getBearerToken()
            } catch(e : Exception) {
                navController.navigate("login")
            }
            try {
                val getPropertyRes = ApiClient.apiService.getProperty(bearerToken, propertyId)
                _property.value = _property.value.copy(
                    id = propertyId,
                    image = "",
                    name = getPropertyRes.name,
                    zipCode = getPropertyRes.postal_code,
                    city = getPropertyRes.city,
                    country = getPropertyRes.country,
                    address = getPropertyRes.address,
                    tenant = getPropertyRes.tenant,
                    available = getPropertyRes.status == "available",
                    startDate = if (getPropertyRes.start_date != null) OffsetDateTime.parse(getPropertyRes.start_date) else null,
                    endDate = if (getPropertyRes.end_date != null) OffsetDateTime.parse(getPropertyRes.end_date) else null,
                    area = getPropertyRes.area_sqm.toInt(),
                    rent = getPropertyRes.rental_price_per_month,
                    deposit = getPropertyRes.deposit_price,
                    documents = arrayOf()
                )
            } catch (e : Exception) {
                e.printStackTrace()
            }
        }
    }

    suspend fun editProperty(property: AddPropertyInput, propertyId: String) {
        val authServ = AuthService(navController.context.dataStore)
        val bearerToken = try {
            authServ.getBearerToken()
        } catch (e : Exception) {
            authServ.onLogout(navController)
            println("error getting token")
            return
        }
        val newProperty = ApiClient.apiService.updateProperty(bearerToken, property, propertyId)
        this._property.value = this._property.value.copy(
            address = newProperty.address,
            tenant = newProperty.tenant,
            available = newProperty.status == "available",
            area = newProperty.area_sqm.toInt(),
            rent = newProperty.rental_price_per_month,
            deposit = newProperty.deposit_price,
        )
    }
}

class RealPropertyDetailsViewModelFactory(private val navController: NavController) :
    ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(RealPropertyDetailsViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return RealPropertyDetailsViewModel(navController) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}
