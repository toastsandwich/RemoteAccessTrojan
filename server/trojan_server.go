package main

import "net"

type TrojanServer struct {
	Host     string
	Port     string
	Listener net.Listener
	Victim   net.Conn
}

func NewTrojanServer(config *Config) *TrojanServer {
	return &TrojanServer{
		Host:     config.Server.Host,
		Port:     config.Server.Port,
		Listener: nil,
		Victim:   nil,
	}
}

func (s *TrojanServer) Run() error {
	ln, err := net.Listen("tcp", s.Host+s.Port)
	if err != nil {
		return err
	}
	s.Listener = ln
	conn, err := s.Listener.Accept()
	if err != nil {
		return err
	}
	s.Victim = conn
	return handleConn(s.Victim)
}

func handleConn(conn net.Conn) error {
	// to do 
	return nil
}
