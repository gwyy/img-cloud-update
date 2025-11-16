import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
import * as dotenv from 'dotenv'
import * as fs from 'fs'


// https://vite.dev/config/

export default ({ mode }) => {
  // 读取环境变量
  const NODE_ENV = mode || 'development'
  const envFiles = [`.env.${NODE_ENV}`]
  for (const file of envFiles) {
    const envConfig = dotenv.parse(fs.readFileSync(file))
    for (const k in envConfig) {
      process.env[k] = envConfig[k]
    }
  }

  const alias = {
    '@': path.resolve(__dirname, './src'),
    vue$: 'vue/dist/vue.runtime.esm-bundler.js'
  }


  const config = {
    host: process.env.VITE_HOST,
    base: '/', // 编译后js导入的资源路径
    root: './', // index.html文件所在位置
    publicDir: 'public', // 静态资源文件夹
   
    resolve: {
      alias
    },
    server: {
      // 如果使用docker-compose开发模式，设置为false
      open: true,
      port: process.env.VITE_CLI_PORT,
      proxy: {
        // 把key的路径代理到target位置
        // detail: https://cli.vuejs.org/config/#devserver-proxy
        [process.env.VITE_BASE_API]: {
          // 需要代理的路径   例如 '/api'
          target: `${process.env.VITE_BASE_PATH}:${process.env.VITE_SERVER_PORT}/`, // 代理到 目标路径
          changeOrigin: true,
          // rewrite: (path) =>  path.replace(new RegExp('^' + process.env.VITE_BASE_API), '')
        }
      }
    },
    plugins: [vue()],
  }

  return config
}
