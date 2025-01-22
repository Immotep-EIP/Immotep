package com.example.immotep.components

import androidx.compose.animation.AnimatedVisibility
import androidx.compose.animation.AnimatedVisibilityScope
import androidx.compose.animation.core.tween
import androidx.compose.animation.fadeIn
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.getValue
import androidx.compose.runtime.setValue
import androidx.compose.runtime.saveable.rememberSaveable


@Composable
internal fun InitialFadeIn(durationMs: Int = 1000, content: @Composable() AnimatedVisibilityScope.() -> Unit) {
    var visibility by rememberSaveable { mutableStateOf(false) }
    LaunchedEffect(key1 = Unit, block = { visibility = true })
    AnimatedVisibility(visible = visibility, enter = fadeIn(tween(durationMs)), content = content)
}