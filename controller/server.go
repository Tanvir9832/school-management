package controller

import "management/store"

type Server struct {
	db store.StoreOperations
}

func (s *Server) NewServer(pgSotre store.Postgres) {
	s.db = &pgSotre
	s.db.NewStore()
}

type ServerOperations interface {
	NewServer(pgSotre store.Postgres)
}
