package com.example.keyz.components

import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import com.example.keyz.R

@Composable
fun Header() {
    Row(verticalAlignment = Alignment.CenterVertically, modifier = Modifier.testTag("header")) {
        Image(
            painter = painterResource(id = R.drawable.keyz_png_logo_blue),
            contentDescription = stringResource(id = R.string.immotep_logo_desc),
            modifier = Modifier.size(50.dp).padding(end = 10.dp),
        )
        Text(stringResource(R.string.app_name), fontSize = 30.sp, color = MaterialTheme.colorScheme.primary)
    }
}
