package fr.keyz.utils

object RegexUtils {
    fun isValidEmail(email: String): Boolean {
        return Regex("^[A-Za-z0-9+_.-]+@[A-Za-z0-9.-]+$").matches(email)
    }
}