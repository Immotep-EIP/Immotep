package com.example.immotep.ApiClient

import retrofit2.Call
import retrofit2.http.GET
import retrofit2.http.Path

data class Post(
    val userId: Int,
    val id: Int,
    val title: String,
    val body: String,
)

interface ApiService {
    @GET("posts/{id}")
    fun getPostById(
        @Path("id") postId: Int,
    ): Call<Post>
}
