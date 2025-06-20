package fr.keyz.dashboard.widgets

import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowColumn
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import fr.keyz.R
import fr.keyz.apiCallerServices.DashBoardReminder

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
                .fillMaxWidth()
        ) {
            //sort the array by gravity
            reminders.take(5).forEach { reminder ->
                Text(reminder.title)
            }
        }
    }
}