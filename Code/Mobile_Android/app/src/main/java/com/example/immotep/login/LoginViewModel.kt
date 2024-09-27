package com.example.immotep.login

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.immotep.ApiClient.ApiClient
import com.example.immotep.ApiClient.Post
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response

data class LoginState(
    val email: String = "",
    val password: String = "",
)

class LoginViewModel : ViewModel() {
    private val _postState = MutableStateFlow<Post?>(null)
    private val _emailAndPassword = MutableStateFlow(LoginState())
    val post: StateFlow<Post?> = _postState.asStateFlow()
    val emailAndPassword: StateFlow<LoginState> = _emailAndPassword.asStateFlow()

    fun updateEmailAndPassword(
        email: String?,
        password: String?,
    ) {
        _emailAndPassword.value = _emailAndPassword.value.copy(
            email = if (email != null) { email } else { _emailAndPassword.value.email },
            password = if (password != null) { password } else { _emailAndPassword.value.password }
        )
    }

    init {
        _postState.value = null
        viewModelScope.launch {
            fetchPost()
        }
    }

    private fun fetchPost() {
        ApiClient.apiService.getPostById(1).enqueue(
            object : Callback<Post> {
                override fun onResponse(
                    call: Call<Post>,
                    response: Response<Post>,
                ) {
                    if (response.isSuccessful) {
                        _postState.value = response.body()
                    } else {
                        // Handle error
                    }
                }

                override fun onFailure(
                    call: Call<Post>,
                    t: Throwable,
                ) {
                    // Handle error
                }
            },
        )
    }
}
