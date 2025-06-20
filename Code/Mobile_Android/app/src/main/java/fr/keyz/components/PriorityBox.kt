package fr.keyz.components

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import fr.keyz.R
import fr.keyz.apiCallerServices.Priority
import fr.keyz.utils.ThemeUtils

@Composable
fun PriorityBox(priority: Priority) {
    val color = ThemeUtils.getStatusColor(priority = priority)
    val text = when (priority) {
        Priority.low -> stringResource(R.string.low)
        Priority.medium -> stringResource(R.string.medium)
        Priority.high -> stringResource(R.string.high)
        Priority.urgent -> stringResource(R.string.urgent)
    }
    Box(
        modifier = Modifier
            .padding(3.dp)
            .border(1.dp, color, MaterialTheme.shapes.small)
            .background(color = color.copy(alpha = 0.1f))
    ) {
        Text(text, color = color, modifier = Modifier.padding(3.dp))
    }
}