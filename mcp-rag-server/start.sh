#!/bin/bash
#export MODEL_RUNNER_BASE_URL=http://model-runner.docker.internal/engines/llama.cpp/v1
export MODEL_RUNNER_BASE_URL=http://localhost:12434/engines/llama.cpp/v1
export EMBEDDING_MODEL=ai/granite-embedding-multilingual:latest
export MCP_HTTP_PORT=9095
export LIMIT=0.45
export MAX_RESULTS=30
export JSON_STORE_FILE_PATH=store/rag-memory-store.json
export CHUNK_SIZE=1024
export CHUNK_OVERLAP=512
go run main.go