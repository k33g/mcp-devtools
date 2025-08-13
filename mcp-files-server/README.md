# MCP Files Server

A lightweight HTTP streamable MCP (Model Context Protocol) server implemented in Go that provides secure file operations within a designated workspace.

## Goal

The MCP Files Server enables AI assistants and automation tools to safely read and write text files through the standardized MCP protocol. This server provides:

- **Workspace Isolation**: All file operations are contained within a configurable workspace directory via `LOCAL_WORKSPACE_FOLDER`
- **Secure File Access**: Proper path validation and cleaning to prevent directory traversal attacks
- **Simple API**: Two essential tools for file manipulation - read and write operations
- **HTTP Streaming**: Built on the MCP streamable HTTP protocol for real-time communication
- **Zero Dependencies**: Minimal external dependencies for easy deployment and maintenance

## Use Cases

- **Content Management**: Allow AI assistants to read configuration files, documentation, or data files
- **Code Generation**: Enable automated code writing and modification within project boundaries
- **Data Processing**: Read input files and write processed results
- **Template Systems**: Read templates and write generated content
- **Log Analysis**: Read log files and write analysis reports
- **Backup Operations**: Read source files and write backup copies

## Tools

### `read_file`
Reads the content of a text file within the workspace.

**Parameters:**
- `file_path` (string, required): Relative path to the file within the workspace

**Returns:** The complete file content as text

### `write_file`
Writes content to a text file within the workspace.

**Parameters:**
- `file_path` (string, required): Relative path to the file within the workspace
- `content` (string, required): Text content to write to the file

**Returns:** Success message with file path and byte count

## Architecture

The server follows the MCP protocol specification and provides:

1. **Session Management**: Each client connection gets a unique session ID
2. **Tool Registration**: Both file tools are registered with proper parameter validation
3. **Error Handling**: Comprehensive error responses for missing files, permissions, etc.
4. **Health Monitoring**: `/health` endpoint for service monitoring
5. **Workspace Security**: All file paths are prefixed with `LOCAL_WORKSPACE_FOLDER`

## Configuration

Configure the server using environment variables in `mcp.server.env`:

```bash
MCP_HTTP_PORT=9096                          # HTTP port for the server
LOCAL_WORKSPACE_FOLDER=/path/to/workspace   # Base directory for all file operations
```

## Quick Start

1. **Start the server:**
   ```bash
   go run main.go
   ```

2. **Initialize a session:**
   ```bash
   ./initialize.sh
   ```

3. **Test file operations:**
   ```bash
   ./tool.write.call.sh  # Write a test file
   ./tool.call.sh        # Read the test file
   ```

4. **List available tools:**
   ```bash
   ./tools.list.sh
   ```

## Security Features

- **Path Sanitization**: All file paths are cleaned using `filepath.Clean()`
- **Workspace Confinement**: Files can only be accessed within `LOCAL_WORKSPACE_FOLDER`
- **Input Validation**: Strict parameter validation for all tool calls
- **Error Boundaries**: Graceful error handling without exposing system details

## Integration

This MCP server can be integrated with:
- AI development environments
- Code generation pipelines
- Content management systems
- Automated testing frameworks
- Documentation generators
- Data processing workflows

The standardized MCP protocol ensures compatibility with any MCP-compliant client or AI assistant.