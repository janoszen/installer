package platformv2registry

import (
	"github.com/pkg/errors"

	"github.com/openshift/installer/pkg/platformv2"
	"github.com/openshift/installer/pkg/platformv2/platformv2errors"
)

// GetByName returns a platform by its name.
//
// Typical usage:
//
// p, err := platformv2registry.Get(installConfig.Config.Platform.Name())
//
// returns platformv2errors.ErrPlatformNotRegistered if the platform is not registered.
func GetByName(platform string) (platformv2.PlatformV2, error) {
	if platform == "" {
		return nil, errors.Errorf("platform name is empty")
	}
	if factory, ok := platforms[platform]; ok {
		return factory.Create()
	}
	return nil, platformv2errors.ErrPlatformNotRegistered
}
