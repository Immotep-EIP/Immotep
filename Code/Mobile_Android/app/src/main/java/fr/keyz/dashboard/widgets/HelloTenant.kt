package fr.keyz.dashboard.widgets

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.height
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import fr.keyz.R

@Composable
fun HelloTenant(userName : String) {
    WidgetBase(dropDownItems = arrayOf(), testTag = "helloTenantWidget") {
        Column {
            Text(
                "${stringResource(R.string.welcome)} $userName",
                fontWeight = FontWeight.Bold,
                fontSize = 20.sp
            )
            Spacer(modifier = Modifier.height(5.dp))
            Text(stringResource(R.string.dashboard_tenant_resume))
        }
    }
}