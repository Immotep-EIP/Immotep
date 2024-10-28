# Global overview of the Android Mobile Application

## Technologies used

In order to build this android application, we used the most common and fastest way to build native android applications, [Kotlin](https://www.kotlinlang.org/). We used android studio as an IDE and for build this application. In order to build the UI and make it reactive we used [JetPack Compose](https://www.developer.android.com/compose) and the Model View View Model architecture ([MVVM architecture](https://www.en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93viewmodel)) in order to manage the data and the states used in the application. There are also tools packages that handle all the stuff related to external tools like API handling...


### Versions of technolgies used

- Java : 21.0.4
- Android SDK minimal version : 27
- Android SDK target version : 34
- Kotlin : 1.9.0
- JetPack Compose : 1.5.4


## Pages hierarchy 

All the pages are located in the Mobile_Android/app/src/main/java/com/example/immotep folder and each of them shares the same architecture according to the MVVM architecture. So, there is a folder for each of the pages with their names and each of the pages is a package with at least two files, one for the UI, usually the one that shares the same name with the package name, and another one, who ends with ViewModel, it's the view model. So, the UI page handle all the interface and visible to the end user logic and the ViewModel handles the data and state management, so it's here that the API calls will be made, state variables created... All the pages packages folders start with a lowercase letter.


## Tools packages

The tools packages are the packages that not belongs to specific pages and are just here to implements external packages like [retrofit](https://square.github.io/retrofit/) for API calls. Their folders name start with an uppercase letter and their hierarchy depends on the need of the tool they implement.


## Testing

In this application two ways are present for test and ensure that all the parts of the application works well even if we add new features to it.

### Unit testing

The first way used to test the application is the simple and basic unit testing, we used this way of testing when there is functions and algorithms that are not direct API calls and UI. It's used for testing and ensure that implemented algorithms still done the rights things even if we change the applications. The unit test folder is located in test/java/com/example/immotep. [Unit testing](https://en.wikipedia.org/wiki/Unit_testing)

To run those tests you will need to go the Mobile_Android folder and run :
```
./gradlew test
```


### Instrumented Testing

The other way to test the application, this time in real life simulation situation is by using the instrumented tests. The instrumented test are tests that uses an android emulator to run the application on it in order to test UI, API calls in the application for ensure that the applications is still the same for end user. Those tests differ from the unit test because we check a bigger chunk of the application at each test, it's a part per part testing for the application. The instrumented tests folder is located in androidTest/java/com/example/immotep. [Instrumented testing](https://developer.android.com/training/testing/instrumented-tests)

To run those test you will need an android emulator setup in your android studio [Setup an android emulator](https://developers.google.com/privacy-sandbox/private-advertising/setup/android/download#:~:text=Set%20up%20an%20Android%20device%20emulator%20image,-To%20set%20up&text=In%20Android%20Studio%2C%20go%20to,it%20isn't%20already%20installed.), then you will need to go to the Mobile_Android folder and run : 

```
./gradlew connectedAndroidTests
```
