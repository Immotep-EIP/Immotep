package fr.keyz.dashboard.widgets

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowColumn
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import fr.keyz.R
import fr.keyz.apiCallerServices.Damage
import fr.keyz.apiCallerServices.DashBoardReminder
import fr.keyz.apiCallerServices.Priority
import fr.keyz.components.PriorityBox
import fr.keyz.utils.ThemeUtils

@Composable
fun OneDamagesInWidget(reminder : DashBoardReminder, isLast : Boolean) {
    Row(modifier = Modifier.fillMaxWidth().drawBehind {
        if (!isLast) {
            val y = size.height - 2.dp.toPx() / 2
            drawLine(
                Color.LightGray,
                Offset(0f, y),
                Offset(size.width, y),
                2.dp.toPx()
            )
        }
    }, verticalAlignment = Alignment.CenterVertically) {
        PriorityBox(reminder.priority)
        Spacer(modifier = Modifier.width(5.dp))
        Text(reminder.title, color = MaterialTheme.colorScheme.primary)
    }
}

@OptIn(ExperimentalLayoutApi::class)
@Composable
fun DamagesListWidget(damages : Array<Damage>) {
    var moreImportantDamages = damages.copyOf()
    moreImportantDamages.sortBy { it.priority }
    moreUsefulReminders.reverse()
    moreUsefulReminders = moreUsefulReminders.take(5).toTypedArray()

    WidgetBase(
        title = stringResource(R.string.reminders),
        dropDownItems = arrayOf(),
        testTag = "remindersWidget",
        isEmpty = reminders.isEmpty()
    ) {
        FlowColumn(
            modifier = Modifier
                .testTag("realPropertyDetailsDamagesTab")
                .fillMaxSize()
        ) {
            moreUsefulReminders.forEachIndexed { index, reminder ->
                OneReminder(reminder, index == moreUsefulReminders.size - 1)
            }
        }
    }
}