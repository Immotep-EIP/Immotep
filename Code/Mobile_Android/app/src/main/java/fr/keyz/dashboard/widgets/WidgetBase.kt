package fr.keyz.dashboard.widgets

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.MoreVert
import androidx.compose.material3.DropdownMenu
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.IconButtonDefaults
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import fr.keyz.realProperty.details.RealPropertyDropDownMenuItem

data class WidgetMenuItem(
    val label : String,
    val onClick : (() -> Unit)?,
    val disabled : Boolean = false,
    val testTag : String
)

@Composable
fun WidgetNumberBase(
    title : String,
    value : Int,
    titleColor : Color = MaterialTheme.colorScheme.primary,
    valueColor : Color = MaterialTheme.colorScheme.primary
) {
    Column(horizontalAlignment = Alignment.CenterHorizontally) {
        Text(value.toString(), color = valueColor, fontSize = 24.sp)
        Spacer(modifier = Modifier.height(5.dp))
        Text(title, color = titleColor, fontSize = 18.sp, fontWeight = FontWeight.SemiBold)
    }

}

@Composable
fun WidgetBase(
    title : String? = null,
    isEmpty : Boolean = false,
    dropDownItems : Array<WidgetMenuItem>,
    testTag : String,
    content : @Composable () -> Unit
) {
    var expanded by rememberSaveable { mutableStateOf(false) }
    if (isEmpty) {
        return
    }
    Column(modifier = Modifier.fillMaxWidth().padding(top = 5.dp, bottom = 20.dp, start = 5.dp, end = 5.dp)) {
        if (title != null) {
            Text(title)
        }
        Box(
            modifier = Modifier
                .fillMaxWidth()
                .shadow(
                    elevation = 6.dp,
                    shape = RoundedCornerShape(10.dp),
                    clip = false
                )
                .clip(RoundedCornerShape(10.dp))
                .background(color = MaterialTheme.colorScheme.primaryContainer, shape = RoundedCornerShape(10.dp))
        ) {
            Box(modifier = Modifier.padding(start = 10.dp, top = 10.dp, bottom = 10.dp, end = 40.dp)) {
                content()
            }
            Box(
                modifier = Modifier
                    .align(Alignment.TopStart)
                    .padding(top = 0.dp, start = 0.dp)
            ) {
                Row(
                    modifier = Modifier.fillMaxWidth(),
                    horizontalArrangement = Arrangement.End,
                    verticalAlignment = Alignment.CenterVertically
                ) {
                    Box {
                        IconButton(
                            onClick = { expanded = true },
                            colors = IconButtonDefaults.iconButtonColors(containerColor = MaterialTheme.colorScheme.background),
                            modifier = Modifier.testTag("moreVertWidget$testTag"),
                        ) {
                            Icon(
                                Icons.Outlined.MoreVert,
                                contentDescription = "More options",
                                tint = MaterialTheme.colorScheme.onBackground
                            )
                        }
                        DropdownMenu(
                            expanded = expanded,
                            onDismissRequest = { expanded = false }
                        ) {
                            dropDownItems.forEach {
                                RealPropertyDropDownMenuItem(
                                    name = it.label,
                                    onClick = it.onClick,
                                    disabled = it.disabled,
                                    closeDropDown = { expanded = false },
                                    testTag = it.testTag
                                )
                            }
                        }
                    }
                }
            }
        }
    }
}