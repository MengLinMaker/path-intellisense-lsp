import { build } from 'esbuild'

build({
  entryPoints: ['src/extension.ts'],
  outfile: 'dist/index.js',
  format: 'cjs',
  platform: 'node',
  bundle: true,
  minify: false,
  sourcemap: true,
  external: ['vscode'],
})
