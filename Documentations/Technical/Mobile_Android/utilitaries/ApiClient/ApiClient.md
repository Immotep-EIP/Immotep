## RetrofitClient and ApiClient Utility Classes Documentation

The `RetrofitClient` and `ApiClient` utility classes provide centralized, reusable configurations for HTTP networking in the Immotep application using Retrofit. By centralizing these configurations, the app can communicate with RESTful APIs consistently and efficiently.

### Overview

1. **`RetrofitClient`**: Configures the Retrofit instance with a base URL and a JSON converter for handling API requests.
2. **`ApiClient`**: Initializes and provides an instance of the `ApiService` interface, allowing the app to make HTTP requests to defined endpoints.

### Class Details

#### 1. `RetrofitClient`

* **Purpose**:
  - Provides a single Retrofit instance configured with the base URL and JSON deserialization.

* **Properties**:
  - **`BASE_URL`**: A constant URL defining the base endpoint for the API. 
    - This URL can be switched between different endpoints based on the environment (local or production).
  - **`retrofit`**: A `Retrofit` instance that initializes lazily, meaning it only gets created when first accessed.

* **Configuration**:
  - **`baseUrl(BASE_URL)`**: Sets the base URL for the HTTP client.
  - **`addConverterFactory(GsonConverterFactory.create())`**: Adds a Gson converter to automatically convert JSON responses into Kotlin objects.

* **Lazy Initialization**:
  - The `retrofit` instance is created only when needed, optimizing memory usage.

* **Usage**:
  - Access `RetrofitClient.retrofit` to obtain a configured Retrofit instance.

#### 2. `ApiClient`

* **Purpose**:
  - Provides an instance of `ApiService`, which defines the endpoints and methods for network requests.

* **Properties**:
  - **`apiService`**: A lazily initialized property that represents the API service, using the `RetrofitClient` to create an instance of `ApiService`.

* **Usage**:
  - Access `ApiClient.apiService` to invoke any defined methods in `ApiService` for network calls.

### Example Usage

To make an API call using `ApiClient`, define your API endpoints in `ApiService`, then call them through `ApiClient.apiService`.

```kotlin
// Example API Call
suspend fun fetchExampleData() {
    val response = ApiClient.apiService.getExampleData()
    if (response.isSuccessful) {
        val data = response.body()
        // Use the data here
    } else {
        // Handle error
    }
}
```

### Benefits of Using `RetrofitClient` and `ApiClient`

- **Single Source of Truth**: Centralizing API configuration ensures consistent network settings and simplifies maintenance.
- **Lazy Initialization**: Optimizes resources by only creating instances when needed.
- **Scalability**: Additional API clients can be easily added or modified by updating the base URL or other configurations in one place.

### Testing and Debugging

* **Mocking API Responses**: Use mocking tools or libraries like MockWebServer to simulate API responses when testing functions that depend on `ApiClient`.
* **Swappable `BASE_URL`**: The `BASE_URL` can be changed to target different environments, which is useful for local testing or staging environments.

