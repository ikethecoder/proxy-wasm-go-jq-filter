# Proxy-Wasm Go JQ Filter

A plugin for the [Kong Microservice API Gateway](https://konghq.com/solutions/gateway/) to transform payloads using [JQ](https://jqlang.github.io/jq/). It is a [Proxy-Wasm](https://github.com/proxy-wasm/spec) filter that uses the [jqgo library](https://github.com/itchyny/gojq) for transforming JSON Responses.

## Tested and working for

| Kong Version |   Tests passing    |
| ------------ | :----------------: |
| 3.4.x        | :white_check_mark: |

## Installation

### Using docker

```
docker build --tag kong-wasm-test  .
```

### Running Kong

```
docker run --rm --name kong \
     -e "KONG_LOG_LEVEL=info" \
     -e "KONG_WASM=on" \
     -e "KONG_DATABASE=off" \
     -e "KONG_DECLARATIVE_CONFIG=/kong/config/kong.yml" \
     -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
     -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
     -e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
     -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
     -e "KONG_ADMIN_LISTEN=0.0.0.0:8001, 0.0.0.0:8444 ssl" \
     -e "KONG_UNTRUSTED_LUA_SANDBOX_REQUIRES=cjson,cjson.safe" \
     -e "KONG_UNTRUSTED_LUA_SANDBOX_ENVIRONMENT=table.concat" \
     -v `pwd`/local/kong_config.yml:/kong/config/kong.yml \
     -p 6200:8000 \
     -p 6443:8443 \
     -p 8001:8001 \
     -p 127.0.0.1:8444:8444 \
     kong-wasm-test
```

## Usage

### Filter Structure

```
filter_chains:
  - name: my-filter-chain
    tags: [ns.sample]
    filters:
     - name: jq-filter
       config: >-
            {
            "Query": ".data | fromjson | .[] | select(.Priority == \"Urgent\") | .Id"
            }

```

### Parameters

| Parameter    | Required | Default | Description                                              |
| ------------ | -------- | ------- | -------------------------------------------------------- |
| name         | yes      |         | The name of the filter to use, in this case `jq-filter`. |
| config       | yes      |         | A JSON object represented as a string.                   |
| config.Query | no       | "."     | The JQ query to perform on the response payload.         |

### Example

```
curl -d @local/sample_payload.json \
  -H "Content-Type: application/json" \
  http://127.0.0.1:6200/post
```

Result:

```json
["1", "3"]
```

## Testing

### Prerequisites

- [tinygo](https://tinygo.org) - a Go compiler that can produce WebAssembly code.

### Build

Once the Go environment is set up and tinygo is in the PATH, build the filter running
`make`.

This will produce a .wasm file in the root of the project.

### Running Tests

```sh
tinygo test
```
