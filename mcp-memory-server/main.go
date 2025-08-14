package main

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/sea-monkeys/artemia"
)

type Message struct {
	ID       int       `json:"id"`
	Date     time.Time `json:"date"`
	Content  string    `json:"content"`
	Role     string    `json:"role"`
	Agent    string    `json:"agent"`
}

var prevalenceLayer *artemia.PrevalenceLayer
var messageIDCounter int
var messageKeys []string
var messageKeysMutex sync.RWMutex

func init() {
	gob.Register(Message{})
}

func main() {
	err := godotenv.Load("mcp.server.env")
	if err != nil {
	}

	// Determine storage path using MEMORY_FOLDER environment variable
	memoryFolder := os.Getenv("MEMORY_FOLDER")
	storagePath := "messages.gob"
	if memoryFolder != "" {
		// Create directory if it doesn't exist
		if err := os.MkdirAll(memoryFolder, 0755); err != nil {
			log.Fatalf("Failed to create memory folder: %v", err)
		}
		storagePath = filepath.Join(memoryFolder, "messages.gob")
	}

	prevalenceLayer, err = artemia.NewPrevalenceLayer(storagePath)
	if err != nil {
		log.Fatalf("Failed to create prevalence layer: %v", err)
	}

	prevalenceLayer.CreateIndex(reflect.TypeOf(Message{}), "Date")
	prevalenceLayer.CreateIndex(reflect.TypeOf(Message{}), "Role")
	prevalenceLayer.CreateIndex(reflect.TypeOf(Message{}), "Agent")

	messageIDCounter = getNextMessageID()

	s := server.NewMCPServer(
		"mcp-memory-server",
		"0.0.1",
	)

	saveMessageTool := mcp.NewTool("save_message",
		mcp.WithDescription("Save a message to memory"),
		mcp.WithString("content",
			mcp.Required(),
			mcp.Description("Content of the message"),
		),
		mcp.WithString("role",
			mcp.Description("Role of the message creator (assistant, user, system). Defaults to 'assistant' if not provided"),
		),
		mcp.WithString("agent",
			mcp.Description("Name of the agent. Defaults to 'unknown' if not provided"),
		),
	)
	s.AddTool(saveMessageTool, saveMessageHandler)

	getLastMessageTool := mcp.NewTool("get_last_message",
		mcp.WithDescription("Get the last message"),
	)
	s.AddTool(getLastMessageTool, getLastMessageHandler)

	getLast3MessagesTool := mcp.NewTool("get_last_3_messages",
		mcp.WithDescription("Get the last 3 messages"),
	)
	s.AddTool(getLast3MessagesTool, getLast3MessagesHandler)

	getLastNMessagesTool := mcp.NewTool("get_last_n_messages",
		mcp.WithDescription("Get the last N messages"),
		mcp.WithString("n",
			mcp.Required(),
			mcp.Description("Number of messages to retrieve"),
		),
	)
	s.AddTool(getLastNMessagesTool, getLastNMessagesHandler)

	deleteOlderThanHoursTool := mcp.NewTool("delete_older_than_hours",
		mcp.WithDescription("Delete messages older than N hours"),
		mcp.WithString("hours",
			mcp.Required(),
			mcp.Description("Number of hours"),
		),
	)
	s.AddTool(deleteOlderThanHoursTool, deleteOlderThanHoursHandler)

	deleteOlderThanDaysTool := mcp.NewTool("delete_older_than_days",
		mcp.WithDescription("Delete messages older than N days"),
		mcp.WithString("days",
			mcp.Required(),
			mcp.Description("Number of days"),
		),
	)
	s.AddTool(deleteOlderThanDaysTool, deleteOlderThanDaysHandler)

	deleteAllMessagesTool := mcp.NewTool("delete_all_messages",
		mcp.WithDescription("Delete all messages from memory"),
	)
	s.AddTool(deleteAllMessagesTool, deleteAllMessagesHandler)

	searchMessagesTool := mcp.NewTool("search_messages",
		mcp.WithDescription("Search messages by keywords in content"),
		mcp.WithString("keywords",
			mcp.Required(),
			mcp.Description("Keywords to search for in message content"),
		),
	)
	s.AddTool(searchMessagesTool, searchMessagesHandler)

	httpPort := os.Getenv("MCP_HTTP_PORT")
	if httpPort == "" {
		httpPort = "9091"
	}

	log.Println("MCP Memory Server is running on port", httpPort)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthCheckHandler)

	httpServer := server.NewStreamableHTTPServer(s,
		server.WithEndpointPath("/mcp"),
	)

	mux.Handle("/mcp", httpServer)

	log.Fatal(http.ListenAndServe(":"+httpPort, mux))
}

func getNextMessageID() int {
	maxID := 0
	messageKeysMutex.RLock()
	for _, key := range messageKeys {
		value, exists := prevalenceLayer.Get(key)
		if exists {
			if msg, ok := value.(Message); ok {
				if msg.ID > maxID {
					maxID = msg.ID
				}
			}
		}
	}
	messageKeysMutex.RUnlock()
	return maxID + 1
}

func saveMessageHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	contentArg, exists := args["content"]
	if !exists || contentArg == nil {
		return nil, fmt.Errorf("missing required parameter 'content'")
	}
	content, ok := contentArg.(string)
	if !ok {
		return nil, fmt.Errorf("parameter 'content' must be a string")
	}

	role := "assistant"
	if roleArg, exists := args["role"]; exists && roleArg != nil {
		if r, ok := roleArg.(string); ok {
			role = r
		}
	}

	agent := "unknown"
	if agentArg, exists := args["agent"]; exists && agentArg != nil {
		if a, ok := agentArg.(string); ok {
			agent = a
		}
	}

	message := Message{
		ID:      messageIDCounter,
		Date:    time.Now(),
		Content: content,
		Role:    role,
		Agent:   agent,
	}

	key := fmt.Sprintf("message_%d", messageIDCounter)
	prevalenceLayer.Set(key, message)
	
	messageKeysMutex.Lock()
	messageKeys = append(messageKeys, key)
	messageKeysMutex.Unlock()

	messageIDCounter++

	log.Printf("Saved message ID: %d, Role: %s, Agent: %s, Content: %s", message.ID, message.Role, message.Agent, message.Content)
	return mcp.NewToolResultText(fmt.Sprintf("Message saved with ID: %d", message.ID)), nil
}

func getLastMessageHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	messages := getAllMessagesSorted()
	if len(messages) == 0 {
		return mcp.NewToolResultText("No messages found"), nil
	}

	lastMessage := messages[len(messages)-1]
	jsonData, err := json.Marshal(lastMessage)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error marshaling message: %v", err)), nil
	}

	return mcp.NewToolResultText(string(jsonData)), nil
}

func getLast3MessagesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	messages := getAllMessagesSorted()
	if len(messages) == 0 {
		return mcp.NewToolResultText("No messages found"), nil
	}

	start := len(messages) - 3
	if start < 0 {
		start = 0
	}
	lastMessages := messages[start:]

	jsonData, err := json.Marshal(lastMessages)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error marshaling messages: %v", err)), nil
	}

	return mcp.NewToolResultText(string(jsonData)), nil
}

func getLastNMessagesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	nArg, exists := args["n"]
	if !exists || nArg == nil {
		return nil, fmt.Errorf("missing required parameter 'n'")
	}

	var n int
	switch v := nArg.(type) {
	case int:
		n = v
	case float64:
		n = int(v)
	case string:
		var err error
		n, err = strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("parameter 'n' must be a valid integer")
		}
	default:
		return nil, fmt.Errorf("parameter 'n' must be an integer")
	}

	messages := getAllMessagesSorted()
	if len(messages) == 0 {
		return mcp.NewToolResultText("No messages found"), nil
	}

	start := len(messages) - n
	if start < 0 {
		start = 0
	}
	lastMessages := messages[start:]

	jsonData, err := json.Marshal(lastMessages)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error marshaling messages: %v", err)), nil
	}

	return mcp.NewToolResultText(string(jsonData)), nil
}

func deleteOlderThanHoursHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	hoursArg, exists := args["hours"]
	if !exists || hoursArg == nil {
		return nil, fmt.Errorf("missing required parameter 'hours'")
	}

	var hours int
	switch v := hoursArg.(type) {
	case int:
		hours = v
	case float64:
		hours = int(v)
	case string:
		var err error
		hours, err = strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("parameter 'hours' must be a valid integer")
		}
	default:
		return nil, fmt.Errorf("parameter 'hours' must be an integer")
	}

	cutoffTime := time.Now().Add(-time.Duration(hours) * time.Hour)
	deletedCount := deleteOlderThan(cutoffTime)

	log.Printf("Deleted %d messages older than %d hours", deletedCount, hours)
	return mcp.NewToolResultText(fmt.Sprintf("Deleted %d messages older than %d hours", deletedCount, hours)), nil
}

func deleteOlderThanDaysHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	daysArg, exists := args["days"]
	if !exists || daysArg == nil {
		return nil, fmt.Errorf("missing required parameter 'days'")
	}

	var days int
	switch v := daysArg.(type) {
	case int:
		days = v
	case float64:
		days = int(v)
	case string:
		var err error
		days, err = strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("parameter 'days' must be a valid integer")
		}
	default:
		return nil, fmt.Errorf("parameter 'days' must be an integer")
	}

	cutoffTime := time.Now().Add(-time.Duration(days) * 24 * time.Hour)
	deletedCount := deleteOlderThan(cutoffTime)

	log.Printf("Deleted %d messages older than %d days", deletedCount, days)
	return mcp.NewToolResultText(fmt.Sprintf("Deleted %d messages older than %d days", deletedCount, days)), nil
}

func getAllMessagesSorted() []Message {
	var messages []Message
	messageKeysMutex.RLock()
	for _, key := range messageKeys {
		value, exists := prevalenceLayer.Get(key)
		if exists {
			if msg, ok := value.(Message); ok {
				messages = append(messages, msg)
			}
		}
	}
	messageKeysMutex.RUnlock()

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].Date.Before(messages[j].Date)
	})

	return messages
}

func deleteOlderThan(cutoffTime time.Time) int {
	deletedCount := 0
	messageKeysMutex.Lock()
	defer messageKeysMutex.Unlock()
	
	newKeys := make([]string, 0)
	
	for _, key := range messageKeys {
		value, exists := prevalenceLayer.Get(key)
		if exists {
			if msg, ok := value.(Message); ok {
				if msg.Date.Before(cutoffTime) {
					prevalenceLayer.Delete(key)
					deletedCount++
				} else {
					newKeys = append(newKeys, key)
				}
			} else {
				newKeys = append(newKeys, key)
			}
		}
	}
	
	messageKeys = newKeys
	return deletedCount
}

func deleteAllMessagesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	messageKeysMutex.Lock()
	defer messageKeysMutex.Unlock()
	
	deletedCount := 0
	
	for _, key := range messageKeys {
		prevalenceLayer.Delete(key)
		deletedCount++
	}
	
	messageKeys = make([]string, 0)
	
	log.Printf("Deleted all %d messages", deletedCount)
	return mcp.NewToolResultText(fmt.Sprintf("Deleted all %d messages", deletedCount)), nil
}

func searchMessagesHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.GetArguments()

	keywordsArg, exists := args["keywords"]
	if !exists || keywordsArg == nil {
		return nil, fmt.Errorf("missing required parameter 'keywords'")
	}
	keywords, ok := keywordsArg.(string)
	if !ok {
		return nil, fmt.Errorf("parameter 'keywords' must be a string")
	}

	messages := getAllMessagesSorted()
	var matchingMessages []Message

	for _, message := range messages {
		if containsKeywords(message.Content, keywords) {
			matchingMessages = append(matchingMessages, message)
		}
	}

	if len(matchingMessages) == 0 {
		return mcp.NewToolResultText("No messages found matching the keywords"), nil
	}

	jsonData, err := json.Marshal(matchingMessages)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error marshaling messages: %v", err)), nil
	}

	log.Printf("Found %d messages matching keywords: %s", len(matchingMessages), keywords)
	return mcp.NewToolResultText(string(jsonData)), nil
}

func containsKeywords(content, keywords string) bool {
	contentLower := strings.ToLower(content)
	keywordsLower := strings.ToLower(keywords)
	keywordList := strings.Fields(keywordsLower)
	
	for _, keyword := range keywordList {
		if strings.Contains(contentLower, keyword) {
			return true
		}
	}
	return false
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]interface{}{
		"status": "healthy",
		"server": "mcp-memory-server",
	}
	json.NewEncoder(w).Encode(response)
}