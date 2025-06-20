package fr.keyz.dashboard.widgets

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowColumn
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
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
import fr.keyz.apiCallerServices.DashBoardReminder
import fr.keyz.apiCallerServices.Priority

@Composable
fun PriorityBox(priority: Priority) {
    val color = when (priority) {
        Priority.low -> Color(0xFF90CAF9)
        Priority.medium -> Color(0xFFFFF176)
        Priority.high -> Color(0xFFFF9862)
        Priority.urgent -> Color(0xFFF44336)
    }
    val text = when (priority) {
        Priority.low -> stringResource(R.string.low)
        Priority.medium -> stringResource(R.string.medium)
        Priority.high -> stringResource(R.string.high)
        Priority.urgent -> stringResource(R.string.urgent)
    }
    Box(
        modifier = Modifier
            .padding(3.dp)
            .border(1.dp, color, MaterialTheme.shapes.small)
            .background(color = color.copy(alpha = 0.1f))
    ) {
        Text(text, color = color, modifier = Modifier.padding(3.dp))
    }
}

@Composable
fun OneReminder(reminder : DashBoardReminder, isLast : Boolean) {
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
        Text(reminder.title, color = MaterialTheme.colorScheme.primary)
    }
}

@OptIn(ExperimentalLayoutApi::class)
@Composable
fun RemindersWidget(reminders : Array<DashBoardReminder>) {
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
            //sort the array by gravity
            reminders.take(5).forEachIndexed { index, reminder ->
                OneReminder(reminder, index == reminders.size - 1 || index == 4 )
            }
        }
    }
}