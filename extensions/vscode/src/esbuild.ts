import { build } from 'esbuild'

build({
  entryPoints: ['src/extension.ts'],
  outfile: 'out/index.cjs',
  format: 'cjs',
  platform: 'node',
  bundle: true,
  minify: false,
  sourcemap: true,
  external: ['vscode'],
})
