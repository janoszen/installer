package platformv2

import (
	"github.com/pkg/errors"

	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/platformv2/abstract"
)

// NotRegistered is returned if a platform with the specified name is not registered.
var NotRegistered = errors.New("platform is not registered")

// Get creates a platform with the name specified in the install config.
func Get(installConfig *installconfig.InstallConfig) (abstract.PlatformV2, error) {
	return GetByName(installConfig.Config.Platform.Name())
}

//GetByName returns a platform by its name
func GetByName(platform string) (abstract.PlatformV2, error) {
	if platform == "" {
		return nil, errors.Errorf("platform name is empty")
	}
	if factory, ok := platforms[platform]; ok {
		return factory.Create()
	}
	return nil, NotRegistered
}
