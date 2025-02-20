package reader

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/msquared-io/etherbase/backend/internal/config"
	"github.com/msquared-io/etherbase/backend/internal/sources"
)

// corsMiddleware adds CORS headers to all responses
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SetupRoutes(r *mux.Router) {
	// Apply CORS middleware to all routes
	r.Use(corsMiddleware)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"status":"healthy"}`)
	}).Methods("GET", "OPTIONS")

	r.HandleFunc("/etherbase-address", func(w http.ResponseWriter, r *http.Request) {
		log.Println("etherbase address")
		cfg, err := config.LoadConfig()
		if err != nil {
			http.Error(w, "Failed to load config", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cfg.EtherbaseAddress.Hex())
	}).Methods("GET", "OPTIONS")

	r.HandleFunc("/sources", func(w http.ResponseWriter, r *http.Request) {
		log.Println("sources")
		cfg, err := config.LoadConfig()
		if err != nil {
			http.Error(w, "Failed to load config", http.StatusInternalServerError)
			return
		}
 
		// Get sources from registry
		registry := sources.GetSourceRegistry()
		if registry == nil {
			registry = sources.NewSourceRegistry(cfg.EtherbaseAddress, nil)
		}

		// Get sources with owners and convert to response format
		sources := registry.GetSourcesWithOwners()
		sourcesResponse := make([]map[string]string, 0, len(sources))
		for _, info := range sources {
			sourcesResponse = append(sourcesResponse, map[string]string{
				"sourceAddress": info.SourceAddress.Hex(),
				"owner":         info.Owner.Hex(),
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(sourcesResponse)
	}).Methods("GET", "OPTIONS")

	r.HandleFunc("/event-definitions", func(w http.ResponseWriter, r *http.Request) {
		log.Println("source definitions")
		sourceAddress := r.URL.Query().Get("sourceAddress")
		if sourceAddress == "" {
			http.Error(w, "Missing 'sourceAddress' query param", http.StatusBadRequest)
			return
		}

		// Get source registry
		registry := sources.GetSourceRegistry()
		if registry == nil {
			http.Error(w, "Source registry not initialized", http.StatusInternalServerError)
			return
		}

		// Get event definitions from the contract
		definitions, err := registry.GetEventDefinitions(sourceAddress)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get event definitions: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(definitions)
	}).Methods("GET", "OPTIONS")

	r.HandleFunc("/emit-event", func(w http.ResponseWriter, r *http.Request) {
		// parse JSON, send to transaction pool
		w.Write([]byte("OK"))
	}).Methods("POST", "OPTIONS")

	r.HandleFunc("/set-state", func(w http.ResponseWriter, r *http.Request) {
		// parse JSON, send to transaction pool
		w.Write([]byte("OK"))
	}).Methods("POST", "OPTIONS")
}
