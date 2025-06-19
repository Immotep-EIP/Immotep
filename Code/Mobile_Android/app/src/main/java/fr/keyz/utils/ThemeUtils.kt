package fr.keyz.utils

import fr.keyz.R

object ThemeUtils {
    fun getIcon(isDark: Boolean): Int {
        return if (isDark) {
            R.drawable.keyz_png_logo_white
        } else {
            R.drawable.keyz_png_logo_blue
        }
    }
}