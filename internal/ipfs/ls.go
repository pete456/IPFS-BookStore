package ipfs

import (
	"fmt"
	cid "github.com/ipfs/go-cid"
	icore "github.com/ipfs/interface-go-ipfs-core"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
)

type node struct {
	Ent		icore.DirEntry
	Children	map[cid.Cid]node
}

/* Creates tree of directorys and files */
func itterate(ctx *IPFSCtx, c cid.Cid) map[cid.Cid]node {
	cns := make(map[cid.Cid]node)
	del,err := LsCid(ctx, c)
	if err != nil {
		fmt.Printf("%s\n",err)
		panic("")
	}
	for _,cde := range del {
		nn := node{Ent:cde}
		if cde.Type == icore.TDirectory {
			nn.Children = itterate(ctx,cde.Cid)
		}
		cns[cde.Cid]=nn
	}
	return cns
}

func Printn(n node, path string) {
	for k := range n.Children {
		cn := n.Children[k].Ent
		fmt.Printf("path/name %v/%v, %v, cid %v\n",path,cn.Name,cn.Type,cn.Cid)
		if cn.Type == icore.TDirectory {
			Printn(n.Children[k],path+"/"+cn.Name)
		}
	}
}

func IndexDir(ctx *IPFSCtx, lbrycid string) node {
	rcid, err := GetCID(lbrycid)
	if err != nil {
		fmt.Printf("Failed to convert cid: %s\n",err)
	}
	rootnode := node{}
	rootnode.Children = itterate(ctx,rcid)
	return rootnode
}

func LsCid(ctx *IPFSCtx, path cid.Cid) ([]icore.DirEntry, error) {
	c, err := ctx.IPFS.Unixfs().Ls(ctx.Ctx,icorepath.IpfsPath(path))
	if err != nil {
		return nil, err
	}
	nodes := make([]icore.DirEntry,0)
	for i := range c {
		nodes = append(nodes,i)
	}
	return nodes, nil
}


