package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {

	// Create MCP server
	s := server.NewMCPServer(
		"mcp-files-server",
		"0.0.0",
	)

	// Read file tool
	readFileTool := mcp.NewTool("read_file",
		mcp.WithDescription("Read the content of a text file"),
		mcp.WithString("file_path",
			mcp.Required(),
			mcp.Description("Path to the file to read"),
		),
	)
	s.AddTool(readFileTool, readFileHandler)

	// Write file tool
	writeFileTool := mcp.NewTool("write_file",
		mcp.WithDescription("Write content to a text file"),
		mcp.WithString("file_path",
			mcp.Required(),
			mcp.Description("Path to the file to write"),
		),
		mcp.WithString("content",
			mcp.Required(),
			mcp.Description("Content to write to the file"),
		),
	)
	s.AddTool(writeFileTool, writeFileHandler)

	// Delete file tool
	deleteFileTool := mcp.NewTool("delete_file",
		mcp.WithDescription("Delete a file from the filesystem"),
		mcp.WithString("file_path",
			mcp.Required(),
			mcp.Description("Path to the file to delete"),
		),
	)
	s.AddTool(deleteFileTool, deleteFileHandler)

	// Start the HTTP server
	httpPort := os.Getenv("MCP_HTTP_PORT")
	if httpPort == "" {
		httpPort = "9090"
	}

	log.Println("MCP Files Server is running on port", httpPort)

	// Create a custom mux to handle both MCP and health endpoints
	mux := http.NewServeMux()

	// Add healthcheck endpoint
	mux.HandleFunc("/health", healthCheckHandler)

	// Add MCP endpoint
	httpServer := server.NewStreamableHTTPServer(s,
		server.WithEndpointPath("/mcp"),
	)

	// Register MCP handler with the mux
	mux.Handle("/mcp", httpServer)

	// Start the HTTP server with custom mux
	log.Fatal(http.ListenAndServe(":"+httpPort, mux))
}

func readFileHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	filePathArg, exists := args["file_path"]
	if !exists || filePathArg == nil {
		return nil, fmt.Errorf("missing required parameter 'file_path'")
	}

	filePath, ok := filePathArg.(string)
	if !ok {
		return nil, fmt.Errorf("parameter 'file_path' must be a string")
	}

	// Clean and validate the file path
	cleanPath := filepath.Clean(filePath)

	// Prefix with LOCAL_WORKSPACE_FOLDER if set
	if workspaceFolder := os.Getenv("LOCAL_WORKSPACE_FOLDER"); workspaceFolder != "" {
		cleanPath = filepath.Join(workspaceFolder, cleanPath)
	}

	// Read the file
	content, err := os.ReadFile(cleanPath)
	if err != nil {
		if os.IsNotExist(err) {
			return mcp.NewToolResultError(fmt.Sprintf("File not found: %s", cleanPath)), nil
		}
		return mcp.NewToolResultError(fmt.Sprintf("Error reading file: %v", err)), nil
	}

	log.Printf("Successfully read file: %s (%d bytes)", cleanPath, len(content))
	return mcp.NewToolResultText(string(content)), nil
}

func writeFileHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	filePathArg, exists := args["file_path"]
	if !exists || filePathArg == nil {
		return nil, fmt.Errorf("missing required parameter 'file_path'")
	}

	filePath, ok := filePathArg.(string)
	if !ok {
		return nil, fmt.Errorf("parameter 'file_path' must be a string")
	}

	contentArg, exists := args["content"]
	if !exists || contentArg == nil {
		return nil, fmt.Errorf("missing required parameter 'content'")
	}

	content, ok := contentArg.(string)
	if !ok {
		return nil, fmt.Errorf("parameter 'content' must be a string")
	}

	// Clean the file path
	cleanPath := filepath.Clean(filePath)

	// Prefix with LOCAL_WORKSPACE_FOLDER if set
	if workspaceFolder := os.Getenv("LOCAL_WORKSPACE_FOLDER"); workspaceFolder != "" {
		cleanPath = filepath.Join(workspaceFolder, cleanPath)
	}

	// Create directory if it doesn't exist
	dir := filepath.Dir(cleanPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error creating directory: %v", err)), nil
	}

	// Write the file
	err := os.WriteFile(cleanPath, []byte(content), 0644)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error writing file: %v", err)), nil
	}

	log.Printf("Successfully wrote file: %s (%d bytes)", cleanPath, len(content))
	return mcp.NewToolResultText(fmt.Sprintf("Successfully wrote %d bytes to %s", len(content), cleanPath)), nil
}

func deleteFileHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	filePathArg, exists := args["file_path"]
	if !exists || filePathArg == nil {
		return nil, fmt.Errorf("missing required parameter 'file_path'")
	}

	filePath, ok := filePathArg.(string)
	if !ok {
		return nil, fmt.Errorf("parameter 'file_path' must be a string")
	}

	// Clean the file path
	cleanPath := filepath.Clean(filePath)

	// Prefix with LOCAL_WORKSPACE_FOLDER if set
	if workspaceFolder := os.Getenv("LOCAL_WORKSPACE_FOLDER"); workspaceFolder != "" {
		cleanPath = filepath.Join(workspaceFolder, cleanPath)
	}

	// Check if file exists
	if _, err := os.Stat(cleanPath); os.IsNotExist(err) {
		return mcp.NewToolResultError(fmt.Sprintf("File not found: %s", cleanPath)), nil
	}

	// Delete the file
	err := os.Remove(cleanPath)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error deleting file: %v", err)), nil
	}

	log.Printf("Successfully deleted file: %s", cleanPath)
	return mcp.NewToolResultText(fmt.Sprintf("Successfully deleted file: %s", cleanPath)), nil
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"status": "healthy",
		"server": "mcp-files-server",
	}
	json.NewEncoder(w).Encode(response)
}
