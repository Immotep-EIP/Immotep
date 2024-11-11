package com.example.immotep.ApiClient

import retrofit2.http.Body
import retrofit2.http.Field
import retrofit2.http.FormUrlEncoded
import retrofit2.http.GET
import retrofit2.http.Header
import retrofit2.http.POST

data class LoginResponse(
    val access_token: String,
    val refresh_token: String,
    val token_type: String,
    val expires_in: Int,
    val properties: Map<String, Any>,
)

data class RegistrationInput(
    val email: String,
    val password: String,
    val firstName: String,
    val lastName: String,
)

data class RegistrationResponse(
    val id: String,
    val email: String,
    val firstname: String,
    val lastname: String,
    val role: String,
    val created_at: String,
    val updated_at: String,
)

data class ProfileResponse(
    val id: String,
    val email: String,
    val firstname: String,
    val lastname: String,
    val role: String,
    val created_at: String,
    val updated_at: String,
)

const val API_PREFIX = "/api/v1"

interface ApiService {
    @FormUrlEncoded
    @POST("${API_PREFIX}/auth/token")
    suspend fun login(
        @Field("grant_type") grantType: String = "password",
        @Field("username") username: String,
        @Field("password") password: String,
    ): LoginResponse

    @FormUrlEncoded
    @POST("${API_PREFIX}/auth/token")
    suspend fun refreshToken(
        @Field("grant_type") grantType: String = "refresh_token",
        @Field("refresh_token") refreshToken: String,
    ): LoginResponse

    @POST("${API_PREFIX}/auth/register")
    suspend fun register(
        @Body registrationInput: RegistrationInput,
    ): RegistrationResponse

    @GET("${API_PREFIX}/profile")
    suspend fun getProfile(@Header("Authorization") authHeader : String): ProfileResponse
}
