package com

import (
	"bs/internal/files"
	"net"
)

type Request interface {
	Execute() []byte
}

// Request structs
type base struct {
	Rtype string
	Path string
}

type Get struct {
	base
}

type Put struct {
	base
	files.FileData
}

type Move struct {
	//TODO
	base
}

func Execute(r Request, c net.Conn) {
	data := r.Execute()
	_, err := c.Write(data)
	if err != nil {
	}
	c.Close()
}

func (g Get)Execute() []byte {
	//TODO
	return []byte("")
}

