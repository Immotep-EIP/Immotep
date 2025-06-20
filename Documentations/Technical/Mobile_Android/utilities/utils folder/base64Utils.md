# Base64Utils

## Overview

`Base64Utils` is a utility class providing helper functions to handle file and image encoding/decoding with Base64, specifically tailored for Android URIs and content. It supports converting files from URIs to Base64 strings, decoding Base64 back to images, saving Base64 PDFs to cache, and extracting file names from URIs.

---

## Functions

### `getPathFromUri(uri: Uri, context: Context): String?`

Resolves the absolute file system path for a given Android `Uri`.

* Handles various URI authorities including:

  * External storage documents
  * Downloads provider
  * Media documents (images, video, audio)
  * Content and file schemes
* Returns the file path as a string or `null` if it cannot be resolved.

*Private helper used internally.*

---

### `getDataColumn(context: Context, uri: Uri?, selection: String?, selectionArgs: Array<String>?): String?`

Queries the content resolver for the data column (`_data`) corresponding to the given URI.

* Used to extract file paths from content URIs.
* Returns file path string or `null`.

*Private helper used internally.*

---

### `isExternalStorageDocument(uri: Uri): Boolean`

Checks if the URI authority corresponds to external storage documents.

---

### `isDownloadsDocument(uri: Uri): Boolean`

Checks if the URI authority corresponds to downloads documents.

---

### `isMediaDocument(uri: Uri): Boolean`

Checks if the URI authority corresponds to media documents.

---

### `encodeImageToBase64(fileUri: Uri, context: Context, withPrefix: Boolean = true): String`

Encodes an image file located at the provided `Uri` into a Base64 string.

* Resolves file path from `Uri`.
* Reads bytes from the file and encodes to Base64.
* Adds data URI prefix (`data:image/extension;base64,`) by default if `withPrefix` is true.
* Returns empty string if any error occurs.

---

### `decodeBase64ToImage(originalBase64String: String): ImageBitmap?`

Decodes a Base64-encoded image string into an `ImageBitmap` usable by Compose.

* Strips the data URI prefix if present.
* Decodes Base64 to bytes and converts to bitmap.
* Returns `null` if decoding fails.

---

### `saveBase64PdfToCache(context: Context, base64String: String, name: String? = null): File?`

Saves a Base64-encoded PDF string to a temporary file in the app cache directory.

* Strips PDF Base64 prefix (`data:application/pdf;base64,`).
* Decodes Base64 to bytes and writes to file.
* Filename defaults to `temp_pdf.pdf` or uses provided `name`.
* Returns the saved `File` or `null` on failure.

---

### `getFileNameFromUri(context: Context, uri: Uri): String?`

Extracts a user-friendly file name from a content or file `Uri`.

* Queries the content resolver for display name if available.
* Falls back to extracting the last path segment.
* Removes `.pdf` extension if present.
* Returns the file name or `null` if not found.

---

### `convertPdfUriToBase64(context: Context, pdfUri: Uri, withPrefix: Boolean = true): String?`

Converts a PDF file located at the given `Uri` into a Base64 string.

* Opens input stream from content resolver.
* Reads bytes and encodes to Base64.
* Adds `data:application/pdf;base64,` prefix if `withPrefix` is true.
* Returns `null` on error.

---

## Usage Summary

`Base64Utils` simplifies:

* Converting Android URIs for images and PDFs to Base64-encoded strings (with or without data URI prefixes).
* Decoding Base64 image strings back into Compose `ImageBitmap`.
* Saving Base64 PDF strings as physical files in cache for temporary use.
* Resolving file paths and file names from URIs for further processing.

This utility handles common Android URI and file access patterns, abstracting away the complexity of content provider queries and file handling.
