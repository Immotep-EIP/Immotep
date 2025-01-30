import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tsconfigPaths from 'vite-tsconfig-paths'
import svgr from 'vite-plugin-svgr'
import { VitePWA } from 'vite-plugin-pwa'

export default defineConfig({
  plugins: [
    react(),
    tsconfigPaths(),
    svgr(),
    VitePWA({
      registerType: 'autoUpdate',
      devOptions: {
        enabled: true, // Activer le mode de développement
        type: 'module' // Utiliser les modules ES pour le SW
      },
      injectRegister: 'auto',
      includeAssets: ['**/*.{js,css,html,png,jpg,svg}'],
      manifest: {
        name: 'Immotep',
        short_name: 'Immotep',
        description: 'Application Immotep',
        theme_color: '#ffffff',
        icons: [
          {
            src: '/assets/ImmotepLogo-DTTSIRh9.svg',
            sizes: '192x192',
            type: 'image/svg+xml'
          }
        ]
      }
    })
  ],
  server: {
    https: {},
    port: 4242,
    open: true
  }
})
