package com.example.immotep.login

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.lifecycle.viewmodel.compose.viewModel
import androidx.navigation.NavController

@Composable
fun LoginScreen(
    navController: NavController,
    viewModel: LoginViewModel = viewModel(),
) {
    val emailAndPassword = viewModel.emailAndPassword.collectAsState()
    Column(
        modifier = Modifier.fillMaxSize(),
        verticalArrangement = Arrangement.Center,
        horizontalAlignment = Alignment.CenterHorizontally,
    ) {
        Text("Login Screen")
        TextField(label = { Text("Entrez votre email") }, value = emailAndPassword.value.email, onValueChange = { value ->
            viewModel.updateEmailAndPassword(value, null)
        })
        TextField(label = { Text("Entrez votre mot de passe") }, value = emailAndPassword.value.password, onValueChange = { value ->
            viewModel.updateEmailAndPassword(null, value)
        })
        Button(onClick = { navController.navigate("dashboard") }) { (Text("Se connecter")) }
    }
}
