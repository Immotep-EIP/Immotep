package fr.keyz.apiCallerServices

import androidx.navigation.NavController
import fr.keyz.apiClient.ApiService
import fr.keyz.apiClient.CreateOrUpdateResponse
import java.time.OffsetDateTime

data class DamageInput(
    val comment: String = "",
    val pictures: ArrayList<String> = ArrayList(),
    val priority: Priority = Priority.low,
    val room_id: String? = null
) {
    fun toDamage(id : String, roomName: String, tenantName: String) : Damage {
        return Damage(
            id = id,
            comment = comment,
            createdAt = OffsetDateTime.now(),
            fixPlannedAt = null,
            fixStatus = DamageStatus.PENDING,
            fixedAt = null,
            leaseId = "",
            pictures = pictures.toTypedArray(),
            priority = priority,
            read = false,
            roomId = room_id.let { "" },
            roomName = roomName,
            tenantName = tenantName,
            updatedAt = OffsetDateTime.now()
        )
    }
}

data class UpdateDamageInput(
    val fix_planned_at: String,
    val read : Boolean
)

data class DamageOutput(
    val comment: String,
    val created_at: String,
    val fix_planned_at: String?,
    val fix_status: String,
    val fixed_at: String?,
    val id: String,
    val lease_id: String,
    val pictures: Array<String>?,
    val priority: String,
    val read: Boolean,
    val room_id: String,
    val room_name: String,
    val tenant_name: String,
    val updated_at: String
) {
    fun toDamage() : Damage {
        return Damage(
            id = this.id,
            comment = this.comment,
            createdAt = OffsetDateTime.parse(this.created_at),
            fixPlannedAt = this.fix_planned_at?.let { OffsetDateTime.parse(it) },
            fixStatus = stringToDamageStatus(this.fix_status),
            fixedAt = this.fixed_at?.let { OffsetDateTime.parse(it) },
            leaseId = this.lease_id,
            pictures = this.pictures ?: arrayOf(),
            priority = stringToPriority(this.priority),
            read = this.read,
            roomId = this.room_id,
            roomName = this.room_name,
            tenantName = this.tenant_name,
            updatedAt = OffsetDateTime.parse(this.updated_at)
        )
    }
}

data class Damage(
    val id: String,
    val comment: String,
    val createdAt: OffsetDateTime,
    val fixPlannedAt: OffsetDateTime?,
    val fixStatus: DamageStatus,
    val fixedAt: OffsetDateTime?,
    val leaseId: String,
    val pictures: Array<String>,
    val priority: Priority,
    val read: Boolean,
    val roomId: String,
    val roomName: String,
    val tenantName: String,
    val updatedAt: OffsetDateTime
)

class DamageCallerService (
    apiService: ApiService,
    navController: NavController,
) : ApiCallerService(apiService, navController) {
    suspend fun getPropertyDamages(
        propertyId: String,
        leaseId: String
    ) : Array<Damage> = changeRetrofitExceptionByApiCallerException {
        if (this.isOwner()) {
            apiService.getPropertyDamages(getBearerToken(), propertyId, leaseId).map { it.toDamage() }.toTypedArray()
        } else {
            apiService.getPropertyDamagesTenant(getBearerToken(), "current").map { it.toDamage() }.toTypedArray()
        }
    }

    suspend fun getDamage(
        propertyId: String,
        leaseId: String,
        damageId : String
    ) : Damage = changeRetrofitExceptionByApiCallerException {
        if (this.isOwner()) {
            apiService.getPropertyDamage(getBearerToken(), propertyId, leaseId, damageId).toDamage()
        } else {
            apiService.getPropertyDamageTenant(getBearerToken(), "current", damageId).toDamage()
        }
    }

    suspend fun addDamage(damageInput: DamageInput) : CreateOrUpdateResponse = changeRetrofitExceptionByApiCallerException {
        apiService.addDamage(getBearerToken(), "current", damage = damageInput)
    }

    suspend fun fixDamage(propertyId: String?,
                          leaseId: String,
                          damageId : String) : CreateOrUpdateResponse =
        changeRetrofitExceptionByApiCallerException {
            if (!isOwner()) {
                apiService.fixDamageTenant(getBearerToken(), "current", damageId)
            } else if (!propertyId.isNullOrEmpty()) {
                apiService.fixDamageOwner(getBearerToken(), propertyId, leaseId, damageId)
            } else {
                throw IllegalArgumentException("Missing propertyId")
            }
    }

    suspend fun updateDamageOwner(
        propertyId: String,
        leaseId: String,
        damageId : String,
        updateDamageInput: UpdateDamageInput
    ) : CreateOrUpdateResponse = changeRetrofitExceptionByApiCallerException {
        apiService.updateDamageOwner(
            getBearerToken(),
            propertyId,
            leaseId,
            damageId,
            updateDamageInput
        )
    }
}