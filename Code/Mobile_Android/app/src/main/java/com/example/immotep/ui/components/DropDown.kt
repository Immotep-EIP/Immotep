package com.example.immotep.ui.components

import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.DropdownMenu
import androidx.compose.material.DropdownMenuItem
import androidx.compose.material.Icon
import androidx.compose.material.MaterialTheme
import androidx.compose.material.Text
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.ArrowDropDown
import androidx.compose.runtime.Composable
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.example.immotep.R

data class DropDownItem<T>(
    val label : String,
    val value : T
)

@Composable
fun <T>DropDown(
    items : List<DropDownItem<T>>,
    selectedItem : T,
    onItemSelected : (T) -> Unit,
    error : String? = null,
    testTag : String = "dropDown"
    ) {
    val isDropDownExpanded = remember {
        mutableStateOf(false)
    }

    Box(modifier = Modifier
        .padding(top = 10.dp)
        .border(
            width = 2.dp,
            color = if (error == null) MaterialTheme.colors.onSurface else MaterialTheme.colors.error,
            shape = RoundedCornerShape(8.dp)
        )
        .padding(10.dp)
        .fillMaxWidth()
        .testTag(testTag),

    ) {
        Row(
            horizontalArrangement = Arrangement.SpaceBetween,
            verticalAlignment = Alignment.CenterVertically,
            modifier = Modifier.clickable {
                isDropDownExpanded.value = true
            }.fillMaxWidth()
        ) {
            Text(text = items.find { it.value == selectedItem }?.label ?: stringResource(R.string.select_an_element))
            Icon(Icons.Outlined.ArrowDropDown, contentDescription = "Drop Down")
        }
        DropdownMenu(
            expanded = isDropDownExpanded.value,
            onDismissRequest = {
                isDropDownExpanded.value = false
            }) {
            items.forEach { item ->
                DropdownMenuItem(
                    content = {
                        Text(text = item.label)
                    },
                    onClick = {
                        isDropDownExpanded.value = false
                        onItemSelected(item.value)
                    })
            }
        }
    }
    if (error != null )  {
        Text(
            error,
            modifier = Modifier.padding(top = 10.dp),
            color = androidx.compose.material3.MaterialTheme.colorScheme.error)
    }
}