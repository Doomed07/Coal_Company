package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	HTTPHandlers *HTTPHandlers
}

func NewSerber(httpHandlers *HTTPHandlers) *Server {
	return &Server{
		HTTPHandlers: httpHandlers,
	}
}

func (s *Server) ServerStart() error {
	router := mux.NewRouter()

	router.Path("/factory").Methods("GET").HandlerFunc(s.HTTPHandlers.AboutFactory)
	router.Path("/factory").Methods("DELETE").HandlerFunc(s.HTTPHandlers.StopFactory)
	router.Path("/factory/miners/classification").Methods("GET").HandlerFunc(s.HTTPHandlers.ShowMiners)
	router.Path("/factory/miners").Methods("POST").HandlerFunc(s.HTTPHandlers.BuyMiner)
	router.Path("/factory/miners").Methods("GET").Queries("working", "true").HandlerFunc(s.HTTPHandlers.WorkingMiners)
	router.Path("/factory/miners").Methods("GET").HandlerFunc(s.HTTPHandlers.AllMiners)
	router.Path("/factory/items").Methods("POST").HandlerFunc(s.HTTPHandlers.BuyItem)
	router.Path("/factory/items/classification").Methods("GET").HandlerFunc(s.HTTPHandlers.AllItems)
	router.Path("/factory/items").Methods("GET").HandlerFunc(s.HTTPHandlers.AllBoughtItems)

	if err := http.ListenAndServe(":9999", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		} else {
			fmt.Println("Fail to start server")
			return err
		}
	}

	return nil
}
