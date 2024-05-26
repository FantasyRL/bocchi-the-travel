import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import VueDevTools from 'vite-plugin-vue-devtools'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    proxy: {
      '/bocchi': {
        target: 'http://127.0.0.1:10001',
        ws: true,// 是否代理websocket
        changeOrigin: true, // 是否将主机头的源更改为目标URL,允许跨域
        pathRewrite: {
          '^/bocchi': '/' // 路径重写，将'/api'替换为空串
        }
      }
    }
  },
  plugins: [
    vue({
      template: {
        compilerOptions: {
          // 所有以 mdui- 开头的标签名都是 mdui 组件
          isCustomElement: (tag) => tag.startsWith('mdui-')
        }
      }
    }),
    VueDevTools(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
