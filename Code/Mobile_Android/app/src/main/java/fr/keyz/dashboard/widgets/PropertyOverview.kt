package fr.keyz.dashboard.widgets

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import fr.keyz.apiCallerServices.DetailedProperty
import fr.keyz.R
import fr.keyz.apiCallerServices.Priority
import fr.keyz.utils.DateFormatter
import fr.keyz.utils.ThemeUtils

@Composable
fun PropertyOverview(detailedProperty: DetailedProperty) {
    WidgetBase(
        dropDownItems = arrayOf(),
        testTag = "propertyOverview",
        title = stringResource(R.string.my_rental)
    ) {
        Column {
            Text(text = detailedProperty.name, fontWeight = FontWeight.SemiBold, fontSize = 16.sp)
            Spacer(modifier = Modifier.height(4.dp))
            Text(text = "${detailedProperty.address}, ${detailedProperty.city}, ${detailedProperty.country}")
            Spacer(modifier = Modifier.height(4.dp))
            Text(text = "${stringResource(R.string.start_date)} :  ${DateFormatter.formatOffsetDateTime(detailedProperty.lease?.startDate)}")
            Spacer(modifier = Modifier.height(4.dp))
            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceBetween,
                verticalAlignment = Alignment.CenterVertically
            ) {
                WidgetNumberBase(
                    title = stringResource(R.string.rentPerMonth),
                    value = detailedProperty.rent,
                    afterValue = "€",
                )
                WidgetNumberBase(
                    title = stringResource(R.string.area),
                    value = detailedProperty.area,
                    afterValue = "m²",
                )
            }
        }
    }
}