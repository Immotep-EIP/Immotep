package fr.keyz.ui.components

import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.unit.dp

@Composable
fun StyledButton(
    onClick: () -> Unit,
    text: String,
    modifier: Modifier = Modifier,
    error: Boolean = false,
    isLoading: Boolean = false,
    testTag: String = "StyledButton"
) {
    Button(
        onClick = if (isLoading) {{}} else onClick,
        colors = ButtonDefaults.buttonColors(
            containerColor = if (error) MaterialTheme.colorScheme.error else MaterialTheme.colorScheme.secondary
        ),
        modifier = modifier
            .clip(RoundedCornerShape(5.dp))
            .padding(5.dp)
            .fillMaxWidth()
            .testTag(testTag)
    ) {
        Text(
            text,
            color = MaterialTheme.colorScheme.onSecondary,
        )
        if (isLoading) {
            Spacer(modifier = Modifier.width(10.dp))
            CircularProgressIndicator(
                modifier = Modifier.width(20.dp).height(20.dp),
                color = MaterialTheme.colorScheme.secondary,
                trackColor = MaterialTheme.colorScheme.onSurfaceVariant,
            )
        }
    }
}