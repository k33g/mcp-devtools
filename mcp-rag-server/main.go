package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/openai/openai-go/v2" // imported as openai
	"github.com/openai/openai-go/v2/option"

	"mcp-rag-server/helpers"
	"mcp-rag-server/rag"
)

var client openai.Client
var store rag.MemoryVectorStore
var embeddingsModel string

func main() {
	ctx := context.Background()

	// Create MCP server
	s := server.NewMCPServer(
		"mcp-rag-server",
		"0.0.0",
	)

	// Ensure MODEL_RUNNER_BASE_URL is set in the environment
	if os.Getenv("MODEL_RUNNER_BASE_URL") == "" {
		os.Setenv("MODEL_RUNNER_BASE_URL", "http://localhost:12434/engines/llama.cpp/v1/")
	}
	if os.Getenv("EMBEDDING_MODEL") == "" {
		os.Setenv("EMBEDDING_MODEL", "ai/mxbai-embed-large:latest")
	}
	// Ensure JSON_STORE_FILE_PATH is set in the environment
	if os.Getenv("JSON_STORE_FILE_PATH") == "" {
		os.Setenv("JSON_STORE_FILE_PATH", "rag-memory-store.json")
	}
	// Ensure DOCUMENTS_PATH is set in the environment
	if os.Getenv("DOCUMENTS_PATH") == "" {
		os.Setenv("DOCUMENTS_PATH", "markdown")
	}

	llmURL := os.Getenv("MODEL_RUNNER_BASE_URL")
	embeddingsModel = os.Getenv("EMBEDDING_MODEL")
	jsonStoreFilePath := os.Getenv("JSON_STORE_FILE_PATH")
	documentsPath := os.Getenv("DOCUMENTS_PATH")

	client = openai.NewClient(
		option.WithBaseURL(llmURL),
		option.WithAPIKey(""),
	)

	// -------------------------------------------------
	// Create a vector store
	// -------------------------------------------------
	store = rag.MemoryVectorStore{
		Records: make(map[string]rag.VectorRecord),
	}

	// Load the vector store from a file if it exists
	err := store.Load(jsonStoreFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("🚀 No existing vector store found, starting fresh.")

			// =================================================
			// CHUNKS:
			// =================================================
			contents, err := helpers.GetContentFiles(documentsPath, ".md")
			if err != nil {
				log.Fatalln("😡 Error getting content files:", err)
			}
			chunks := []string{}
			fmt.Println("💡 Found", len(contents), "content files to process.")
			//fmt.Println("📂 Processing content files...", contents)
			fmt.Println("📝 Processing(Chunking) content files...")

			chunkSize := os.Getenv("CHUNK_SIZE")
			if chunkSize == "" {
				chunkSize = "1024"
			}
			chunkOverlap := os.Getenv("CHUNK_OVERLAP")
			if chunkOverlap == "" {
				chunkOverlap = "256"
			}
			chunkSizeInt, err := strconv.Atoi(chunkSize)
			if err != nil {
				log.Fatalln("😡 Error converting chunk size to int:", err)
			}
			chunkOverlapInt, err := strconv.Atoi(chunkOverlap)
			if err != nil {
				log.Fatalln("😡 Error converting chunk overlap to int:", err)
			}

			for _, content := range contents {
				chunks = append(chunks, rag.ChunkText(content, chunkSizeInt, chunkOverlapInt)...)
				//chunks = append(chunks, rag.SplitTextWithDelimiter(content, "---")...)
				//chunks = append(chunks, rag.ChunkWithMarkdownHierarchy(content)... )

			}

			// -------------------------------------------------
			// Create and save the embeddings from the chunks
			// -------------------------------------------------
			fmt.Println("⏳ Creating the embeddings...")

			for idx, chunk := range chunks {
				embeddingsResponse, err := client.Embeddings.New(ctx, openai.EmbeddingNewParams{
					Input: openai.EmbeddingNewParamsInputUnion{
						OfString: openai.String(chunk),
					},
					Model: embeddingsModel,
				})

				if err != nil {
					fmt.Println(err)
				} else {
					_, errSave := store.Save(rag.VectorRecord{
						Prompt:    chunk,
						Embedding: embeddingsResponse.Data[0].Embedding,
					})
					if errSave != nil {
						fmt.Println("😡:", errSave)
					}
					fmt.Println("✅ Chunk", idx, "saved with embedding:", len(embeddingsResponse.Data[0].Embedding))
				}
			}

			fmt.Println("✋", "Embeddings created, total of records", len(store.Records))
			err = store.Persist(jsonStoreFilePath)
			if err != nil {
				log.Fatalln("😡 Error saving vector store:", err)
			}
			fmt.Println("✅ Vector store saved to", jsonStoreFilePath)
			fmt.Println("💾 Vector store initialized with", len(store.Records), "records.")
			fmt.Println()

		} else {
			log.Fatalln("Error loading vector store:", err)
		}
	} else {
		log.Println("Vector store loaded successfully, total records:", len(store.Records))
	}

	// =================================================
	// TOOLS:
	// =================================================
	searchInDoc := mcp.NewTool("rag_question",
		mcp.WithDescription(`Find an answer in the internal database.`),
		mcp.WithString("search_question",
			mcp.Required(),
			mcp.Description("Search question"),
		),
	)
	s.AddTool(searchInDoc, searchInDocHandler)

	// Start the HTTP server
	httpPort := os.Getenv("MCP_HTTP_PORT")
	fmt.Println("🌍 MCP HTTP Port:", httpPort)
	if httpPort == "" {
		httpPort = "9090"
	}

	log.Println("MCP StreamableHTTP server is running on port", httpPort)

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

func searchInDocHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	args := request.GetArguments()
	userQuestion := args["search_question"].(string)

	fmt.Println("🔍 Searching for question:", userQuestion)

	// -------------------------------------------------
	// Search for similarities
	// -------------------------------------------------

	fmt.Println("⏳ Searching for similarities...")

	// -------------------------------------------------
	// Create embedding from the user question
	// -------------------------------------------------
	embeddingsResponse, err := client.Embeddings.New(ctx, openai.EmbeddingNewParams{
		Input: openai.EmbeddingNewParamsInputUnion{
			OfString: openai.String(userQuestion),
		},
		Model: embeddingsModel,
	})
	if err != nil {
		log.Fatal("😡:", err)
	}

	// -------------------------------------------------
	// Create a vector record from the user embedding
	// -------------------------------------------------
	embeddingFromUserQuestion := rag.VectorRecord{
		Embedding: embeddingsResponse.Data[0].Embedding,
	}

	strLimit := os.Getenv("LIMIT")
	if strLimit == "" {
		strLimit = "0.6"
	}
	strMax := os.Getenv("MAX_RESULTS")
	if strMax == "" {
		strMax = "2"
	}
	// Convert string to float64 and int
	var limit float64
	fmt.Sscanf(strLimit, "%f", &limit)
	var maxResults int
	fmt.Sscanf(strMax, "%d", &maxResults)

	similarities, _ := store.SearchTopNSimilarities(embeddingFromUserQuestion, limit, maxResults)

	documentsContent := "Documents:\n"

	for _, similarity := range similarities {
		fmt.Println("✅ CosineSimilarity:", similarity.CosineSimilarity, "Chunk:", similarity.Prompt)
		documentsContent += similarity.Prompt
	}
	documentsContent += "\n"
	fmt.Println("✋", "Similarities found, total of records", len(similarities))
	fmt.Println()

	// -------------------------------------------------
	// Generate embeddings from user question
	// -------------------------------------------------
	// EMBEDDINGS...
	return mcp.NewToolResultText(documentsContent), nil
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check if vector store is initialized and has records
	if len(store.Records) == 0 {
		w.WriteHeader(http.StatusServiceUnavailable)
		response := map[string]interface{}{
			"status": "unhealthy",
			"reason": "vector store not initialized",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"status":           "healthy",
		"records":          len(store.Records),
		"embeddings_model": embeddingsModel,
	}
	json.NewEncoder(w).Encode(response)
}
