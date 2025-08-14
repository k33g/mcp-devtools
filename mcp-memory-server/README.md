# MCP Memory Server

A streamable HTTP MCP (Model Context Protocol) server in Go that stores and manages messages using the [Artemia](https://github.com/sea-monkeys/artemia) persistence library.

## Features

The server provides the following tools for message management:

- **save_message** - Save a message with content, role, and agent
- **get_last_message** - Retrieve the most recent message
- **get_last_3_messages** - Get the last 3 messages
- **get_last_n_messages** - Get the last N messages (specify N)
- **delete_older_than_hours** - Delete messages older than N hours
- **delete_older_than_days** - Delete messages older than N days
- **delete_all_messages** - Delete all messages from memory

## Message Structure

Each message contains:
- `id`: Unique integer identifier
- `date`: Creation timestamp (ISO 8601 format)
- `content`: The message content
- `role`: Who created the message (assistant, user, system)
- `agent`: Name of the agent

## Setup

1. **Build the server:**
   ```bash
   go build -o mcp-memory-server main.go
   ```

2. **Start the server:**
   ```bash
   ./mcp-memory-server
   ```
   The server runs on port 9091 by default (configurable via `MCP_HTTP_PORT` environment variable).

## Configuration

- `mcp.server.env`: Contains server configuration:
  - `MCP_HTTP_PORT=9091` - Server port
  - `MEMORY_FOLDER=./data` - Directory for storing messages.gob file
- `mcp.env`: Contains the session ID for MCP communication

The `MEMORY_FOLDER` environment variable determines where the `messages.gob` persistence file is stored:
- If not set: stores in current working directory
- If set: creates the directory if it doesn't exist and stores the file there

## Testing

Use the provided test scripts:

- `./test_tools_list.sh` - List available tools
- `./test_save_message.sh` - Save a test message
- `./test_get_last.sh` - Get the last message
- `./test_get_last_3.sh` - Get the last 3 messages
- `./test_get_last_n.sh` - Get the last N messages
- `./test_delete_hours.sh` - Delete messages older than specified hours

## Health Check

The server provides a health check endpoint at `/health`:

```bash
curl http://localhost:9091/health
```

## Data Persistence

Messages are persisted using the Artemia library, which provides:
- In-memory storage with disk persistence
- Automatic data structure serialization
- Index support for efficient querying

Data is stored in `messages.gob` file in the working directory.

## MCP Protocol

The server implements the MCP (Model Context Protocol) over HTTP with:
- JSON-RPC 2.0 messaging
- Session-based communication using `Mcp-Session-Id` header
- Streamable HTTP transport

## Example Usage

```bash
# Save a message
curl -X POST "http://localhost:9091/mcp" \
  -H "Content-Type: application/json" \
  -H "Mcp-Session-Id: your-session-id" \
  -d '{
    "jsonrpc": "2.0",
    "id": "test",
    "method": "tools/call",
    "params": {
      "name": "save_message",
      "arguments": {
        "content": "Hello World!",
        "role": "user",
        "agent": "my-agent"
      }
    }
  }'
```