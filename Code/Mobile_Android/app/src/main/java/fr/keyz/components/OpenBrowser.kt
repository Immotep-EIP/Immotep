package fr.keyz.components

import android.content.Intent
import android.net.Uri
import android.widget.Toast
import androidx.compose.foundation.clickable
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.unit.sp
import fr.keyz.R

@Composable
fun OpenBrowserAnnotatedString(url: String, title: String) {
    val context = LocalContext.current
    val noBrowserFound = stringResource(R.string.no_browser_found)
    Text(
        AnnotatedString(
            title
        ),
        fontSize = 12.sp,
        color = MaterialTheme.colorScheme.secondary,
        modifier =
        Modifier.clickable {
            val browserIntent = Intent(Intent.ACTION_VIEW, Uri.parse(url)).apply {
                addCategory(Intent.CATEGORY_BROWSABLE)
                flags = Intent.FLAG_ACTIVITY_NEW_TASK
            }
            try {
                context.startActivity(browserIntent)
            } catch (e: Exception) {
                Toast.makeText(context, noBrowserFound, Toast.LENGTH_SHORT).show()
            }
       },
    )
}