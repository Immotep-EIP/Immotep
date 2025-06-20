# DateFormatter

## Overview

`DateFormatter` is a utility object providing date and time formatting functions tailored for handling `OffsetDateTime` instances and generating the current date-time string in a specific ISO format.

---

## Functions

### `formatOffsetDateTime(dateTime: OffsetDateTime?): String?`

Formats an `OffsetDateTime` instance to a string in `yyyy/MM/dd` format.

* If `dateTime` is `null`, returns `null`.
* Uses the standard ISO local date format (`yyyy-MM-dd`) and replaces dashes (`-`) with slashes (`/`).
* Catches exceptions and returns `"Invalid date"` if formatting fails.

---

### `currentDateAsOffsetDateTimeString(): String`

Returns the current date and time as a string formatted like `yyyy-MM-dd'T'HH:mm:ss.SSSXXX`.

* Uses `SimpleDateFormat` with the pattern to produce an ISO 8601–like timestamp with milliseconds and timezone.
* On failure, returns a fixed fallback timestamp string (`"2025-03-09T13:52:54.823Z"`).

---

## Summary

* Provides safe, reusable formatting for date-time values in the app.
* Handles null inputs and formatting errors gracefully.
* Useful for standardizing date output in UI or API payloads.
