package vsphere

import (
	"github.com/solo-io/unik/pkg/types"
)

func (p *VsphereProvider) ListImages() ([]*types.Image, error) {
	images := []*types.Image{}
	for _, image := range p.state.GetImages() {
		images = append(images, image)
	}
	return images, nil
}
