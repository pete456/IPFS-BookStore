package ipfs

import (
	cid "github.com/ipfs/go-cid"
)

func GetCID(path string) (cid.Cid, error) {
	return cid.Decode(path)
}
