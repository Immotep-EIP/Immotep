const CACHE_NAME = 'immotep-cache-v1'
const ASSETS_TO_CACHE = ['/', '/index.html']

const putInCache = async (request, response) => {
  if (request.method !== 'GET') return
  const cache = await caches.open(CACHE_NAME)
  await cache.put(request, response)
}

const cacheFirst = async ({ request, preloadResponsePromise, fallbackUrl }) => {
  const responseFromCache = await caches.match(request)
  if (responseFromCache) {
    return responseFromCache
  }

  const preloadResponse = await preloadResponsePromise
  if (preloadResponse) {
    putInCache(request, preloadResponse.clone())
    return preloadResponse
  }

  try {
    const responseFromNetwork = await fetch(request.clone())
    putInCache(request, responseFromNetwork.clone())
    return responseFromNetwork
  } catch (error) {
    const fallbackResponse = await caches.match(fallbackUrl)
    if (fallbackResponse) {
      return fallbackResponse
    }
    return new Response('Network error happened', {
      status: 408,
      headers: { 'Content-Type': 'text/plain' }
    })
  }
}

const enableNavigationPreload = async () => {
  // eslint-disable-next-line no-restricted-globals
  if (self.registration.navigationPreload) {
    // eslint-disable-next-line no-restricted-globals
    await self.registration.navigationPreload.enable()
  }
}

// eslint-disable-next-line no-restricted-globals
self.addEventListener('install', event => {
  event.waitUntil(
    caches.open(CACHE_NAME).then(cache => cache.addAll(ASSETS_TO_CACHE))
  )
})

// eslint-disable-next-line no-restricted-globals
self.addEventListener('fetch', event => {
  event.respondWith(
    cacheFirst({
      request: event.request,
      preloadResponsePromise: event.preloadResponse,
      fallbackUrl: {
        status: 'offline',
        message:
          'Vous êtes hors ligne. Les données réelles seront affichées dès que vous serez reconnecté.',
        properties: []
      }
    })
  )
})

// eslint-disable-next-line no-restricted-globals
self.addEventListener('activate', event => {
  event.waitUntil(
    Promise.all([
      enableNavigationPreload(),
      caches.keys().then(cacheNames =>
        Promise.all(
          cacheNames.map(cacheName => {
            if (cacheName !== CACHE_NAME) {
              return caches.delete(cacheName)
            }
            return null
          })
        )
      )
    ])
  )
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
