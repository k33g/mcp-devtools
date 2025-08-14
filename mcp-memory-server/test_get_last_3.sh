#!/bin/bash

# Load the session ID from the environment file
source mcp.env
source mcp.server.env

MCP_SERVER=${MCP_SERVER:-"http://localhost:${MCP_HTTP_PORT}"}

curl -X POST "${MCP_SERVER}/mcp" \
  -H "Content-Type: application/json" \
  -H "Mcp-Session-Id: $SESSION_ID" \
  -d '{
    "jsonrpc": "2.0",
    "id": "get-3-test",
    "method": "tools/call",
    "params": {
      "name": "get_last_3_messages"
    }
  }' | jq