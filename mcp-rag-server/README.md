# MCP RAG Server

A Model Context Protocol (MCP) server that implements Retrieval-Augmented Generation (RAG) functionality. This server processes markdown documents, creates vector embeddings, and provides semantic search capabilities through MCP tools.

## Features

- **Document Processing**: Automatically processes markdown files from a specified directory
- **Text Chunking**: Splits documents into configurable chunks with overlap for better context retention
- **Vector Embeddings**: Creates embeddings using configurable embedding models
- **Persistent Storage**: Saves vector store to JSON for persistence across restarts
- **Semantic Search**: Provides `rag_question` tool for finding relevant information in your document collection
- **MCP Integration**: Exposes functionality through the Model Context Protocol

## Configuration

The server uses environment variables for configuration with sensible defaults:

| Variable | Default Value | Description |
|----------|---------------|-------------|
| `MODEL_RUNNER_BASE_URL` | `http://localhost:12434/engines/llama.cpp/v1/` | Base URL for the LLM/embedding service |
| `EMBEDDING_MODEL` | `ai/mxbai-embed-large:latest` | Model name for generating embeddings |
| `JSON_STORE_FILE_PATH` | `rag-memory-store.json` | Path to persist the vector store |
| `DOCUMENTS_PATH` | `markdown` | Directory containing markdown files to process |
| `CHUNK_SIZE` | `1024` | Size of text chunks in characters |
| `CHUNK_OVERLAP` | `256` | Overlap between chunks in characters |
| `MCP_HTTP_PORT` | `9090` | Port for the MCP HTTP server |
| `LIMIT` | `0.6` | Minimum similarity threshold for search results |
| `MAX_RESULTS` | `2` | Maximum number of search results to return |

## How It Works

1. **Initialization**: On first run, the server processes all markdown files in the specified directory
2. **Chunking**: Documents are split into overlapping chunks for better semantic retrieval
3. **Embedding Creation**: Each chunk is converted to a vector embedding using the specified model
4. **Storage**: The vector store is persisted to disk for future use
5. **Search**: The `rag_question` tool finds the most semantically similar chunks to answer questions

## MCP Tools

### `rag_question`
Searches the document collection for information relevant to a given question.

**Parameters:**
- `question` (required): The search question to find relevant information

**Returns:** The most relevant document chunks that can help answer the question.

## Usage

1. Place your markdown documents in the configured directory (default: `markdown/`)
2. Start the server: `go run main.go`
3. Connect your MCP client to `http://localhost:9090/mcp`
4. Use the `rag_question` tool to search your document collection

## Requirements

- Go 1.19+
- Access to an embedding service (e.g., Ollama with mxbai-embed-large model)
- Markdown documents to index