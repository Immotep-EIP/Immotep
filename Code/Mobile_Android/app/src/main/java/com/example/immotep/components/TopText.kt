package com.example.immotep.components

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp

@Composable
fun TopText(
    title: String,
    subtitle: String,
    limitMarginTop: Boolean = false,
) {
    Column(
        modifier = Modifier.fillMaxWidth().padding(top = if (limitMarginTop) 25.dp else 90.dp),
        horizontalAlignment = Alignment.CenterHorizontally,
    ) {
        Text(title, fontSize = 30.sp, fontWeight = FontWeight.SemiBold, color = MaterialTheme.colorScheme.primary)
        Text(subtitle, fontSize = 15.sp, color = MaterialTheme.colorScheme.primary)
    }
}
