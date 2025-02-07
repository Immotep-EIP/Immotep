const version = 1
const CACHE_NAME = `immotep-cache-v1`
const ASSETS_TO_CACHE = ['/', '/index.html', '/vite.svg', '/assets/*']

const SENSITIVE_API_URLS = [
  '/owner/properties/',
  '/auth/token/',
  '/profile/',
  '/user/',
  '/real-property'
]

// eslint-disable-next-line no-restricted-globals
self.addEventListener('install', event => {
  event.waitUntil(
    caches.open(CACHE_NAME).then(cache => cache.addAll(ASSETS_TO_CACHE))
  )
})

// eslint-disable-next-line no-restricted-globals
self.addEventListener('activate', event => {
  const cacheWhitelist = [CACHE_NAME]
  event.waitUntil(
    caches.keys().then(cacheNames =>
      Promise.all(
        // eslint-disable-next-line array-callback-return, consistent-return
        cacheNames.map(cacheName => {
          if (!cacheWhitelist.includes(cacheName)) {
            return caches.delete(cacheName)
          }
        })
      )
    )
  )
})

const cacheFirst = async event => {
  try {
    const cachedResponse = await caches.match(event.request)
    if (cachedResponse) {
      return cachedResponse
    }

    const response = await fetch(event.request)
    if (
      response &&
      response.status === 200 &&
      !SENSITIVE_API_URLS.some(url => event.request.url.includes(url))
    ) {
      const clonedResponse = response.clone()
      const cache = await caches.open(CACHE_NAME)
      await cache.put(event.request, clonedResponse)
    }

    return response
  } catch (error) {
    console.error('Error in cacheFirst:', error)
    return new Response('An error occurred while fetching the resource.', {
      status: 500,
      statusText: 'Internal Server Error'
    })
  }
}

// eslint-disable-next-line no-restricted-globals
self.addEventListener('fetch', event => {
  if (SENSITIVE_API_URLS.some(url => event.request.url.includes(url))) {
    event.respondWith(fetch(event.request))
  } else {
    event.respondWith(cacheFirst(event))
  }
})

const deleteCache = async () => {
  const cacheNames = await caches.keys()
  await Promise.all(cacheNames.map(cacheName => caches.delete(cacheName)))
}

// eslint-disable-next-line no-restricted-globals
self.addEventListener('message', event => {
  if (event.data && event.data.type === 'LOGOUT') {
    event.waitUntil(deleteCache())
  }
})
