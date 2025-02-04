# Technical Documentation: Service Worker for Web Application Caching

## Overview

This document explains how the web application utilizes a Service Worker to manage caching efficiently. The Service Worker enables offline functionality, reduces network requests, and improves overall application performance by controlling the caching of assets and API responses.

---

## What is a Service Worker?

A Service Worker is a script that runs in the background of a web browser, separate from the main browser thread. It is primarily used for features like background sync, push notifications, and intercepting network requests to cache resources.

---

## Features of the Service Worker in Our Application

1. **Caching Static Assets**: The Service Worker pre-caches critical assets such as HTML, CSS, JavaScript, and image files to ensure the application loads quickly.
2. **Runtime Caching**: Dynamic content and API responses are cached on-the-fly to optimize data fetching.
3. **Offline Support**: Users can access the application even without an internet connection by leveraging cached assets.
4. **Cache Versioning**: The Service Worker uses versioned cache keys to manage updates and remove outdated resources automatically.

---

## Implementation Details

### Registering the Service Worker

The Service Worker is registered in the `index.js` file:

```javascript
if ("serviceWorker" in navigator) {
  (async () => {
    try {
      const registration = await navigator.serviceWorker.register("/sw.js", {
        scope: "./",
      });
      if (registration.installing) {
        console.log("Service worker installing");
      } else if (registration.waiting) {
        console.log("Service worker installed");
      } else if (registration.active) {
        console.log("Service worker active");
      }
    } catch (error) {
      console.error(`Registration failed with ${error}`);
    }
  })();
}
```

### Cache Name and Assets

```javascript
const CACHE_NAME = "immotep-cache-v1";
const ASSETS_TO_CACHE = ["/", "/index.html"];
```

- CACHE_NAME: The name of the cache storage.

- ASSETS_TO_CACHE: The list of assets to cache during the installation phase.

### Helper Functions

`putInCache`

Stores a resource in the cache.

```javascript
const putInCache = async (request, response) => {
  const cache = await caches.open(CACHE_NAME);
  await cache.put(request, response);
};
```

`cacheFirst`

Implements the Cache First strategy:

1. Check if the resource is available in the cache.
2. If not, try to fetch it from the network.
3. If the network fetch fails, return a fallback response.

```javascript
const cacheFirst = async ({ request, preloadResponsePromise, fallbackUrl }) => {
  const responseFromCache = await caches.match(request);
  if (responseFromCache) {
    return responseFromCache;
  }

  const preloadResponse = await preloadResponsePromise;
  if (preloadResponse) {
    putInCache(request, preloadResponse.clone());
    return preloadResponse;
  }

  try {
    const responseFromNetwork = await fetch(request.clone());
    putInCache(request, responseFromNetwork.clone());
    return responseFromNetwork;
  } catch (error) {
    const fallbackResponse = await caches.match(fallbackUrl);
    if (fallbackResponse) {
      return fallbackResponse;
    }
    return new Response("Network error happened", {
      status: 408,
      headers: { "Content-Type": "text/plain" },
    });
  }
};
```

### Service Worker Events

`install`

Caches specified assets during the installation phase.

```javascript
self.addEventListener("install", (event) => {
  event.waitUntil(
    caches.open(CACHE_NAME).then((cache) => cache.addAll(ASSETS_TO_CACHE))
  );
});
```

`activate`

Activates the service worker and enables navigation preload. Also clears old caches.

```javascript
self.addEventListener("activate", (event) => {
  event.waitUntil(
    Promise.all([
      caches.keys().then((cacheNames) =>
        Promise.all(
          cacheNames.map((cacheName) => {
            if (cacheName !== CACHE_NAME) {
              return caches.delete(cacheName);
            }
            return null;
          })
        )
      ),
      enableNavigationPreload(),
    ])
  );
});

const enableNavigationPreload = async () => {
  if (self.registration.navigationPreload) {
    await self.registration.navigationPreload.enable();
  }
};
```

#### `fetch`

Intercepts network requests and serves responses based on the Cache First strategy.

```javascript
self.addEventListener("fetch", (event) => {
  event.respondWith(
    cacheFirst({
      request: event.request,
      preloadResponsePromise: event.preloadResponse,
      fallbackUrl: {
        status: "offline",
        message:
          "Vous êtes hors ligne. Les données réelles seront affichées dès que vous serez reconnecté.",
        properties: [],
      },
    })
  );
});
```

---

## Key Features

- **Offline Support**: Ensures that essential assets are available offline.
- **Fallback Response**: Provides a user-friendly message when the network is unavailable.
- **Cache Management**: Handles cache updates and old cache cleanup.
- **Navigation Preload**: Improves performance by preloading navigation requests during activation.

---

## Conclusion

This implementation of a service worker with a Cache First strategy enhances user experience by ensuring faster load times and offline functionality. Regular updates to cached assets and efficient cache management are crucial to maintaining optimal performance.
