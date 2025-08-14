#!/bin/bash

extism call ../gen-name.wasm tools_information \
  --log-level "info" \
  --wasi
echo ""

extism call ../gen-name.wasm generate_name \
  --input '{"name":"Philippe","race":"elf"}' \
  --log-level "info" \
  --wasi
echo ""
