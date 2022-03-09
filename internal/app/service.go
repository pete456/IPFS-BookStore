package app

import (
	"bs/internal/ipfs"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	IPFS *ipfs.IPFSCtx
	Log *log.Logger
}
