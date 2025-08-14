# MCP WASM Server

A Model Context Protocol (MCP) server that dynamically loads and exposes WebAssembly (WASM) plugins as tools. This server allows you to extend its functionality by adding WASM plugins without modifying the core server code.

## Overview

This MCP server provides a plugin architecture where WebAssembly modules can be loaded as tools and exposed through the MCP protocol. Each WASM plugin can define multiple tools with their own schemas and implementations.

## Architecture

- **Main Server** (`main.go`): HTTP server that implements the MCP protocol
- **WASM Runtime** (`wasm/`): Handles loading and executing WASM plugins using Extism
- **Plugin Interface**: WASM plugins implement specific exported functions to register tools
- **Tool Registration**: Dynamic discovery and registration of tools from WASM modules

## Features

- 🔌 **Dynamic Plugin Loading**: Load WASM plugins from a configurable directory
- 🛠️ **Tool Registration**: Automatically discover and register tools from plugins
- 🌐 **HTTP Server**: Streamable HTTP server with MCP endpoint and health checks
- ⚙️ **Environment Configuration**: Configure plugins and server through environment variables
- 🔧 **Extism Integration**: Uses Extism for secure WASM runtime execution

## Quick Start

### Prerequisites

- Go 1.24.0 or later
- WASM plugins in the `./plugins` directory

### Running the Server

1. Clone the repository
2. Build and run:
   ```bash
   go mod tidy
   go run main.go
   ```

3. The server will start on port 9090 (configurable via `MCP_HTTP_PORT`)

### Configuration

Environment variables:
- `MCP_HTTP_PORT`: HTTP server port (default: 9090)
- `PLUGINS_PATH`: Path to WASM plugins directory (default: ./plugins)
- `WASM_*`: Any environment variables starting with `WASM_` are passed to plugins

## Plugin Development

### Plugin Interface

WASM plugins must implement the following exported functions:

#### `tools_information`
Returns a JSON array of available tools:
```json
[
  {
    "name": "tool_name",
    "description": "Tool description",
    "inputSchema": {
      "type": "object",
      "required": ["param1"],
      "properties": {
        "param1": {
          "type": "string",
          "description": "Parameter description"
        }
      }
    }
  }
]
```

#### Tool Functions
Each tool must have a corresponding exported function with the same name that:
- Takes JSON input via `pdk.InputString()`
- Returns results via `pdk.OutputString()`

### Example Plugin

See the included plugins for examples:

**D&D Greetings Plugin** (`plugins/dnd/`):
- `orc_greetings`: Greets as an Orc
- `vulcan_greetings`: Greets as a Vulcan

**Name Generator Plugin** (`plugins/gen-name/`):
- `generate_name`: Transforms names into D&D character names based on race

### Building Plugins

Each plugin directory contains build scripts:
```bash
cd plugins/your-plugin
./build.sh
```

## API Endpoints

- `GET /health`: Health check endpoint
- `POST /mcp`: MCP protocol endpoint for tool execution

## Example Usage

Once the server is running, you can interact with it using any MCP client:

1. **List available tools**: The server automatically exposes all tools from loaded WASM plugins
2. **Execute tools**: Send tool execution requests via the MCP protocol

## Project Structure

```
mcp-wasm-server/
├── main.go              # Main server implementation
├── wasm/
│   ├── wasm.go         # WASM plugin loading logic
│   └── tools.go        # Tool registration and execution
├── tools/
│   └── tools.go        # Tool interface definitions
├── plugins/
│   ├── dnd.wasm        # Compiled D&D greetings plugin
│   ├── gen-name.wasm   # Compiled name generator plugin
│   ├── dnd/            # D&D greetings plugin source
│   └── gen-name/       # Name generator plugin source
└── scripts/            # Utility scripts
```

## Dependencies

- **Extism Go SDK**: WebAssembly runtime for plugins
- **MCP Go**: Model Context Protocol implementation
- **Wazero**: WebAssembly runtime engine

## Security

- WASM plugins run in a sandboxed environment
- Plugins have controlled access to system resources
- Network access can be configured per plugin
