import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tsconfigPaths from 'vite-tsconfig-paths'
import svgr from 'vite-plugin-svgr'
import EnvironmentPlugin from 'vite-plugin-environment'

export default defineConfig({
  plugins: [react(), tsconfigPaths(), svgr(), EnvironmentPlugin('all')],
  server: {
    port: 4242,
    open: true
  }
})
