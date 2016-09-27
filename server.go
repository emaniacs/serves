package main

type Server struct {
	hostname string
	path string
	port int
}

func (server *Server) SetHostname (hostname string) {
	server.hostname = hostname
}
func (server *Server) GetHostname () string {
	return server.hostname
}
func (server *Server) SetPort (port int) {
	server.port = port
}
func (server *Server) GetPort () int {
	return server.port
}
func (server *Server) SetPath (path string) {
	server.path = path
}
func (server *Server) GetPath () string {
	return server.path
}
