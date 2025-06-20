# PdfsUtils

## Overview

`PdfsUtils` is a utility object designed to simplify the process of opening PDF files on Android devices using a secure `FileProvider` URI.

---

## Functions

### `openPdfFile(context: Context, pdfFile: File)`

* Opens the given PDF file using an external app that can handle PDFs.
* Generates a content URI for the file via `FileProvider` to comply with Android's file access security.
* Creates an `Intent` with action `ACTION_VIEW` configured for PDF mime type.
* Adds the flag `FLAG_GRANT_READ_URI_PERMISSION` to allow temporary read access.
* Launches a chooser dialog to let the user pick their preferred PDF viewer app.

---

## Summary

* Provides a secure and user-friendly way to open PDF files from the app.
* Handles Android file sharing restrictions by leveraging `FileProvider`.
* Ensures compatibility with various PDF viewer apps installed on the device.
