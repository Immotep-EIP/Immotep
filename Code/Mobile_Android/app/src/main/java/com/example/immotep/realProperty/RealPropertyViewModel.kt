package com.example.immotep.realProperty

import androidx.lifecycle.ViewModel
import androidx.lifecycle.ViewModelProvider
import androidx.navigation.NavController
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import java.util.Date
import java.util.Vector

data class Property(
    val image : String = "",
    val address : String = "",
    val tenant : String = "",
    val available : Boolean = true,
    val startDate : Date = Date(),
    val endDate : Date? = null
)


class RealPropertyViewModel(private val navController: NavController) : ViewModel() {
    private val _properties = MutableStateFlow(Array<Property>(1) {Property()})
    val properties: StateFlow<Array<Property>> = _properties.asStateFlow()
    init {
        val newTestArray = arrayOf(Property(
            "https://plus.unsplash.com/premium_photo-1661915661139-5b6a4e4a6fcc?fm=jpg&q=60&w=3000&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MXx8aG91c2V8ZW58MHx8MHx8fDA%3D",
            "19 rue de la paix, Paris 75000",
            "John Doe",
            true,
            Date(),
            null
        ),
        Property(
            "https://plus.unsplash.com/premium_photo-1661915661139-5b6a4e4a6fcc?fm=jpg&q=60&w=3000&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MXx8aG91c2V8ZW58MHx8MHx8fDA%3D",
            "30 rue de la source, Lyon 69000",
            "Tom Nook",
            false,
            Date(),
            null
        ),
        Property(
            "https://plus.unsplash.com/premium_photo-1661915661139-5b6a4e4a6fcc?fm=jpg&q=60&w=3000&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MXx8aG91c2V8ZW58MHx8MHx8fDA%3D",
            "1 rue de la companie des indes, Marseille 13000",
            "Crash Bandicoot",
            false,
            Date(),
            Date()
        ))
        this._properties.value = newTestArray

    }

}

class RealPropertyViewModelFactory(private val navController: NavController) :
    ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        if (modelClass.isAssignableFrom(RealPropertyViewModel::class.java)) {
            @Suppress("UNCHECKED_CAST")
            return RealPropertyViewModel(navController) as T
        }
        throw IllegalArgumentException("Unknown ViewModel class")
    }
}