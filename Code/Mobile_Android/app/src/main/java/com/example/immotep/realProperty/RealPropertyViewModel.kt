package com.example.immotep.realProperty

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.navigation.NavController
import com.example.immotep.apiClient.newTestArray
import com.example.immotep.realProperty.details.toProperty
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import java.util.Date
import kotlin.collections.map
import kotlin.collections.toTypedArray

interface IProperty {
    val id: String
    val image: String
    val address: String
    val tenant: String
    val available: Boolean
    val startDate: Date
    val endDate: Date?
}

data class Property(
    override val id: String = "",
    override val image: String = "",
    override val address: String = "",
    override val tenant: String = "",
    override val available: Boolean = true,
    override val startDate: Date = Date(),
    override val endDate: Date? = null
) : IProperty

class RealPropertyViewModel(private val navController: NavController) : ViewModel() {
    private val _properties = MutableStateFlow(Array(1) { Property() })
    val properties: StateFlow<Array<Property>> = _properties.asStateFlow()
    init {
        val properties: Array<Property> = newTestArray.map { it.toProperty() }.toTypedArray()
        this._properties.value = properties
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
