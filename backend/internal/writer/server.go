package writer

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
	"github.com/msquared-io/etherbase/backend/internal/txpool"
)

type Server struct {
	cfg     *config.Config
	router  *mux.Router
	srv     *http.Server
	txPool  *txpool.TransactionPool
	manager *client.Manager
}

func NewServer(cfg *config.Config) (*Server, error) {
	r := mux.NewRouter()

	// Create eth clients
	manager, err := client.NewManager(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create client manager: %w", err)
	}

	txpool.MakeTransactionPool(cfg)

	// Set up websocket upgrader
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	txpool.GetTransactionPool()
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
	
	// Set up routes for write operations
	r.HandleFunc("/write", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		handleWriteWebSocket(conn, r)
	})

	// Create server with proper address format
	addr := fmt.Sprintf(":%s", cfg.WriterPort)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return &Server{
		cfg:     cfg,
		router:  r,
		srv:     srv,
		manager: manager,
	}, nil
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	s.manager.Close()
	return s.srv.Shutdown(ctx)
}

func (s *Server) Router() *mux.Router {
	return s.router
} 