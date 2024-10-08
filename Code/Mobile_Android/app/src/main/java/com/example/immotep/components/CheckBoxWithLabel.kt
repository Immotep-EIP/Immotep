package com.example.immotep.components

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.material3.Checkbox
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.sp

@Composable
fun CheckBoxWithLabel(
    label: String,
    isChecked: Boolean,
    onCheckedChange: (Boolean) -> Unit,
    errorMessage: String? = null,
    modifier: Modifier = Modifier,
) {
    Column {
        Row(verticalAlignment = Alignment.CenterVertically) {
            Checkbox(
                checked = isChecked,
                onCheckedChange = onCheckedChange,
                modifier = modifier,
            )
            Text(label, color = MaterialTheme.colorScheme.primary, fontSize = 12.sp)
        }
        if (errorMessage != null) {
            Text(errorMessage, color = MaterialTheme.colorScheme.error, fontSize = 10.sp)
        }
    }
}
