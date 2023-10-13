FROM tinygo/tinygo:0.30.0 as WASM
COPY --chown=tinygo . /home/tinygo
RUN make

FROM kong:3.4.2
ARG KONG_WASM_FILTERS_PATH=/usr/local/kong/wasm_filters
ENV KONG_WASM_FILTERS_PATH=${KONG_WASM_FILTERS_PATH}
COPY --from=WASM /home/tinygo/jq-filter.wasm ${KONG_WASM_FILTERS_PATH}/jq-filter.wasm

USER kong