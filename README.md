# plugin.tree-generator

A [GAMS](https://github.com/kkgams/gams) WASM plugin for procedural tree generation.

| Unit | Source / artifact | Screenshot | License | Original work |
| --- | --- | --- | --- | --- |
| Procedural tree generator | `plugins/treegen.comp/main.go` / `plugin.tree-generator.wasm` | Not applicable | [Apache-2.0](LICENSE), [third-party terms](LICENSING.md) | GAMS |

## API

`plugins/treegen.comp/wit/package.wit` defines `gams:tree-generator/tree-generator@1.0.0`: `gen(config)` accepting node count, maximum depth/branching and root branch count, returning nodes with string-pair properties and parent IDs. The `generator` world uses Host-provided WASI imports.

## Development

Use `nix develop` (or direnv) for pinned Go 1.25, TinyGo, wasm-tools and documentation tools. No sibling checkout is required.

| Command | Purpose |
| --- | --- |
| `make build` | Compile and validate the component with embedded notices. |
| `make package` | Prepare `build.nosync/plugin.tree-generator.wasm`, `LICENSES.zip` and `SHA256SUMS` incrementally. |
| `make test` | Run the existing native tree generator and SDK tilemap Go tests. |
| `make check` | Check Python tooling syntax; native tests are separate from WASM runtime coverage. |
| `make open` | Build and open this local Project in the released macOS Apple Silicon Host. |
| `make docs` | Build the wiki to `build.nosync/site`. |
| `make serve-docs` | Preview the wiki; `DOCS_PORT` defaults to 8080. |
| `make validate-docs` | Validate wiki content. |
| `make clean` | Remove `build.nosync` and identified Host/Unit downloads; refuse tracked files, symlinks and unidentified supplied tools/caches. |

The local Project browses files with pinned published Units; it is not a tree editor or runtime test. Downloads preserve accompanying notices and existing Host tools. Packaging does not verify native Host/API behavior.

Distribute the readable `LICENSES.zip` terms together with the raw WASM.
