package com.example.keyz.apiCallerServices

import androidx.navigation.NavController
import com.example.keyz.apiClient.ApiService


data class DashBoardOpenDamage(
    val listToFix: Array<Damage>,
    val nbrHigh: Int,
    val nbrLow: Int,
    val nbrMedium: Int,
    val nbrPlannedToFixThisWeek: Int,
    val nbrTotal: Int,
    val nbrUrgent: Int
)

data class DashBoardOpenDamageOutput(
    val list_to_fix: Array<DamageOutput>,
    val nbr_high: Int,
    val nbr_low: Int,
    val nbr_medium: Int,
    val nbr_planned_to_fix_this_week: Int,
    val nbr_total: Int,
    val nbr_urgent: Int
) {
    fun toDashBoardOpenDamage() : DashBoardOpenDamage {
        return DashBoardOpenDamage(
            listToFix = this.list_to_fix.map { it.toDamage() }.toTypedArray(),
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
    val listRecentlyAdded: Array<DetailedProperty>,
    val nbrArchived: Int,
    val nbrAvailable: Int,
    val nbrOccupied: Int,
    val nbrPendingInvites: Int,
    val nbrTotal: Int
)

data class DashBoardPropertiesOutput(
    val list_recently_added: Array<DashBoardPropertyOutput>,
    val nbr_archived: Int,
    val nbr_available: Int,
    val nbr_occupied: Int,
    val nbr_pending_invites: Int,
    val nbr_total: Int
)  {
    fun toDashBoardProperties() : DashBoardProperties {
        return DashBoardProperties(
            listRecentlyAdded = this.list_recently_added.map { it.toDetailedProperty() }
                .toTypedArray(),
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
    val priority: String,
    val title: String
)

data class GetDashBoardOutput(
    val open_damages : DashBoardOpenDamageOutput,
    val properties: DashBoardPropertiesOutput,
    val reminders: Array<DashBoardReminder>
) {
    fun toGetDashBoard() : GetDashBoard {
        return GetDashBoard(
            openDamages = this.open_damages.toDashBoardOpenDamage(),
            properties = this.properties.toDashBoardProperties(),
            reminders = this.reminders
        )
    }
}

data class GetDashBoard(
    val openDamages : DashBoardOpenDamage,
    val properties: DashBoardProperties,
    val reminders: Array<DashBoardReminder>
)


class DashBoardCallerService (
    apiService: ApiService,
    navController: NavController,
) : ApiCallerService(apiService, navController) {

    suspend fun getDashBoard(
    ) : Array<GetDashBoard> {
        return changeRetrofitExceptionByApiCallerException {
            apiService.getDashboard(this.getBearerToken(), "eng").map { it.toGetDashBoard() }.toTypedArray()
        }
    }
}