#!/bin/bash
: <<'COMMENT'
# Use tool "read_file" or "write_file"
COMMENT

# STEP 1: Load the session ID from the environment file
source mcp.env
source mcp.server.env

MCP_SERVER=${MCP_SERVER:-"http://localhost:${MCP_HTTP_PORT}"}

# Example: Read a file
read -r -d '' DATA <<- EOM
{
  "jsonrpc": "2.0",
  "id": "test",
  "method": "tools/call",
  "params": {
    "name": "read_file",
    "arguments": {
      "file_path": "test.txt"
    }
  }
}
EOM

# Uncomment this example to write a file instead:
# read -r -d '' DATA <<- EOM
# {
#   "jsonrpc": "2.0",
#   "id": "test",
#   "method": "tools/call",
#   "params": {
#     "name": "write_file",
#     "arguments": {
#       "file_path": "example.txt",
#       "content": "Hello from MCP Files Server!"
#     }
#   }
# }
# EOM

curl ${MCP_SERVER}/mcp \
  -H "Content-Type: application/json" \
  -H "Mcp-Session-Id: $SESSION_ID" \
  -d "${DATA}" | jq