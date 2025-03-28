package com.example.immotep.addOrEditPropertyModal

import android.net.Uri
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.AddPropertyInput
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

data class PropertyFormError(
    var address: Boolean = false,
    var zipCode: Boolean = false,
    var country: Boolean = false,
    var area: Boolean = false,
    var rental: Boolean = false,
    var deposit: Boolean = false,
    var name: Boolean = false,
    var city: Boolean = false
)

class AddOrEditPropertyViewModel : ViewModel() {
    private val _propertyForm = MutableStateFlow(AddPropertyInput())
    private val _propertyFormError = MutableStateFlow(PropertyFormError())
    val pictures = mutableStateListOf<Uri>()
    val propertyForm: StateFlow<AddPropertyInput> = _propertyForm.asStateFlow()
    val propertyFormError: StateFlow<PropertyFormError> = _propertyFormError.asStateFlow()

    fun setBaseValue(property: AddPropertyInput) {
        _propertyForm.value = property
    }

    fun setAddress(address: String) {
        _propertyForm.value = _propertyForm.value.copy(address = address)
    }
    fun setZipCode(zipCode: String) {
        _propertyForm.value = _propertyForm.value.copy(postal_code = zipCode)
    }

    fun setCountry(country: String) {
        _propertyForm.value = _propertyForm.value.copy(country = country)
    }

    fun setArea(area: Double) {
        _propertyForm.value = _propertyForm.value.copy(area_sqm = area)
    }

    fun setRental(rental: Int) {
        _propertyForm.value = _propertyForm.value.copy(rental_price_per_month = rental)
    }

    fun setDeposit(deposit: Int) {
        _propertyForm.value = _propertyForm.value.copy(deposit_price = deposit)
    }

    fun setName(name: String) {
        _propertyForm.value = _propertyForm.value.copy(name = name)
    }

    fun setCity(city: String) {
        _propertyForm.value = _propertyForm.value.copy(city = city)
    }

    fun addPicture(picture: Uri) {
        pictures.add(picture)
    }

    fun setAppartementNumber(appartementNumber: String) {
        _propertyForm.value = _propertyForm.value.copy(apartment_number = appartementNumber)
    }

    fun reset(baseValue: AddPropertyInput? = null) {
        if (baseValue != null) {
            _propertyForm.value = baseValue
            return
        }
        _propertyForm.value = AddPropertyInput()
    }

    fun onSubmit(onClose : () -> Unit, sendFormFn : suspend (property : AddPropertyInput) -> Unit) {
        val newPropertyErrors = PropertyFormError()
        if (_propertyForm.value.address.length < 3) {
            newPropertyErrors.address = true
        }
        if (_propertyForm.value.postal_code.length != 5) {
            newPropertyErrors.zipCode = true
        }
        if (_propertyForm.value.country.length < 3) {
            newPropertyErrors.country = true
        }
        if (_propertyForm.value.area_sqm < 1) {
            newPropertyErrors.area = true
        }
        if (_propertyForm.value.rental_price_per_month < 1) {
            newPropertyErrors.rental = true
        }
        if (_propertyForm.value.deposit_price < 1) {
            newPropertyErrors.deposit = true
            }
        if (newPropertyErrors.address || newPropertyErrors.zipCode || newPropertyErrors.country || newPropertyErrors.area || newPropertyErrors.rental || newPropertyErrors.deposit) {
            _propertyFormError.value = newPropertyErrors
            println("ERROR $newPropertyErrors")
            return
        }
            viewModelScope.launch {
                try {
                    sendFormFn(propertyForm.value)
                    onClose()
                    reset()
                } catch (e: Exception) {
                    println("Error during property creation: ${e.message}")
                    e.printStackTrace()
                }
            }
    }
}
