# Proxy-Wasm Go JQ Filter

This repository contains a [Proxy-Wasm](https://github.com/proxy-wasm/spec)
filter using the JQ JSON processor for transforming JSON Responses.

## Requirements

- [tinygo](https://tinygo.org) - a Go compiler that can produce WebAssembly code.

## Build

Once the Go environment is set up and tinygo is in the PATH, build the filter running
`make`.

This will produce a .wasm file in the root of the project.
