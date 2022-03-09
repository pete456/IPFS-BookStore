package ipfs

import (
	cid "github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
)

func GetFile(ctx *IPFSCtx, path cid.Cid) (files.Node, error) {
	rootNodeFile, err := ctx.IPFS.Unixfs().Get(ctx.Ctx,icorepath.IpfsPath(path))
	if err != nil {
		return nil, err
	}
	return rootNodeFile, nil
}

func SaveFile(fn files.Node, savepath string) error {
	err := files.WriteTo(fn,savepath)
	return err
}
