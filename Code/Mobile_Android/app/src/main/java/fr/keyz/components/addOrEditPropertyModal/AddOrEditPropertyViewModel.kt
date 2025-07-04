package fr.keyz.components.addOrEditPropertyModal

import android.content.Context
import android.net.Uri
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import fr.keyz.apiCallerServices.AddPropertyInput
import fr.keyz.apiCallerServices.ApiCallerServiceException
import fr.keyz.apiCallerServices.RealPropertyCallerService
import fr.keyz.apiClient.ApiService
import fr.keyz.utils.Base64Utils
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

class AddOrEditPropertyViewModel(apiService: ApiService, navController: NavController) : ViewModel() {
    private val _callerService = RealPropertyCallerService(apiService, navController)
    private val _propertyForm = MutableStateFlow(AddPropertyInput())
    private val _propertyFormError = MutableStateFlow(PropertyFormError())
    private val _picture = MutableStateFlow<Uri?>(null)
    private val _isLoading = MutableStateFlow(false)
    val picture = _picture.asStateFlow()
    val propertyForm: StateFlow<AddPropertyInput> = _propertyForm.asStateFlow()
    val propertyFormError: StateFlow<PropertyFormError> = _propertyFormError.asStateFlow()
    val isLoading: StateFlow<Boolean> = _isLoading.asStateFlow()

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

    fun setPicture(picture: Uri) {
        _picture.value = picture
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

    fun onSubmit(
        onClose : () -> Unit,
        sendFormFn : suspend (property : AddPropertyInput) -> String,
        updateUserPicture : (propertyId : String, picture : String) -> Unit,
        context : Context
    ) {
        _isLoading.value = true
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
            _isLoading.value = false
            println("ERROR $newPropertyErrors")
            return
        }
        viewModelScope.launch {
            try {
                val propertyId = sendFormFn(propertyForm.value)
                if (picture.value != null) {
                    val pictureBase64 = Base64Utils.encodeImageToBase64(
                        fileUri = picture.value!!,
                        context = context
                    )
                    _callerService.updatePropertyPicture(propertyId, pictureBase64)
                    updateUserPicture(propertyId, pictureBase64)
                }
                onClose()
                reset()
            } catch (e: ApiCallerServiceException) {
                println("Error on the api during property creation: ${e.message}")
                e.printStackTrace()
            } catch (e : Exception) {
                println("Error during property creation: ${e.message}")
                e.printStackTrace()
            }
            finally {
                _isLoading.value = false
            }
        }
    }
}
