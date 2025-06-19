package fr.keyz.utils

import android.content.ContentUris
import android.content.Context
import android.database.Cursor
import android.graphics.BitmapFactory
import android.net.Uri
import android.os.Build
import android.os.Environment
import android.provider.DocumentsContract
import android.provider.MediaStore
import android.provider.OpenableColumns
import androidx.compose.ui.graphics.ImageBitmap
import androidx.compose.ui.graphics.asImageBitmap
import java.io.File
import java.io.FileOutputStream
import java.util.Base64


class Base64Utils {
    companion object {
        private fun getPathFromUri(uri: Uri, context: Context): String? {
            val isKitKat = Build.VERSION.SDK_INT >= Build.VERSION_CODES.KITKAT

            // DocumentProvider
            if (isKitKat && DocumentsContract.isDocumentUri(context, uri)) {
                // ExternalStorageProvider
                if (isExternalStorageDocument(uri)) {
                    val docId = DocumentsContract.getDocumentId(uri)
                    val split = docId.split(":".toRegex()).dropLastWhile { it.isEmpty() }
                        .toTypedArray()
                    val type = split[0]

                    if ("primary".equals(type, ignoreCase = true)) {
                        return Environment.getExternalStorageDirectory().toString() + "/" + split[1]
                    }

                    // TODO handle non-primary volumes
                } else if (isDownloadsDocument(uri)) {
                    val id = DocumentsContract.getDocumentId(uri)
                    val contentUri = ContentUris.withAppendedId(
                        Uri.parse("content://downloads/public_downloads"), id.toLong()
                    )

                    return getDataColumn(context, contentUri, null, null)
                } else if (isMediaDocument(uri)) {
                    val docId = DocumentsContract.getDocumentId(uri)
                    val split = docId.split(":".toRegex()).dropLastWhile { it.isEmpty() }
                        .toTypedArray()
                    val type = split[0]

                    var contentUri: Uri? = null
                    if ("image" == type) {
                        contentUri = MediaStore.Images.Media.EXTERNAL_CONTENT_URI
                    } else if ("video" == type) {
                        contentUri = MediaStore.Video.Media.EXTERNAL_CONTENT_URI
                    } else if ("audio" == type) {
                        contentUri = MediaStore.Audio.Media.EXTERNAL_CONTENT_URI
                    }

                    val selection = "_id=?"
                    val selectionArgs = arrayOf(
                        split[1]
                    )

                    return getDataColumn(context, contentUri, selection, selectionArgs)
                }
            } else if ("content".equals(uri.scheme, ignoreCase = true)) {
                return getDataColumn(context, uri, null, null)
            } else if ("file".equals(uri.scheme, ignoreCase = true)) {
                return uri.path
            }

            return null
        }

        private fun getDataColumn(
            context: Context, uri: Uri?, selection: String?,
            selectionArgs: Array<String>?
        ): String? {
            var cursor: Cursor? = null
            val column = "_data"
            val projection = arrayOf(
                column
            )

            try {
                cursor = context.contentResolver.query(
                    uri!!, projection, selection, selectionArgs,
                    null
                )
                if (cursor != null && cursor.moveToFirst()) {
                    val column_index = cursor.getColumnIndexOrThrow(column)
                    return cursor.getString(column_index)
                }
            } finally {
                cursor?.close()
            }
            return null
        }

        private fun isExternalStorageDocument(uri: Uri): Boolean {
            return "com.android.externalstorage.documents" == uri.authority
        }

        private fun isDownloadsDocument(uri: Uri): Boolean {
            return "com.android.providers.downloads.documents" == uri.authority
        }

        private fun isMediaDocument(uri: Uri): Boolean {
            return "com.android.providers.media.documents" == uri.authority
        }

        fun encodeImageToBase64(
            fileUri: Uri,
            context: Context,
            withPrefix: Boolean = true
        ): String {
            try {
                val path = getPathFromUri(fileUri, context) ?: return ""
                val file = File(path)
                val bytes = file.readBytes()
                val base64 = Base64.getEncoder().encodeToString(bytes)
                if (!withPrefix) {
                    return base64
                }
                return "data:image/${file.extension};base64,$base64"
            } catch (e : Exception) {
                println("error : ${e.message}")
                return ""
            }
        }

        fun decodeBase64ToImage(originalBase64String: String): ImageBitmap? {
            try {
                var base64String = originalBase64String
                if (base64String.startsWith("data:image")) {
                    base64String = base64String.substringAfter(",")
                }
                val byteArrayDecoded =
                    Base64.getDecoder().decode(base64String)
                return BitmapFactory.decodeByteArray(byteArrayDecoded, 0, byteArrayDecoded.size)
                    .asImageBitmap()
            } catch (e: Exception) {
                println("error : ${e.message}")
            }
            return null
        }

        fun saveBase64PdfToCache(context: Context, base64String: String, name : String? = null): File? {
            return try {
                val pdfName = if (name != null) "$name.pdf" else "temp_pdf.pdf"
                val newBase64String = base64String.replace("data:application/pdf;base64,", "")
                val pdfAsBytes = Base64.getDecoder().decode(newBase64String)

                val pdfFile = File(context.cacheDir, pdfName)
                FileOutputStream(pdfFile).use { it.write(pdfAsBytes) }

                pdfFile
            } catch (e: Exception) {
                println("Impossible to save pdf file ${e.message}")
                e.printStackTrace()
                null
            }
        }

        fun getFileNameFromUri(context: Context, uri: Uri): String? {
            val contentResolver = context.contentResolver
            var name: String? = null

            if (uri.scheme == "content") {
                val cursor = contentResolver.query(uri, null, null, null, null)
                cursor?.use {
                    val nameIndex = it.getColumnIndex(OpenableColumns.DISPLAY_NAME)
                    if (it.moveToFirst() && nameIndex != -1) {
                        name = it.getString(nameIndex)
                    }
                }
            }

            if (name == null) {
                name = uri.path?.substringAfterLast('/')
            }

            if (name != null && name!!.endsWith(".pdf")) {
                name = name!!.substringBeforeLast(".pdf")
            }

            return name
        }

        fun convertPdfUriToBase64(
            context: Context,
            pdfUri: Uri,
            withPrefix: Boolean = false
        ): String? {
            try {
                val inputStream = context.contentResolver.openInputStream(pdfUri) ?: throw Exception("Input stream is null")
                val bytes = inputStream.readBytes()
                val base64 = Base64.getEncoder().encodeToString(bytes)
                if (!withPrefix) {
                    return base64
                }
                return "data:application/pdf;base64,$base64"
            } catch (e : Exception) {
                println("error during conversion of pdf file : ${e.message}")
                e.printStackTrace()
                return null
            }
        }
    }
}