// Top-level build file where you can add configuration options common to all sub-projects/modules.
plugins {
    alias(libs.plugins.android.application) apply false
    alias(libs.plugins.kotlin.android) apply false
    @Suppress("ktlint:standard:no-wildcard-imports")
    id("org.jlleitschuh.gradle.ktlint") version "11.0.0"
}
