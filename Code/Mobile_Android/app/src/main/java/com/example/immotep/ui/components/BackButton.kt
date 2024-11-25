package com.example.immotep.ui.components

import androidx.compose.foundation.layout.padding
import androidx.compose.material.Icon
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.KeyboardArrowLeft
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.example.immotep.R

@Composable
fun BackButton(onClick: () -> Unit) {
    Button(onClick = onClick, modifier = Modifier.testTag("backButton")) {
        Icon(Icons.Outlined.KeyboardArrowLeft, contentDescription = stringResource(R.string.back), tint = MaterialTheme.colorScheme.onPrimary)
        Text(stringResource(R.string.back), modifier = Modifier.padding(start = 10.dp))
    }
}
