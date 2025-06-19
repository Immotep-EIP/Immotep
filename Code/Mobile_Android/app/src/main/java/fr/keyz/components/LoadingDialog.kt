package fr.keyz.components

import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.width
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import androidx.compose.ui.window.Dialog

@Composable
fun LoadingDialog(isOpen : Boolean) {
    if (!isOpen) {
        return
    }
    Dialog(onDismissRequest = { }) {
            CircularProgressIndicator(
                modifier = Modifier.width(100.dp).height(100.dp),
                color = MaterialTheme.colorScheme.secondary,
                trackColor = MaterialTheme.colorScheme.onSurfaceVariant,
            )
    }
}