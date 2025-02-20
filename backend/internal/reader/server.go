package reader

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"github.com/msquared-io/etherbase/backend/internal/client"
	"github.com/msquared-io/etherbase/backend/internal/config"
	"github.com/msquared-io/etherbase/backend/internal/schema"
	"github.com/msquared-io/etherbase/backend/internal/sources"
	"github.com/msquared-io/etherbase/backend/internal/streaming"
	"github.com/msquared-io/etherbase/backend/internal/subscription"
)

type Server struct {
	cfg     *config.Config
	router  *mux.Router
	srv     *http.Server
	sView   *streaming.StreamingView
	manager *client.Manager
}

func NewServer(cfg *config.Config) (*Server, error) {
	manager, err := client.NewManager(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create client manager: %w", err)
	}
	sProv := schema.NewSchemaProvider(cfg.EtherbaseAddress, manager.Client())

	eReg := sources.NewSourceRegistry(cfg.EtherbaseAddress, nil)
	updateCh := eReg.OnUpdate()
	go func() {
		for newSources := range updateCh {
			fmt.Printf("Received new sources from registry: %v\n", newSources)
			for _, source := range newSources {
				sProv.AddSource(source)
			}
		}
	}()

	subscription.NewManager() 

	schema.NewSchemaProvider(cfg.EtherbaseAddress, manager.Client())
	streamingView := streaming.NewStreamingView(manager.Client(), BroadcastUpdates)

	r := mux.NewRouter()
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	SetupRoutes(r)

	// Set up routes for subscriptions
	r.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		HandleReadWebSocket(conn, r)
	})

	// Create server with proper address format
	addr := fmt.Sprintf(":%s", cfg.ReaderPort)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return &Server{
		cfg:     cfg,
		router:  r,
		srv:     srv,
		sView:   streamingView,
		manager: manager,
	}, nil
}

func (s *Server) Start() error {
	// Start streaming view
	go s.sView.Start(context.Background())
	
	// Start HTTP server
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	s.sView.Stop()
	s.manager.Close()
	return s.srv.Shutdown(ctx)
}

func (s *Server) Router() *mux.Router {
	return s.router
} 