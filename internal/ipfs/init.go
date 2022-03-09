package ipfs

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	config "github.com/ipfs/go-ipfs-config"
	libp2p "github.com/ipfs/go-ipfs/core/node/libp2p"
	icore "github.com/ipfs/interface-go-ipfs-core"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/plugin/loader" // This package is needed so that all the preloaded plugins are loaded automatically
	"github.com/ipfs/go-ipfs/repo/fsrepo"
)

type IPFSCtx struct {
	Ctx	context.Context
	Cancel	context.CancelFunc
	IPFS	icore.CoreAPI
}

func New() (*IPFSCtx, error) {
	ctx, cancel := context.WithCancel(context.Background())
	
	ipfs, err := spawnEphemeral(ctx)

	ipfsCtx := IPFSCtx{
		Ctx: ctx,
		Cancel: cancel,
		IPFS: ipfs,
	}

	return &ipfsCtx, err
}

func createNode(ctx context.Context, path string) (icore.CoreAPI, error) {
	repo, err := fsrepo.Open(path)
	if err != nil {
		return nil, err
	}

	nodeOptions := &core.BuildCfg{
		Online: true,
		Routing: libp2p.DHTOption,
		Repo: repo,
	}

	node, err := core.NewNode(ctx, nodeOptions)
	if err != nil {
		return nil, err
	}

	return coreapi.NewCoreAPI(node)
}

func createTempRepo(ctx context.Context) (string, error) {
	repoPath, err := getRepoPath()
	if err != nil {
		return "", fmt.Errorf("failed to get tempdir: %s", err)
	}

	cfg, err := config.Init(ioutil.Discard, 2048)
	if err != nil {
		return "", err
	}

	err = fsrepo.Init(repoPath, cfg)
	if err != nil {
		return "", fmt.Errorf("Failed to init ephemeral node: %s", err)
	}

	return repoPath, nil
}

func getRepoPath() (string, error) {
	return fsrepo.BestKnownPath()
}

func setupPlugins(externalPluginsPath string) error {
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))

	if err != nil {
		return fmt.Errorf("Error loading plugins: %s",err)
	}

	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("Error initializing plugins: %s",err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("Error injecting plugins: %s", err)
	}
	
	return nil
}

func spawnEphemeral(ctx context.Context) (icore.CoreAPI, error) {
	if err := setupPlugins(""); err != nil {
		return nil, err
	}

	repoPath, err := createTempRepo(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to create temp repo: %s\n", err)
	}
	return createNode(ctx,repoPath)
}
