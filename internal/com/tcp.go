package com

import (
	"bs/internal/app"

	md "github.com/mitchellh/mapstructure"

	"bufio"
	"encoding/json"
	"errors"
	"net"
)

func sendError(err error, c net.Conn, srv *app.Service) {	
	srv.Log.Debug(err)
	c.Write([]byte(err.Error()))
	c.Close()
}

func handleConnections(srv *app.Service, c net.Conn) {
	// Read Data from connection
	jd, err := bufio.NewReader(c).ReadBytes('\n')
	if err != nil {
		sendError(err,c,srv)
		return
	}
	
	// Unmarshell data to json
	var req RawRequest
	err = json.Unmarshal(jd, &req)
	if err != nil {
		sendError(err,c,srv)
		return
	}

	switch req.Rtype {
	case "LIB":
		lb := Library{}
		lb.Rtype = req.Rtype
		lb.Path = req.Path
		err = md.Decode(req.Data,&lb)
		if err != nil {
			sendError(err,c,srv)
			return
		}
		err = Execute(lb,c)
	case "GET":
		gr := Get{}
		gr.Rtype = req.Rtype
		gr.Path = req.Path
		err = md.Decode(req.Data,&gr)
		if err != nil {
			sendError(err,c,srv)
			return
		}
		err = Execute(gr,c)
	case "PUT":
		srv.Log.Debug("PUT not implemented")
	default:
		sendError(errors.New("Invalid Request Type"),c,srv)
	}
	if err != nil {
		sendError(err,c,srv)
	} else {
		c.Close()
	}
}

func StartServer(srv *app.Service, port string) error {
	s, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	defer s.Close()

	for {
		conn, err := s.Accept()
		if err != nil {
			return err
		}
		go handleConnections(srv,conn)
	}
}
