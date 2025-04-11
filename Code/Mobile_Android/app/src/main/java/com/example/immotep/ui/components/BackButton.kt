package com.example.immotep.ui.components

import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.ChevronLeft
import androidx.compose.material.icons.outlined.KeyboardArrowLeft
import androidx.compose.material3.Button
import androidx.compose.material3.IconButtonDefaults
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
    IconButton(
        onClick = onClick,
        colors = IconButtonDefaults.iconButtonColors(containerColor = MaterialTheme.colorScheme.background),
        modifier = Modifier.testTag("backButton"),
    ) {
        Icon(Icons.Outlined.ChevronLeft, contentDescription = stringResource(R.string.back), tint = MaterialTheme.colorScheme.onBackground)
    }
}
