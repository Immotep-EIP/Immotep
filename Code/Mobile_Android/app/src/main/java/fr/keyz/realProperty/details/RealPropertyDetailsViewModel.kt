package fr.keyz.realProperty.details

import android.content.Context
import android.net.Uri
import android.widget.Toast
import androidx.compose.runtime.mutableStateListOf
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import androidx.navigation.NavController
import fr.keyz.apiCallerServices.AddPropertyInput
import fr.keyz.apiCallerServices.Damage
import fr.keyz.apiCallerServices.DamageCallerService
import fr.keyz.apiCallerServices.DetailedProperty
import fr.keyz.apiCallerServices.Document
import fr.keyz.apiCallerServices.DocumentInput
import fr.keyz.apiCallerServices.InviteDetailedProperty
import fr.keyz.apiCallerServices.InviteTenantCallerService
import fr.keyz.apiCallerServices.PropertyStatus
import fr.keyz.apiCallerServices.RealPropertyCallerService
import fr.keyz.apiClient.ApiService
import fr.keyz.utils.Base64Utils
import fr.keyz.utils.PdfsUtils
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import kotlinx.coroutines.sync.Mutex
import java.time.Instant
import java.time.OffsetDateTime
import java.time.ZoneOffset

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
    private val damageApiCaller = DamageCallerService(apiService, navController)
    private val inviteApiCaller = InviteTenantCallerService(apiService, navController)
    private var _property = MutableStateFlow(DetailedProperty())
    private val _apiError = MutableStateFlow(ApiErrors.NONE)
    private val _isLoading = MutableStateFlow(false)

    val property: StateFlow<DetailedProperty> = _property.asStateFlow()
    val documents = mutableStateListOf<Document>()
    val damages = mutableStateListOf<Damage>()
    val apiError = _apiError.asStateFlow()
    val isLoading = _isLoading.asStateFlow()

    fun setIsLoading(value : Boolean) {
        _isLoading.value = value
    }

    fun loadProperty(newProperty: DetailedProperty) {
        _apiError.value = ApiErrors.NONE
        _property.value = newProperty
        documents.clear()
        damages.clear()
        if (newProperty.lease?.id == null) {
            return
        }
        viewModelScope.launch {
            try {
                setIsLoading(true)
                val propertyDocuments = apiCaller.getPropertyDocuments(newProperty.id, newProperty.lease.id)
                val propertyDamages = damageApiCaller.getPropertyDamages(newProperty.id, newProperty.lease.id)
                documents.addAll(propertyDocuments)
                damages.addAll(propertyDamages)
            } catch (e : Exception) {
                println("Error loading property documents ${e.message}")
                e.printStackTrace()
            } finally {
               setIsLoading(false)
            }
        }
    }

    suspend fun editProperty(propertyEdited: AddPropertyInput, propertyId: String) : String {
        _apiError.value = ApiErrors.NONE
        setIsLoading(true)
        try {
            val (id) = apiCaller.updateProperty(propertyEdited, propertyId)
            _property.value = _property.value.copy(
                name = propertyEdited.name,
                address = propertyEdited.address,
                zipCode = propertyEdited.postal_code,
                city = propertyEdited.city,
                area = propertyEdited.area_sqm.toInt(),
                deposit = propertyEdited.deposit_price,
                rent = propertyEdited.rental_price_per_month,
                country = propertyEdited.country,
                appartementNumber = propertyEdited.apartment_number
            )
            return id
        } catch (e : Exception) {
            e.printStackTrace()
            _apiError.value = ApiErrors.UPDATE_PROPERTY
        } finally {
            setIsLoading(false)
        }
        return propertyId
    }

    fun onSubmitPicture(picture : String) {
        try {
            val pictureDecoded = Base64Utils.decodeBase64ToImage(picture)
                ?: throw Exception("Picture is not a valid base64 string")
            _property.value = _property.value.copy(
                picture = pictureDecoded
            )
        } catch (e : Exception) {
            e.printStackTrace()
        }
    }

    fun openPdf(documentId : String, context: Context) {
        try {
            val document = documents.find { it.id == documentId }?: throw Exception("Document not found")
            val pdfFile = Base64Utils.saveBase64PdfToCache(context, document.data, document.name)
            pdfFile?.let { PdfsUtils.openPdfFile(context, it) }
        } catch (e : Exception) {
            println("Error opening pdf file: ${e.message}")
            Toast.makeText(context, "Error opening pdf file", Toast.LENGTH_SHORT).show()
        }
    }

    fun addDocument(documentUri: Uri, context: Context) {
        viewModelScope.launch {
            try {
                property.value.lease?.id?: throw Exception("Lease id is null")
                val document = Base64Utils.convertPdfUriToBase64(context, documentUri)
                    ?: throw Exception("Error converting uri to base64")
                val documentName = Base64Utils.getFileNameFromUri(context, documentUri)?: "Document"
                val input = DocumentInput(
                    name = documentName,
                    data = document
                )
                val (id) = apiCaller.uploadDocument(
                    propertyId = property.value.id,
                    leaseId = property.value.lease!!.id,
                    document = input
                )
                documents.add(input.toDocument(id))
            } catch (e: Exception) {
                println("Error adding document: ${e.message}")
                Toast.makeText(context, "Error adding document", Toast.LENGTH_SHORT).show()
            }
        }
    }

    fun onSubmitInviteTenant(email: String, startDate: Long, endDate: Long) {
        val newInvite = InviteDetailedProperty(
            startDate = OffsetDateTime.ofInstant(Instant.ofEpochMilli(startDate), ZoneOffset.UTC),
            endDate = OffsetDateTime.ofInstant(Instant.ofEpochMilli(endDate), ZoneOffset.UTC),
            tenantEmail = email
        )
        _property.value = _property.value.copy(
            invite = newInvite,
            status = PropertyStatus.invite_sent
        )
    }

    fun onCancelInviteTenant() {
        viewModelScope.launch {
            setIsLoading(true)
            try {
                inviteApiCaller.cancelInvite(_property.value.id)
                _property.value = _property.value.copy(
                    invite = null,
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

    fun addDamage(damage: Damage) {
        try {
            this.damages.add(damage)
        } catch (e : Exception) {
            e.printStackTrace()
        }
    }
}
