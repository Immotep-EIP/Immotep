package com.example.immotep.addPropertyModal

import android.net.Uri
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow

data class PropertyForm(
    val address: String = "",
    val zipCode: String = "",
    val country: String = "",
    val area: Int = 0,
    val rental: Int = 0,
    val deposit: Int = 0,
)

data class PropertyFormError(
    var address: Boolean = false,
    var zipCode: Boolean = false,
    var country: Boolean = false,
    var area: Boolean = false,
    var rental: Boolean = false,
    var deposit: Boolean = false,
)

class AddPropertyViewModelViewModel() : ViewModel() {
    private val _propertyForm = MutableStateFlow(PropertyForm())
    private val _propertyFormError = MutableStateFlow(PropertyFormError())
    val pictures = mutableStateListOf<Uri>()
    val propertyForm: StateFlow<PropertyForm> = _propertyForm.asStateFlow()
    val propertyFormError: StateFlow<PropertyFormError> = _propertyFormError.asStateFlow()

    fun setAddress(address: String) {
        _propertyForm.value = _propertyForm.value.copy(address = address)
    }
    fun setZipCode(zipCode: String) {
        _propertyForm.value = _propertyForm.value.copy(zipCode = zipCode)
    }

    fun setCountry(country: String) {
        _propertyForm.value = _propertyForm.value.copy(country = country)
    }

    fun setArea(area: Int) {
        _propertyForm.value = _propertyForm.value.copy(area = area)
    }

    fun setRental(rental: Int) {
        _propertyForm.value = _propertyForm.value.copy(rental = rental)
    }

    fun setDeposit(deposit: Int) {
        _propertyForm.value = _propertyForm.value.copy(deposit = deposit)
    }

    fun addPicture(picture: Uri) {
        pictures.add(picture)
    }

    fun reset() {
        _propertyForm.value = PropertyForm()
    }

    fun onSubmit(onClose : () -> Unit) {
        val newPropertyErrors : PropertyFormError = PropertyFormError()
        if (_propertyForm.value.address.length < 3) {
            newPropertyErrors.address = true
        }
        if (_propertyForm.value.zipCode.length < 3) {
            newPropertyErrors.zipCode = true
        }
        if (_propertyForm.value.country.length < 3) {
            newPropertyErrors.country = true
        }
        if (_propertyForm.value.area < 1) {
            newPropertyErrors.area = true
        }
        if (_propertyForm.value.rental < 1) {
            newPropertyErrors.rental = true
        }
        if (_propertyForm.value.deposit < 1) {
            newPropertyErrors.deposit = true
            }
        if (newPropertyErrors.address || newPropertyErrors.zipCode || newPropertyErrors.country || newPropertyErrors.area || newPropertyErrors.rental || newPropertyErrors.deposit) {
            _propertyFormError.value = newPropertyErrors
            return
        }
    }
}
