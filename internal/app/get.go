package app

import (
	"errors"
	"fmt"
	"time"

	"bs/configs"
	"bs/internal/ipfs"

	files "github.com/ipfs/go-ipfs-files"
	cid "github.com/ipfs/go-cid"
)

type fetchreport struct {
	Fn	files.Node
	Err	error
}

func fetchFileFromIPFS(ctx *ipfs.IPFSCtx, c cid.Cid) (*files.Node, error) {
	fst,err := time.ParseDuration(configs.FileSearchTimeout)
	if err != nil {
		return nil,err
	}

	timeoutc := make(chan fetchreport,1)
	go func() {
		fn, err := ipfs.GetFile(ctx, c)
		timeoutc <- fetchreport{fn, err}
	}()

	select {
	case res := <-timeoutc: 
		return &res.Fn, res.Err
	case <-time.After(fst):
		return nil, errors.New("File search timed out")
	}


}

func SaveFile(fn *files.Node, n string) error {
	err := ipfs.SaveFile(*fn,n)
	return err
}

func Download(ctx *ipfs.IPFSCtx, path string, tag string) {
	cid, err := ipfs.GetCID(path)
	if err != nil {
		 fmt.Printf("Failed to convert CID: %s\n",err)
		 return
	}

	_, err = fetchFileFromIPFS(ctx, cid)
	if err != nil {
		fmt.Printf("Error downloading file: %s\n", err)
		return
	}
}

