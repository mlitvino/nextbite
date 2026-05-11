import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'node:path'
import { fileURLToPath } from 'node:url'

// https://vite.dev/config/
const rootDir = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..')

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, rootDir, '')
  const port = Number(env.FRONTEND_PORT) || 5173

  return {
    envDir: rootDir,
    plugins: [react()],
    server: {
      port,
    },
  }
})
