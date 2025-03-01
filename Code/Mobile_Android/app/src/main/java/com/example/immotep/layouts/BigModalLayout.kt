package com.example.immotep.layouts

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalConfiguration
import androidx.compose.ui.unit.dp

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun BigModalLayout(height: Float, open : Boolean, close : () -> Unit, content: @Composable () -> Unit)
{
    val sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = true)
    val configuration = LocalConfiguration.current
    val screenHeight = configuration.screenHeightDp.dp
    val modalHeight = screenHeight * height
    LaunchedEffect(open) {
        if (!open) {
            sheetState.hide()
        } else {
            sheetState.show()
        }
    }
    ModalBottomSheet(
            onDismissRequest = {
                close()
            },
            sheetState = sheetState,
            modifier = Modifier.fillMaxWidth().heightIn(modalHeight),
        ) {
            Column(
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(bottom = 8.dp, top = 0.dp, start = 8.dp, end = 8.dp)
            ) {
                content()
            }
        }
}