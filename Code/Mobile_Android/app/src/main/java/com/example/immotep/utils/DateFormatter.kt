package com.example.immotep.utils

import java.time.OffsetDateTime
import java.time.format.DateTimeFormatter

object DateFormatter {
    fun formatOffsetDateTime(dateTime : OffsetDateTime?) : String? {
        if (dateTime == null) return null
        try {
            val formattedDate = dateTime.format(DateTimeFormatter.ISO_LOCAL_DATE)
            return formattedDate
        } catch (e : Exception) {
            return "Invalid date"
        }
    }
}