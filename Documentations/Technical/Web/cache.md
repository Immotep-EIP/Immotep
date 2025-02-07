# Technical Documentation: Service Worker for Web Application Caching

## Overview

This application employs a robust caching strategy using a combination of Service Worker and IndexedDB to improve performance, reduce network dependency, and ensure a smooth user experience even in offline scenarios.

---

## Service Worker

A Service Worker is a script that runs in the background of a web browser, separate from the main browser thread. It is primarily used for features like background sync, push notifications, and intercepting network requests to cache resources.

### Features of the Service Worker in Our Application

1. **Caching Static Assets**: The Service Worker pre-caches critical assets such as HTML, CSS, TypeScript, and image files to ensure the application loads quickly.
2. **Runtime Caching**: Dynamic content and API responses are cached on-the-fly to optimize data fetching.
3. **Offline Support**: Users can access the application even without an internet connection by leveraging cached assets.
4. **Cache Versioning**: The Service Worker uses versioned cache keys to manage updates and remove outdated resources automatically.

By intercepting network requests, the Service Worker checks for cached versions of these files before reaching out to the network. This approach significantly reduces loading times and ensures the application remains functional even without an active internet connection.

---

## IndexedDB

IndexedDB is used for caching and managing more sensitive or dynamic data. This includes:

- User information

- Property details

- Other application-specific data

IndexedDB provides a secure and structured way to store and retrieve large amounts of data, making it ideal for managing application state and user data locally.

### Implementation Highlights

**Data Storage:**

- User data and property details are stored in an IndexedDB database with well-defined object stores.

- Data is indexed for efficient querying and retrieval.

**Synchronization:**

- Data is synchronized with the server when the user is online, ensuring consistency.

- When the user is offline, they can still use the app thanks to the cached data.

**Security Measures:** (To do)

- Sensitive data is encrypted before storage.

- Access to IndexedDB is restricted to the application domain.

---

## Workflow Diagram

**User requests an asset:**

- The Service Worker intercepts the request.

- If the request is for a static asset, it checks the cache. If available, the cached version is returned; otherwise, it fetches from the network and updates the cache.

**User requests sensitive or dynamic data:**

- The application queries IndexedDB for the data.

- If the data is available, it is served from the database. If not, a network request is made, and the response is cached in IndexedDB for future use.

---

## Advantages of This Approach

**Performance:**

- Static assets are served instantly from the cache, reducing load times.

- Local storage of dynamic data reduces the need for frequent network requests.

**Offline Support:**

- The application remains functional even when the user is offline.

- User data and application state are preserved locally.

**Scalability:**

- IndexedDB supports large datasets and complex queries, making it suitable for managing growing application data.

---

## Conclusion

This caching strategy effectively combines the strengths of Service Worker and IndexedDB to create a robust and user-friendly application. By caching static assets with the Service Worker, the application ensures fast loading times and offline availability. Meanwhile, IndexedDB securely handles sensitive and dynamic data, allowing users to interact seamlessly with the application even when offline. Together, these tools provide a scalable and efficient solution, enhancing performance and reliability for end-users.
