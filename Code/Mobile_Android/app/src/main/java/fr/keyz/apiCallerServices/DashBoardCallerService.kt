package fr.keyz.apiCallerServices

import androidx.navigation.NavController
import fr.keyz.apiClient.ApiService
import java.util.Locale

//enum classes

enum class Priority {
    low, medium, high, urgent,
}

//data classes

data class DashBoardOpenDamage(
    val listToFix: Array<Damage> = arrayOf(),
    val nbrHigh: Int = 0,
    val nbrLow: Int = 0,
    val nbrMedium: Int = 0,
    val nbrPlannedToFixThisWeek: Int = 0,
    val nbrTotal: Int = 0,
    val nbrUrgent: Int = 0
)

data class DashBoardOpenDamageOutput(
    val list_to_fix: Array<DamageOutput>?,
    val nbr_high: Int,
    val nbr_low: Int,
    val nbr_medium: Int,
    val nbr_planned_to_fix_this_week: Int,
    val nbr_total: Int,
    val nbr_urgent: Int
) {
    fun toDashBoardOpenDamage() : DashBoardOpenDamage {
        return DashBoardOpenDamage(
            listToFix = this.list_to_fix?.map { it.toDamage() }?.toTypedArray() ?: arrayOf(),
            nbrHigh = this.nbr_high,
            nbrLow = this.nbr_low,
            nbrMedium = this.nbr_medium,
            nbrPlannedToFixThisWeek = this.nbr_planned_to_fix_this_week,
            nbrTotal = this.nbr_total,
            nbrUrgent = this.nbr_urgent
        )
    }
}

data class DashBoardPropertyOutput(
    val id: String,
    val apartment_number: String?,
    val archived: Boolean,
    val owner_id: String,
    val name: String,
    val address: String,
    val city: String,
    val postal_code: String,
    val country: String,
    val area_sqm: Double,
    val rental_price_per_month: Int,
    val deposit_price: Int,
    val created_at: String,
) {
    fun toDetailedProperty() : DetailedProperty {
        return DetailedProperty(
            id = this.id,
            name = this.name,
            zipCode = this.postal_code,
            city = this.city,
            country = this.country,
            address = this.address,
            appartementNumber = this.apartment_number,
            area = this.area_sqm.toInt(),
            rent = this.rental_price_per_month,
            deposit = this.deposit_price,
        )
    }
}

data class DashBoardProperties(
    val listRecentlyAdded: Array<DetailedProperty> = arrayOf(),
    val nbrArchived: Int = 0,
    val nbrAvailable: Int = 0,
    val nbrOccupied: Int = 0,
    val nbrPendingInvites: Int = 0,
    val nbrTotal: Int = 0
)

data class DashBoardPropertiesOutput(
    val list_recently_added: Array<DashBoardPropertyOutput>?,
    val nbr_archived: Int,
    val nbr_available: Int,
    val nbr_occupied: Int,
    val nbr_pending_invites: Int,
    val nbr_total: Int
)  {
    fun toDashBoardProperties() : DashBoardProperties {
        return DashBoardProperties(
            listRecentlyAdded = if (this.list_recently_added != null) {
                this.list_recently_added.map { it.toDetailedProperty() }.toTypedArray()
            } else arrayOf(),
            nbrArchived = this.nbr_archived,
            nbrAvailable = this.nbr_available,
            nbrOccupied = this.nbr_occupied,
            nbrPendingInvites = this.nbr_pending_invites,
            nbrTotal = this.nbr_total
        )
    }
}


data class DashBoardReminder(
    val advice: String,
    val id: String,
    val link: String,
    val priority: Priority,
    val title: String
)

data class GetDashBoardOutput(
    val open_damages : DashBoardOpenDamageOutput,
    val properties: DashBoardPropertiesOutput,
    val reminders: Array<DashBoardReminder>
) {
    fun toDashBoard() : DashBoard {
        return DashBoard(
            openDamages = this.open_damages.toDashBoardOpenDamage(),
            reminders = this.reminders,
            properties = this.properties.toDashBoardProperties()
        )
    }
}

data class DashBoard(
    val reminders: Array<DashBoardReminder> = arrayOf(),
    val openDamages: DashBoardOpenDamage = DashBoardOpenDamage(),
    val properties: DashBoardProperties = DashBoardProperties()
)


class DashBoardCallerService (
    apiService: ApiService,
    navController: NavController,
) : ApiCallerService(apiService, navController) {

    suspend fun getDashBoard(
    ) : DashBoard {
        return changeRetrofitExceptionByApiCallerException {
            apiService.getDashboard(this.getBearerToken(), Locale.getDefault().language).toDashBoard()
        }
    }
}