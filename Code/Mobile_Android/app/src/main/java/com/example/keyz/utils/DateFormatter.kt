package com.example.keyz.utils

import java.text.SimpleDateFormat
import java.time.OffsetDateTime
import java.time.format.DateTimeFormatter
import java.util.Locale

object DateFormatter {

    fun formatOffsetDateTime(dateTime : OffsetDateTime?) : String? {
        if (dateTime == null) return null
        try {
            val formattedDate = dateTime.format(DateTimeFormatter.ISO_LOCAL_DATE)
            return formattedDate.replace("-", "/")
        } catch (e : Exception) {
            return "Invalid date"
        }
    }

    fun currentDateAsOffsetDateTimeString() : String {
        try {
            val formatter = SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSSXXX", Locale.getDefault())
            return formatter.format(java.util.Date().time)
        } catch (e : Exception) {
            return "2025-03-09T13:52:54.823Z"
        }
    }
}