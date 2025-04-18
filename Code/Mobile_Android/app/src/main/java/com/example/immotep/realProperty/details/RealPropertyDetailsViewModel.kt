package com.example.immotep.realProperty.details

import android.content.Context
import android.widget.Toast
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.AddPropertyInput
import com.example.immotep.apiCallerServices.DetailedProperty
import com.example.immotep.apiCallerServices.InviteTenantCallerService
import com.example.immotep.apiCallerServices.PropertyStatus
import com.example.immotep.apiCallerServices.RealPropertyCallerService
import com.example.immotep.apiClient.ApiClient
import com.example.immotep.apiClient.ApiService
import com.example.immotep.authService.AuthService
import com.example.immotep.login.dataStore
import com.example.immotep.utils.Base64Utils
import com.example.immotep.utils.PdfsUtils
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import kotlinx.coroutines.sync.Mutex
import java.time.Instant
import java.time.LocalDateTime
import java.time.OffsetDateTime
import java.time.ZoneOffset
import java.util.Date

class RealPropertyDetailsViewModel(
    navController: NavController,
    apiService: ApiService
) : ViewModel() {
    enum class ApiErrors {
        GET_PROPERTY,
        UPDATE_PROPERTY,
        NONE
    }
    private val apiCaller = RealPropertyCallerService(apiService, navController)
    private val inviteApiCaller = InviteTenantCallerService(apiService, navController)
    private var _property = MutableStateFlow(DetailedProperty())
    private val _apiError = MutableStateFlow(ApiErrors.NONE)
    private val _isLoading = MutableStateFlow(false)
    private val _isLoadingMutex = Mutex()

    val property: StateFlow<DetailedProperty> = _property.asStateFlow()
    val apiError = _apiError.asStateFlow()
    val isLoading = _isLoading.asStateFlow()

    fun setIsLoading(value : Boolean) {
        viewModelScope.launch {
            _isLoadingMutex.lock()
            _isLoading.value = value
            _isLoadingMutex.unlock()
        }
    }

    fun loadProperty(newProperty: DetailedProperty) {
        setIsLoading(true)
        _apiError.value = ApiErrors.NONE
        _property.value = newProperty
        viewModelScope.launch {
            try {
                val propertyDocuments = apiCaller.getPropertyDocuments(newProperty.id) { }
                _property.value = newProperty.copy(documents = propertyDocuments)
            } catch (e : Exception) {
                println("Error loading property ${e.message}")
                e.printStackTrace()
            } finally {
               setIsLoading(false)
            }
        }
    }

    suspend fun editProperty(property: AddPropertyInput, propertyId: String) {
        _apiError.value = ApiErrors.NONE
        setIsLoading(true)
        try {
            _property.value = apiCaller.updateProperty(property, propertyId) { _apiError.value = ApiErrors.UPDATE_PROPERTY }
        } catch (e : Exception) {
            e.printStackTrace()
            _apiError.value = ApiErrors.UPDATE_PROPERTY
        } finally {
            setIsLoading(false)
        }
    }

    fun openPdf(documentId : String, context: Context) {
        try {
            val document = _property.value.documents.find { it.id == documentId }
            if (document == null) throw Exception("Document not found")
            val pdfFile = Base64Utils.saveBase64PdfToCache(context, document.data, document.name)
            pdfFile?.let { PdfsUtils.openPdfFile(context, it) }
        } catch (e : Exception) {
            println("Error opening pdf file: ${e.message}")
            Toast.makeText(context, "Error opening pdf file", Toast.LENGTH_SHORT).show()
        }
    }

    fun onSubmitInviteTenant(email: String, startDate: Long, endDate: Long) {
        _property.value = _property.value.copy(
            tenant = email,
            startDate = OffsetDateTime.ofInstant(Instant.ofEpochMilli(startDate), ZoneOffset.UTC),
            endDate = OffsetDateTime.ofInstant(Instant.ofEpochMilli(endDate), ZoneOffset.UTC),
            status = PropertyStatus.invite_sent
        )
    }

    fun onCancelInviteTenant() {
        viewModelScope.launch {
            setIsLoading(true)
            try {
                inviteApiCaller.cancelInvite(_property.value.id)
                _property.value = _property.value.copy(
                    tenant = null,
                    startDate = null,
                    endDate = null,
                    status = PropertyStatus.available
                )
            } catch (e : Exception) {
                println(e.message)
                _apiError.value = ApiErrors.UPDATE_PROPERTY
            } finally {
                setIsLoading(false)
            }
        }
    }
}
