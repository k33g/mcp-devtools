#!/bin/bash
: <<'COMMENT'
# Use tool "write_file"
COMMENT

# STEP 1: Load the session ID from the environment file
source mcp.env
source mcp.server.env

MCP_SERVER=${MCP_SERVER:-"http://localhost:${MCP_HTTP_PORT}"}

# Example: Write a file
read -r -d '' DATA <<- EOM
{
  "jsonrpc": "2.0",
  "id": "test",
  "method": "tools/call",
  "params": {
    "name": "write_file",
    "arguments": {
      "file_path": "example.txt",
      "content": "Hello from MCP Files Server!\nThis is a test file created by the write_file tool.\nTimestamp: $(date)"
    }
  }
}
EOM

curl ${MCP_SERVER}/mcp \
  -H "Content-Type: application/json" \
  -H "Mcp-Session-Id: $SESSION_ID" \
  -d "${DATA}" | jq