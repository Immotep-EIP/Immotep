package fr.keyz.components


import androidx.compose.foundation.BorderStroke
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.Button
import androidx.compose.material.Icon
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.AddCircleOutline
import androidx.compose.material3.MaterialTheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.unit.dp

@Composable
fun InventoryCenterAddButton(onClick: () -> Unit, testTag: String) {
    Button(onClick = onClick,
        border = BorderStroke(1.dp, MaterialTheme.colorScheme.primary),
        shape = RoundedCornerShape(5.dp),
        colors = androidx.compose.material.ButtonDefaults.buttonColors(backgroundColor = MaterialTheme.colorScheme.background),
        modifier = Modifier
            .padding(top = 10.dp, bottom = 10.dp)
            .fillMaxWidth()
            .testTag(testTag)
    )  {
        Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.Center, verticalAlignment = Alignment.CenterVertically) {
            Icon(
                Icons.Outlined.AddCircleOutline,
                contentDescription = "Add Icon",
                tint = MaterialTheme.colorScheme.onPrimaryContainer
            )
        }
    }
}