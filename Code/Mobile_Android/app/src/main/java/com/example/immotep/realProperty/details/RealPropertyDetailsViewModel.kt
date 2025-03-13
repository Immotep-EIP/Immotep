package com.example.immotep.realProperty.details

import android.content.Context
import android.widget.Toast
import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import com.example.immotep.apiCallerServices.AddPropertyInput
import com.example.immotep.apiCallerServices.DetailedProperty
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
import java.time.LocalDateTime
import java.time.OffsetDateTime
import java.util.Date

class RealPropertyDetailsViewModel(
    private val navController: NavController,
    apiService: ApiService
) : ViewModel() {
    enum class ApiErrors {
        GET_PROPERTY,
        UPDATE_PROPERTY,
        NONE
    }
    private val apiCaller = RealPropertyCallerService(apiService, navController)
    private var _property = MutableStateFlow(DetailedProperty())
    private val _apiError = MutableStateFlow(ApiErrors.NONE)

    val property: StateFlow<DetailedProperty> = _property.asStateFlow()
    val apiError = _apiError.asStateFlow()

    fun loadProperty(propertyId: String) {
        _apiError.value = ApiErrors.NONE
        viewModelScope.launch {
            try {
                val newProperty = apiCaller.getPropertyWithDetails(propertyId) { _apiError.value = ApiErrors.GET_PROPERTY }
                val propertyDocuments = apiCaller.getPropertyDocuments(propertyId) { _apiError.value = ApiErrors.GET_PROPERTY }
                _property.value = newProperty
                _property.value = newProperty.copy(documents = propertyDocuments)

            } catch (e : Exception) {
                println("Error loading property ${e.message}")
                e.printStackTrace()
            }
        }
    }

    suspend fun editProperty(property: AddPropertyInput, propertyId: String) {
        _apiError.value = ApiErrors.NONE
        _property.value = apiCaller.updateProperty(property, propertyId) { _apiError.value = ApiErrors.UPDATE_PROPERTY }
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
}
