package ovirt

import (
	"os"

	"github.com/openshift/installer/pkg/platformv2/abstract"
)

type ovirtFactory struct {
}

func (o *ovirtFactory) Create() (abstract.PlatformV2, error) {
	_, ok := os.LookupEnv("OPENSHIFT_INSTALL_OS_IMAGE_OVERRIDE")
	return &ovirtPlatform{
		removeTemplate: !ok,
	}, nil
}
