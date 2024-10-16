package com.example.immotep.components

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.example.immotep.R
import com.example.immotep.ui.icons.Release_alert

@Composable
fun ErrorAlert(
    code: Int?,
    login: Boolean?,
) {
    if (code == null) {
        return
    }
    val errorText =
        when (code) {
            400 -> stringResource(R.string.bad_request)
            401 -> if (login != null && login) stringResource(R.string.login_error) else stringResource(R.string.unauthorized)
            403 -> stringResource(R.string.forbidden)
            404 -> stringResource(R.string.not_found)
            405 -> stringResource(R.string.method_not_allowed)
            406 -> stringResource(R.string.not_acceptable)
            409 -> stringResource(R.string.conflict)
            410 -> stringResource(R.string.gone)
            413 -> stringResource(R.string.request_entity_too_large)
            415 -> stringResource(R.string.unsupported_media_type)
            429 -> stringResource(R.string.too_many_requests)
            500 -> stringResource(R.string.internal_server_error)
            501 -> stringResource(R.string.not_implemented)
            502 -> stringResource(R.string.bad_gateway)
            503 -> stringResource(R.string.service_unavailable)
            504 -> stringResource(R.string.gateway_timeout)
            else -> stringResource(R.string.unknown_error)
        }

    Row(
        modifier =
            Modifier
                .fillMaxWidth()
                .background(
                    MaterialTheme.colorScheme.errorContainer,
                    shape = RoundedCornerShape(10.dp),
                ).padding(10.dp),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        Image(Release_alert, contentDescription = "Error alert")
        Spacer(modifier = Modifier.width(10.dp))
        Text(errorText, color = Color.White)
    }
}
