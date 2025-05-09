package com.example.keyz.realProperty.details.tabs

import android.net.Uri
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.defaultMinSize
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.wrapContentSize
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.Button
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.AttachFile
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import com.example.keyz.R
import com.example.keyz.apiCallerServices.Document

@Composable
fun OneDocument(document: Document, openPdf: (String) -> Unit) {
    Column(
        horizontalAlignment = Alignment.CenterHorizontally,
        modifier = Modifier
            .fillMaxWidth(0.33f)
            .padding(5.dp, top = 32.dp)
            .clickable {
                openPdf(document.id)
            }
            .wrapContentSize(Alignment.Center)
            .testTag("OneDocument ${document.id}")
    ) {
        Box(
            modifier = Modifier
                .border(
                    1.dp,
                    color = MaterialTheme.colorScheme.onBackground,
                    shape = RoundedCornerShape(5.dp)
                )
                .padding(start = 25.dp, end = 25.dp, top = 10.dp, bottom = 10.dp)
        ) {
            Icon(
                Icons.Outlined.AttachFile,
                contentDescription = "document icon",
                modifier = Modifier.size(50.dp),
                tint = MaterialTheme.colorScheme.onPrimaryContainer
            )
        }
        Text(
            text = document.name,
            textAlign = TextAlign.Center,
            modifier = Modifier.padding(start = 10.dp, end = 10.dp).fillMaxWidth(),
            maxLines = 2,
            overflow = TextOverflow.Ellipsis,
            color = MaterialTheme.colorScheme.onPrimaryContainer
        )
    }
}

@OptIn(ExperimentalLayoutApi::class)
@Composable
fun DocumentBox(
    openPdf: (String) -> Unit,
    documents : List<Document>,
    addDocument: (Uri) -> Unit
) {
    val pdfLauncher = rememberLauncherForActivityResult(
        contract = ActivityResultContracts.OpenDocument(),
        onResult = { uri ->
            if (uri != null) {
                addDocument(uri)
            }
        }
    )
    Column(
        modifier = Modifier
            .fillMaxSize()
            .verticalScroll(rememberScrollState())
            .testTag("realPropertyDocumentsTab")
    ) {
        FlowRow(modifier = Modifier.defaultMinSize(minHeight = 125.dp)) {
            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.End
            ) {
                Button(
                    shape = RoundedCornerShape(5.dp),
                    colors = androidx.compose.material.ButtonDefaults.buttonColors(
                        backgroundColor = MaterialTheme.colorScheme.secondary,
                        contentColor = MaterialTheme.colorScheme.onPrimaryContainer
                    ),
                    modifier = Modifier.testTag("editInventoryButton"),
                    onClick = { pdfLauncher.launch(arrayOf("application/pdf")) }) {
                    androidx.compose.material.Text(stringResource(R.string.add_document))
                }
            }
            documents.forEach { item ->
                OneDocument(item, openPdf = { openPdf(it) })
            }
        }
    }
}