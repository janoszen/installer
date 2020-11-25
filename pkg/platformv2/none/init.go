package ovirt

import (
	"github.com/openshift/installer/pkg/platformv2"
	"github.com/openshift/installer/pkg/types/none"
)

func init() {
	platformv2.Register(none.Name, &noneFactory{})
}
