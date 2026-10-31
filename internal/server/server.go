package server

import (
	"fmt"
	"go-linkr/internal/config"
	"go-linkr/internal/db"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

type Server struct {
	cfg *config.Config
	mux *chi.Mux
	db  *db.DB
}

func NewServer(cfg *config.Config) *Server {
	mux := chi.NewRouter()

	return &Server{
		cfg: cfg,
		mux: mux,
	}
}

func (s *Server) Run() error {
	openedDB, err := db.Open(s.cfg.DbPath)
	if err != nil {
		return err
	}
	s.db = openedDB

	s.setupHandlers()

	addr := fmt.Sprintf("0.0.0.0:%s", s.cfg.Port)
	log.Printf("Listening on %s...", addr)
	return http.ListenAndServe(addr, s.mux)
}

func (s *Server) setupHandlers() {
	s.mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	})

	s.mux.Get("/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")
		link, err := s.db.FindLinkBySlug(slug)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		target := link.Target
		if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
			target = "https://" + target
		}

		http.Redirect(w, r, target, http.StatusFound)
	})
}
