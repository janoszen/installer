# OpenShift Installer Platform API V2

This folder contains the API that platforms need to implement to integrate with the OpenShift installer. The primary integration point is located in [abstract/abstraction.go](abstract/abstraction.go). Any platform wishing to integrate must implement this interface.

Platforms should also define an init function as follows:

```go
func init() {
	platformv2.Register(none.Name, &noneFactory{})
}
```

Once this method is added the platform package should be imported in [register.go](register.go) to enable calling the `init()` function:

```go
import (
	// Importing platforms for registration
	_ "github.com/openshift/installer/pkg/platformv2/none"
	_ "github.com/openshift/installer/pkg/platformv2/ovirt"
    // <--- Add yours here
)
```
