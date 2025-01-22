package com.example.immotep.utils

import android.content.ContentUris
import android.content.Context
import android.database.Cursor
import android.graphics.BitmapFactory
import android.net.Uri
import android.os.Build
import android.os.Environment
import android.provider.DocumentsContract
import android.provider.MediaStore
import androidx.compose.ui.graphics.ImageBitmap
import androidx.compose.ui.graphics.asImageBitmap
import java.io.File
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
            withPrefix: Boolean = false
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

        fun decodeBase64ToImage(base64String: String): ImageBitmap? {
            try {
                val byteArrayDecoded =
                    Base64.getDecoder().decode(base64String)
                return BitmapFactory.decodeByteArray(byteArrayDecoded, 0, byteArrayDecoded.size)
                    .asImageBitmap()
            } catch (e: Exception) {
                println("error : ${e.message}")
            }
            return null
        }
    }
}