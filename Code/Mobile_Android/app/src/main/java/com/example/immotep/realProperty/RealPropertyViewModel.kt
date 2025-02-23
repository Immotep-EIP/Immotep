package com.example.immotep.realProperty

import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiClient.AddPropertyInput
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.ApiService
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import com.example.immotep.realProperty.details.toProperty
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import java.time.LocalDateTime
import java.time.OffsetDateTime
import java.time.ZonedDateTime
import java.util.Date
import kotlin.collections.map
import kotlin.collections.toTypedArray

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

class RealPropertyViewModel(private val navController: NavController) : ViewModel() {
    val properties = mutableStateListOf<Property>()

    fun getProperties() {
        viewModelScope.launch {
            val bearerToken: String
            val authService = AuthService(navController.context.dataStore)
            try {
                bearerToken = authService.getBearerToken()
            } catch (e : Exception) {
                authService.onLogout(navController)
                println("error getting token")
                return@launch
            }
            properties.clear()
            try {
                val newProperties = ApiClient.apiService.getProperties(bearerToken)
                newProperties.forEach {
                    try {
                        properties.add(
                            Property(
                                id = it.id,
                                image = "",
                                address = it.address,
                                tenant = it.tenant,
                                available = it.status == "available",
                                startDate = if (it.start_date != null) OffsetDateTime.parse(it.start_date) else null,
                                endDate = if (it.end_date != null) OffsetDateTime.parse(it.end_date) else null
                            )
                        )
                    } catch (e : Exception) {
                        println("error adding this property to the list ${e.message}")
                    }
                }
            } catch (e : Exception) {
                println("error getting properties ${e.message}")
                e.printStackTrace()
            }
        }
    }

    suspend fun addProperty(propertyForm: AddPropertyInput) {
        val authServ = AuthService(navController.context.dataStore)
        val bearerToken = try {
            authServ.getBearerToken()
        } catch (e : Exception) {
            authServ.onLogout(navController)
            println("error getting token")
            return
        }
        val newProperty = ApiClient.apiService.addProperty(bearerToken, propertyForm)
        properties.add(newProperty.toProperty())
    }

    fun deleteProperty(propertyId: String) {
        val index = properties.indexOfFirst { it.id == propertyId }
        if (index == -1) {
            return
        }
        viewModelScope.launch {
            val bearerToken: String
            val authService = AuthService(navController.context.dataStore)
            try {
                bearerToken = authService.getBearerToken()
            } catch (e : Exception) {
                authService.onLogout(navController)
                println("error getting token")
                return@launch
            }
            try {
                ApiClient.apiService.archiveProperty(bearerToken, propertyId)
                properties.removeAt(index)
            } catch (e : Exception) {
                println("error deleting property ${e.message}")
                e.printStackTrace()
            }
        }
    }
}

class RealPropertyViewModelFactory(private val navController: NavController) :
    ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(RealPropertyViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return RealPropertyViewModel(navController) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}
