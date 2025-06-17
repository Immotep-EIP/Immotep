import java.util.Locale

plugins {
    alias(libs.plugins.android.application)
    alias(libs.plugins.kotlin.android)
    jacoco
}

val cameraXVersion = "1.0.1"

android {
    namespace = "com.example.keyz"
    compileSdk = 35
    compileSdk = 35

    defaultConfig {
        applicationId = "com.example.keyz"
        minSdk = 27
        targetSdk = 35
        targetSdk = 35
        versionCode = 1
        versionName = "1.0"

        testInstrumentationRunner = "androidx.test.runner.AndroidJUnitRunner"
        vectorDrawables {
            useSupportLibrary = true
        }
    }

    buildTypes {
        release {
            isMinifyEnabled = false
            proguardFiles(
                getDefaultProguardFile("proguard-android-optimize.txt"),
                "proguard-rules.pro",
                "proguard-rules.pro",
            )
        }
        debug {
            enableAndroidTestCoverage = true
            enableUnitTestCoverage = true
        }
    }
    compileOptions {
        sourceCompatibility = JavaVersion.VERSION_1_8
        targetCompatibility = JavaVersion.VERSION_1_8
    }
    kotlinOptions {
        jvmTarget = "1.8"
    }
    buildFeatures {
        compose = true
    }
    composeOptions {
        kotlinCompilerExtensionVersion = "1.5.1"
    }
    packaging {
        resources {
            excludes += "/META-INF/{AL2.0,LGPL2.1}"
        }
    }
}

dependencies {

    implementation(libs.androidx.core.ktx)
    implementation(libs.androidx.lifecycle.runtime.ktx)
    implementation(libs.androidx.activity.compose)
    implementation(platform(libs.androidx.compose.bom))
    implementation(libs.androidx.ui)
    implementation(libs.androidx.ui.graphics)
    implementation(libs.androidx.ui.tooling.preview)
    implementation(libs.androidx.material3)
    implementation("androidx.compose.ui:ui:1.5.4")
    implementation("androidx.compose.ui:ui-graphics:1.5.4")
    implementation("androidx.compose.ui:ui-tooling-preview:1.5.4")
    implementation("androidx.compose.material3:material3:1.3.2")

    //pdf handler
    implementation("io.github.grizzi91:bouquet:1.1.2")


    //retrofit
    implementation("com.squareup.retrofit2:retrofit:2.9.0")
    implementation("com.squareup.retrofit2:converter-gson:2.9.0")
    implementation("androidx.navigation:navigation-compose:2.8.4")
    implementation("com.squareup.okhttp3:logging-interceptor:4.12.0")

    implementation("com.google.accompanist:accompanist-systemuicontroller:0.27.0")
    implementation("androidx.compose.material:material:1.8.0-alpha01")
    implementation("androidx.datastore:datastore-preferences:1.0.0")
    implementation("io.github.osipxd:security-crypto-datastore-preferences:1.0.0-alpha04")
    implementation("androidx.security:security-crypto-ktx:1.1.0-alpha05")
    implementation("io.github.cdimascio:dotenv-kotlin:6.4.2")
    implementation("com.squareup.okhttp3:okhttp:4.12.0")
    implementation("io.coil-kt:coil-compose:2.4.0")
    implementation("androidx.compose.material:material-icons-extended:1.5.4")

    //camer
    implementation("androidx.camera:camera-camera2:1.0.1")
    implementation("androidx.camera:camera-lifecycle:1.0.1")
    implementation("androidx.camera:camera-view:1.0.0-alpha27")

    // Needed for unit testing API
    testImplementation("androidx.arch.core:core-testing:2.1.0")
    testImplementation("io.mockk:mockk:1.13.9")
    testImplementation("org.jetbrains.kotlinx:kotlinx-coroutines-test:1.8.0")
    androidTestImplementation("io.github.aungthiha:compose-ui-test:1.0.1")
    testImplementation("org.junit.jupiter:junit-jupiter-api:5.12.0")
    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine:5.12.0")
    androidTestImplementation(libs.androidx.espresso.core)
    androidTestImplementation(platform(libs.androidx.compose.bom))
    androidTestImplementation(libs.androidx.ui.test.junit4)
    debugImplementation(libs.androidx.ui.tooling)
    debugImplementation(libs.androidx.ui.test.manifest)
}

val exclusions = listOf(
    "**/R.class",
    "**/R\$*.class",
    "**/BuildConfig.*",
    "**/Manifest*.*",
    "**/*Test*.*"
)

tasks.withType(Test::class) {
    configure<JacocoTaskExtension> {
        isIncludeNoLocationClasses = true
        excludes = listOf("jdk.internal.*")
    }
}

android {
    applicationVariants.all(
        closureOf<com.android.build.gradle.internal.api.BaseVariantImpl> {
            val variant = this@closureOf.name.replaceFirstChar {
                if (it.isLowerCase()) it.titlecase(
                    Locale.getDefault()
                ) else it.toString()
            }
            val unitTests = "test${variant}UnitTest"
            val androidTests = "connected${variant}AndroidTest"
            tasks.register<JacocoReport>("Jacoco${variant}CodeCoverage") {
                dependsOn(listOf(unitTests, androidTests))
                group = "Reporting"
                description = "Execute ui and unit tests, generate and combine Jacoco coverage report"
                reports {
                    xml.required.set(true)
                    html.required.set(true)
                }
                sourceDirectories.setFrom(layout.projectDirectory.dir("src/main"))
                classDirectories.setFrom(
                    files(
                        fileTree(layout.buildDirectory.dir("intermediates/javac/")) {
                            exclude(exclusions)
                        },
                        fileTree(layout.buildDirectory.dir("tmp/kotlin-classes/")) {
                            exclude(exclusions)
                        }
                    )
                )
                executionData.setFrom(
                    files(
                        fileTree(layout.buildDirectory) { include(listOf("**/*.exec", "**/*.ec")) }
                    )
                )
            }
        }
    )
}
