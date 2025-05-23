package com.example.keyz.utils

import com.example.keyz.R

object ThemeUtils {
    fun getIcon(isDark: Boolean): Int {
        return if (isDark) {
            R.drawable.keyz_png_logo_white
        } else {
            R.drawable.keyz_png_logo_blue
        }
    }
}