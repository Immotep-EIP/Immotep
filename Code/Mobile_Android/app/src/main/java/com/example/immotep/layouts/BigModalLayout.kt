package com.example.immotep.layouts

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.setValue
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalConfiguration
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.unit.dp
import kotlinx.coroutines.launch

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun BigModalLayout(height: Float, open : Boolean, close : () -> Unit, testTag : String = "bigModalLayout", content: @Composable () -> Unit)
{
    val sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = true)
    val scope = rememberCoroutineScope()
    var innerShowBottomSheet by rememberSaveable { mutableStateOf(false) }

    val configuration = LocalConfiguration.current
    val screenHeight = configuration.screenHeightDp.dp
    val modalHeight = screenHeight * height
    val closeModal = {
        scope.launch { sheetState.hide() }.invokeOnCompletion {
            if (!sheetState.isVisible) {
                innerShowBottomSheet = false
                close()
            }
        }
    }
    LaunchedEffect(open) {
        if (!open) {
            closeModal()
        } else {
            innerShowBottomSheet = true
            sheetState.show()
        }
    }
    if (innerShowBottomSheet) {
        ModalBottomSheet(
            onDismissRequest = {
                closeModal()
            },
            sheetState = sheetState,
            modifier = Modifier.fillMaxWidth().heightIn(modalHeight).testTag(testTag),
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
}