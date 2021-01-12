package vagrant

import (
	"fmt"
	"path/filepath"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type LXCProvider struct{}

func (p *LXCProvider) KeepInputArtifact() bool {
	return false
}

func (p *LXCProvider) Process(ui packersdk.Ui, artifact packersdk.Artifact, dir string) (vagrantfile string, metadata map[string]interface{}, err error) {
	// Create the metadata
	metadata = map[string]interface{}{
		"provider": "lxc",
		"version":  "1.0.0",
	}

	// Copy all of the original contents into the temporary directory
	for _, path := range artifact.Files() {
		ui.Message(fmt.Sprintf("Copying: %s", path))

		dstPath := filepath.Join(dir, filepath.Base(path))
		if err = CopyContents(dstPath, path); err != nil {
			return
		}
	}

	return
}
