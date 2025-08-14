package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mcp-wasm-server/wasm"
	"net/http"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Struct to hold property and type
type PropertyType struct {
	Property string
	Type     string
}

func TransformInputSchema(inputSchemaJSON string) (*mcp.ToolInputSchema, error) {
	var inputSchema mcp.ToolInputSchema

	// Unmarshal the JSON string into the ToolInputSchema struct
	err := json.Unmarshal([]byte(inputSchemaJSON), &inputSchema)
	if err != nil {
		return nil, err
	}

	return &inputSchema, nil
}

func main() {

	// Create MCP server
	s := server.NewMCPServer(
		"mcp-wasm-server",
		"0.0.0",
	)

	pluginsPath := os.Getenv("PLUGINS_PATH")
	if pluginsPath == "" {
		pluginsPath = "./plugins"
	}

	wasm.LoadPlugins(pluginsPath, map[string]string{})
	fmt.Println("🔥 Loaded plugins:", wasm.GetToolSet())

	// =================================================
	// TOOLS:
	// =================================================
	for name, tool := range wasm.GetToolSet() {
		// Register the tool with the server
		fmt.Println("🔥 Registering tool:", name)
		fmt.Println("📝 Description.    :", tool.Description)
		fmt.Println("🔧 Arguments       :", tool.InputSchema)

		// make jsonstring from tool.InputSchema
		inputSchemaJSON, err := json.Marshal(tool.InputSchema)
		if err != nil {
			log.Fatalf("Error marshalling input schema for tool %s: %v", name, err)
		}
		fmt.Println("Input Schema JSON:", string(inputSchemaJSON))

		wasmTool := mcp.NewTool(name, mcp.WithDescription(tool.Description))
		//wasmTool.RawInputSchema = string(inputSchemaJSON)
		inputSchema, _ := TransformInputSchema(string(inputSchemaJSON))
		wasmTool.InputSchema = *inputSchema

		s.AddTool(
			wasmTool,
			func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
				args := request.GetArguments()
				fmt.Println("🔥 Tool arguments:", args)

				result, err := tool.Handler(args)
				if err != nil {
					return nil, fmt.Errorf("error calling tool %s: %w", name, err)
				}
				// Return the result as a tool result
				return mcp.NewToolResultText(fmt.Sprintf("Tool %s executed successfully: %v", name, result)), nil
			},
		)
	}

	// Start the HTTP server
	httpPort := os.Getenv("MCP_HTTP_PORT")
	if httpPort == "" {
		httpPort = "9090"
	}

	log.Println("🟪 MCP WASM StreamableHTTP server is running on port", httpPort)

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


func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"status": "healthy",
		"server": "mcp-files-server",
	}
	json.NewEncoder(w).Encode(response)
}