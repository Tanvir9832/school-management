package controller

import (
	"management/model"
	"management/store"
	"management/utils"
)

type Server struct {
	db store.StoreOperations
}

func (s *Server) NewServer(pgSotre store.Postgres) {
	utils.SetLogger()

	utils.Logger.Infof("Logger setup done ... \n")
	s.db = &pgSotre
	err := s.db.NewStore()
	if err != nil {
		utils.Log(model.LogLevelError, model.Controller, model.NewServer, "database connection failed!", err)
	} else {
		utils.Log(model.LogLevelInfo, model.Controller, model.NewServer, "database connected", nil)
	}
}

type ServerOperations interface {
	NewServer(pgSotre store.Postgres)
}
