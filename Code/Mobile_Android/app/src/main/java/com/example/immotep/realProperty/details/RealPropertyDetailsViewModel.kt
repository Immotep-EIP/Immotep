package com.example.immotep.realProperty.details

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import com.example.immotep.apiClient.newTestArray
import com.example.immotep.realProperty.IProperty
import com.example.immotep.realProperty.Property
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
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
    override val startDate : Date = Date(),
    override val endDate : Date? = null,
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


class RealPropertyDetailsViewModel(private val propertyId: String) : ViewModel() {
    private var _property = MutableStateFlow(DetailedProperty())
    val property: StateFlow<DetailedProperty> = _property.asStateFlow()
    init {
        val newValue = newTestArray.find { it.id == propertyId }
        if (newValue != null) {
            this._property.value = newValue
        }
    }
}

class RealPropertyDetailsViewModelFactory(private val propertyId: String) :
    ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(RealPropertyDetailsViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return RealPropertyDetailsViewModel(propertyId) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}
