package com.example.immotep.realProperty.details

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
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
import java.util.Date

interface IDetailedProperty : IProperty {
    val area : Int
    val rent : Int
    val deposit : Int
    val documents : Array<String>
}

data class DetailedProperty(
    override val id : String = "",
    override val image : String = "",
    override val address : String = "",
    override val tenant : String = "",
    override val available : Boolean = true,
    override val startDate : LocalDateTime? = null,
    override val endDate : LocalDateTime? = null,
    override val area : Int = 0,
    override val rent : Int = 0,
    override val deposit : Int = 0,
    override val documents : Array<String> = arrayOf()
) : IDetailedProperty

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


class RealPropertyDetailsViewModel(private val propertyId: String, private val navController: NavController) : ViewModel() {
    private var _property = MutableStateFlow(DetailedProperty())
    val property: StateFlow<DetailedProperty> = _property.asStateFlow()

    fun loadProperty() {
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
                    address = getPropertyRes.address,
                    tenant = getPropertyRes.tenant,
                    available = getPropertyRes.status == "available",
                    startDate = if (getPropertyRes.start_date != null) LocalDateTime.parse(getPropertyRes.start_date) else null,
                    endDate = if (getPropertyRes.end_date != null) LocalDateTime.parse(getPropertyRes.end_date) else null,
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
}

class RealPropertyDetailsViewModelFactory(private val propertyId: String, private val navController: NavController) :
    ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(RealPropertyDetailsViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return RealPropertyDetailsViewModel(propertyId, navController) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}
