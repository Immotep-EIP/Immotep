package com.example.keyz.profile

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.IconButton
import androidx.compose.material.MaterialTheme
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.outlined.ExitToApp
import androidx.compose.material3.Icon
import androidx.compose.material3.SegmentedButton
import androidx.compose.material3.SegmentedButtonDefaults
import androidx.compose.material3.SingleChoiceSegmentedButtonRow
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.mutableIntStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.runtime.getValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController
import com.example.keyz.LocalApiService
import com.example.keyz.R
import com.example.keyz.components.ErrorAlert
import com.example.keyz.components.LoadingDialog
import com.example.keyz.dashboard.DashBoardLayout
import java.util.Locale

@Composable
fun UserInfoLine(label : String, value : String) {
    Spacer(modifier = Modifier.height(10.dp))
    Text(label, fontWeight = FontWeight.Thin, fontSize = 10.sp)
    Text(value)
}

@Composable
fun LineBottomColumn(content: @Composable () -> Unit) {
    val borderColor = androidx.compose.material3.MaterialTheme.colorScheme.onBackground
    Column(
        modifier = Modifier
            .fillMaxWidth()
            .padding(2.dp)
            .drawBehind {
                val borderSize = 1.dp.toPx()
                drawLine(
                    color = borderColor,
                    start = Offset(0f, size.height),
                    end = Offset(size.width, size.height),
                    strokeWidth = borderSize
                )
            }
    ) {
        Spacer(modifier = Modifier.height(10.dp))
        content()
        Spacer(modifier = Modifier.height(10.dp))
    }
}


@Composable
fun ProfileScreen(
    navController: NavController,
) {
    val context = LocalContext.current
    val apiService = LocalApiService.current
    val viewModel: ProfileViewModel = viewModel {
        ProfileViewModel(navController, apiService)
    }
    val infos = viewModel.infos.collectAsState()
    val isLoading = viewModel.isLoading.collectAsState()

    val currentLanguage = if (Locale.getDefault().language == "fr") 1 else 0
    var selectedIndex by remember { mutableIntStateOf(currentLanguage) }
    val options = listOf("English", "Français")

    LaunchedEffect(Unit) {
        viewModel.initProfile()
    }
    DashBoardLayout(navController, "profile") {
        LoadingDialog(isLoading.value)
        Column(
            verticalArrangement = Arrangement.Center,
            horizontalAlignment = Alignment.CenterHorizontally,
            modifier = Modifier.fillMaxSize()
        ) {
            ErrorAlert(
                null,
                null,
                if (viewModel.apiError.collectAsState().value) stringResource(R.string.profile_api_error) else null
            )
            Column(
                horizontalAlignment = Alignment.CenterHorizontally,
                modifier = Modifier
                    .fillMaxWidth()
                    .fillMaxHeight()
                    .padding(top = 100.dp, bottom = 100.dp, start = 30.dp, end = 30.dp)
                    .shadow(
                        elevation = 6.dp,
                        shape = RoundedCornerShape(10.dp),
                        clip = false
                    )
                    .clip(RoundedCornerShape(10.dp))
                    .background(color = androidx.compose.material3.MaterialTheme.colorScheme.primaryContainer, shape = RoundedCornerShape(10.dp))
                    .padding(15.dp)
            ) {
                LineBottomColumn {
                    Text(stringResource(R.string.user_infos), fontWeight = FontWeight.Bold)
                    UserInfoLine(stringResource(R.string.first_name), infos.value.firstname)
                    UserInfoLine(stringResource(R.string.last_name), infos.value.lastname)
                    UserInfoLine(stringResource(R.string.email), infos.value.email)
                }
                LineBottomColumn {
                    Row(
                        horizontalArrangement = Arrangement.SpaceBetween,
                        verticalAlignment = Alignment.CenterVertically,
                        modifier = Modifier.fillMaxWidth()
                    ) {
                        Text(stringResource(R.string.language))
                        SingleChoiceSegmentedButtonRow(
                            modifier = Modifier.testTag("selectButtonLanguage")
                        ) {
                            options.forEachIndexed { index, label ->
                                SegmentedButton(
                                    shape = SegmentedButtonDefaults.itemShape(index = index, count = options.size),
                                    onClick = {
                                        selectedIndex = index;
                                        viewModel.changeLanguageAndRestart(context, if (index == 0) "en" else "fr" )
                                    },
                                    selected = index == selectedIndex
                                ) {
                                    Text(label)
                                }
                            }
                        }
                    }
                }
                LineBottomColumn {
                    Row(
                        horizontalArrangement = Arrangement.SpaceBetween,
                        verticalAlignment = Alignment.CenterVertically,
                        modifier = Modifier.fillMaxWidth()
                    ) {
                        Text(stringResource(R.string.logout))
                        IconButton(
                            onClick = { viewModel.logout() },
                            modifier = Modifier.testTag("profileLogoutBtn")
                        ) {
                            Icon(
                                Icons.AutoMirrored.Outlined.ExitToApp,
                                contentDescription = stringResource(R.string.logout),
                                tint = MaterialTheme.colors.error
                            )
                        }
                    }
                }

            }
        }
    }
}
