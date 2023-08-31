package handler

import (
	"github.com/RCNRC/dynamic_segmentation/pkg/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	serviceRouter := r.PathPrefix("/segment").Subrouter()
	serviceRouter.HandleFunc("/", h.createSegment).Methods("POST")
	serviceRouter.HandleFunc("/", h.deleteSegment).Methods("DELETE")

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/update/", h.update).Methods("PUT")
	userRouter.HandleFunc("/history/", h.history).Methods("GET")
	userRouter.HandleFunc("/report/{fileName}", h.report).Methods("GET")
	userRouter.HandleFunc("/current/", h.current).Methods("GET")
	return r
}
