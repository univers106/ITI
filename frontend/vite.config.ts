import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import ui from '@nuxt/ui/vite'


// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
    ui({
      ui: {
        colors: {
          primary: 'amber',
          neutral: 'olive'
        },
        header: {
          slots: {
            toggle: 'sm:hidden',
          },
        },
        icons: {
          arrowDown: 'i-tabler-arrow-down',
          arrowLeft: 'i-tabler-arrow-left',
          arrowRight: 'i-tabler-arrow-right',
          arrowUp: 'i-tabler-arrow-up',
          caution: 'i-tabler-alert-square-rounded',
          check: 'i-tabler-check',
          chevronDoubleLeft: 'i-tabler-chevrons-left',
          chevronDoubleRight: 'i-tabler-chevrons-right',
          chevronDown: 'i-tabler-chevron-down',
          chevronLeft: 'i-tabler-chevron-left',
          chevronRight: 'i-tabler-chevron-right',
          chevronUp: 'i-tabler-chevron-up',
          close: 'i-tabler-x',
          copy: 'i-tabler-copy',
          copyCheck: 'i-tabler-copy-check',
          dark: 'i-tabler-moon',
          drag: 'i-tabler-grip-vertical',
          ellipsis: 'i-tabler-dots',
          error: 'i-tabler-square-rounded-x',
          external: 'i-tabler-external-link',
          eye: 'i-tabler-eye',
          eyeOff: 'i-tabler-eye-off',
          file: 'i-tabler-file',
          folder: 'i-tabler-folder',
          folderOpen: 'i-tabler-folder-open',
          hash: 'i-tabler-hash',
          info: 'i-tabler-info-square-rounded',
          light: 'i-tabler-sun',
          loading: 'i-tabler-loader-2',
          menu: 'i-tabler-menu',
          minus: 'i-tabler-minus',
          panelClose: 'i-tabler-layout-sidebar-left-collapse',
          panelOpen: 'i-tabler-layout-sidebar-left-expand',
          plus: 'i-tabler-plus',
          reload: 'i-tabler-reload',
          search: 'i-tabler-search',
          stop: 'i-tabler-player-stop',
          star: 'i-tabler-star',
          success: 'i-tabler-square-rounded-check',
          system: 'i-tabler-device-desktop',
          tip: 'i-tabler-bulb',
          upload: 'i-tabler-upload',
          warning: 'i-tabler-alert-triangle'
        }
      }
    })
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  server: {
      proxy: {
        '/api': {
          target: 'http://localhost:8080',
          changeOrigin: true,
          secure: false,
        }
      }
    }
})
