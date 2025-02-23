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
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"status":"healthy"}`)
	}).Methods("GET", "OPTIONS")

	r.HandleFunc("/etherbase-address", func(w http.ResponseWriter, r *http.Request) {
		log.Println("etherbase address")
		cfg, err := config.LoadConfig()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to load config"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cfg.EtherbaseAddress.Hex())
	}).Methods("GET", "OPTIONS")

	r.HandleFunc("/sources", func(w http.ResponseWriter, r *http.Request) {
		log.Println("sources")
		cfg, err := config.LoadConfig()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to load config"})
			return
		}
 
		// Get sources from registry
		registry := sources.GetSourceRegistry()
		if registry == nil {
			registry = sources.NewSourceRegistry(cfg.EtherbaseAddress)
		}

		// Get sources with owners and convert to response format
		sources := registry.GetSources()
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

	r.HandleFunc("/custom-contracts", func(w http.ResponseWriter, r *http.Request) {
		log.Println("custom contracts")
		cfg, err := config.LoadConfig()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to load config"})
			return
		}

		// Get custom contracts from registry
		registry := sources.GetSourceRegistry()
		if registry == nil {
			registry = sources.NewSourceRegistry(cfg.EtherbaseAddress)
		}

		// Get custom contracts
		customContracts := registry.GetCustomContracts()
		customContractsResponse := make([]map[string]string, 0, len(customContracts))
		for _, contract := range customContracts {
			customContractsResponse = append(customContractsResponse, map[string]string{
				"contractAddress": contract.ContractAddress.Hex(),
				"addedBy":         contract.AddedBy.Hex(),
				"contractABI":     contract.ContractABI,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customContractsResponse)
	}).Methods("GET", "OPTIONS")

	r.HandleFunc("/event-definitions", func(w http.ResponseWriter, r *http.Request) {
		log.Println("source definitions")
		sourceAddress := r.URL.Query().Get("sourceAddress")
		if sourceAddress == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Missing 'sourceAddress' query param"})
			return
		}

		// Get source registry
		registry := sources.GetSourceRegistry()
		if registry == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Source registry not initialized"})
			return
		}

		// Get event definitions from the contract
		definitions, err := registry.GetEventDefinitions(sourceAddress)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Failed to get event definitions: %v", err)})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(definitions)
	}).Methods("GET", "OPTIONS")
}
