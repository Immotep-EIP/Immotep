package com.example.immotep.ui.components

import androidx.compose.foundation.interaction.MutableInteractionSource
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.requiredHeight
import androidx.compose.foundation.text.KeyboardActions
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material3.LocalTextStyle
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.remember
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.input.VisualTransformation
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp

@Composable
fun OutlinedTextField(
    value: String,
    onValueChange: (String) -> Unit,
    modifier: Modifier = Modifier,
    enabled: Boolean = true,
    readOnly: Boolean = false,
    label: String,
    placeholder: @Composable (() -> Unit)? = null,
    leadingIcon: @Composable (() -> Unit)? = null,
    trailingIcon: @Composable (() -> Unit)? = null,
    helperMessage: String? = null,
    errorMessage: String? = null,
    visualTransformation: VisualTransformation = VisualTransformation.None,
    keyboardOptions: KeyboardOptions = KeyboardOptions.Default,
    keyboardActions: KeyboardActions = KeyboardActions(),
    singleLine: Boolean = false,
    maxLines: Int = Int.MAX_VALUE,
    interactionSource: MutableInteractionSource = remember { MutableInteractionSource() },
) {
    Column {
        OutlinedTextField(
            modifier = modifier,
            value = value,
            onValueChange = onValueChange,
            enabled = enabled,
            readOnly = readOnly,
            label = {
                Text(
                    text = label,
                    color = MaterialTheme.colorScheme.onSurface
                )
                    },
            placeholder = placeholder,
            leadingIcon = leadingIcon,
            trailingIcon = trailingIcon,
            isError = errorMessage != null,
            visualTransformation = visualTransformation,
            keyboardOptions = keyboardOptions,
            keyboardActions = keyboardActions,
            singleLine = singleLine,
            maxLines = maxLines,
            interactionSource = interactionSource,
        )

        Box(
            modifier =
                Modifier
                    .requiredHeight(16.dp)
                    .padding(start = 16.dp, end = 12.dp),
        ) {
            CompositionLocalProvider(
                LocalTextStyle provides
                    LocalTextStyle.current.copy(
                        fontSize = 12.sp,
                        color = if (errorMessage != null) MaterialTheme.colorScheme.error else LocalTextStyle.current.color,
                    ),
            ) {
                if (errorMessage != null) {
                    Text(errorMessage)
                } else if (helperMessage != null) {
                    Text(helperMessage)
                }
            }
        }
    }
}
