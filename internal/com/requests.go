package com

import (
	"bs/internal/files"
	"net"
	"fmt"
)

type Request interface {
	Execute() []byte
}

// Request structs
type RawRequest struct {
	Rtype string
	Path  string
	Data  map[string]string
}
type base struct {
	Rtype string
	Path  string
}

type Library struct {
	base
	Operation string // UPLOAD | ADD | REMOVE | SYNC
	CID	  string
}

type Get struct {
	base
	Test string
	Test2 string
}

type MakeDirectory struct {

}

type MoveFile struct {
	base
}

type PutFile struct {
	base
	files.FileData
}


func Execute(r Request, c net.Conn) error {
	data := r.Execute()
	_, err := c.Write(data)
	return err
}

func (l Library)Execute() []byte {
	fmt.Println(l.Operation)	
	return []byte("")
}

func (g Get)Execute() []byte {
	//TODO
	return []byte("")
}

