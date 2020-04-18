package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"server/api"
	"server/services"
	"server/utils"
)

type Server struct {
	ClientTokens []string
	serveMux     *mux.Router
	httpServer   *http.Server
	done         chan int
}

func NewServer(clientTokens []string, service *services.ReportService) (*Server, error) {
	if len(clientTokens) == 0 {
		return nil, errors.New("no one client's token provided")
	}

	for _, clientToken := range clientTokens {
		if len(clientToken) == 0 {
			return nil, errors.New("client's token can not be empty")
		}
	}

	server := Server{
		ClientTokens: clientTokens,
	}

	route := mux.NewRouter()
	route.Use(server.proxyHandler)
	route.HandleFunc("/report/", api.NewReportController(service).Handle).Methods(http.MethodPost)
	server.serveMux = route

	return &server, nil
}

func (this *Server) proxyHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		this.proxy(writer, request, handler)
	})
}

func (this *Server) proxy(writer http.ResponseWriter, request *http.Request, handler http.Handler) {
	currentToken := request.Header.Get("Access-Token")

	forbidden := true
	for _, allowedToken := range this.ClientTokens {
		if currentToken == allowedToken {
			forbidden = false
		}
	}

	if forbidden {
		utils.Forbidden(writer)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	handler.ServeHTTP(writer, request)
}

func (this *Server) Start(port int) {
	host := fmt.Sprintf(":%d", port)

	tcpListener, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Println(err)
		panic("TCP listener wasn't created")
	}

	this.httpServer = &http.Server{
		Addr:    host,
		Handler: this.serveMux,
	}

	go this.httpServer.Serve(tcpListener)
	fmt.Printf("HTTP server started on %d\n", port)
}

func (this *Server) Stop() {
	this.httpServer.Close()

	fmt.Printf("HTTP server stopped\n")
	if this.done != nil {
		this.done <- 1
	}

}
