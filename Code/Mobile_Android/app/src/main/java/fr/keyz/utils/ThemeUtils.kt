package fr.keyz.utils

import androidx.compose.ui.graphics.Color
import fr.keyz.R
import fr.keyz.apiCallerServices.Priority

object ThemeUtils {
    fun getIcon(isDark: Boolean): Int {
        return if (isDark) {
            R.drawable.keyz_png_logo_white
        } else {
            R.drawable.keyz_png_logo_blue
        }
    }

    fun getStatusColor(priority: Priority): Color {
        return when (priority) {
            Priority.low -> Color(0xFF90CAF9)
            Priority.medium -> Color(0xFFFFD176)
            Priority.high -> Color(0xFFFF9862)
            Priority.urgent -> Color(0xFFF44336)
        }
    }
}