# OpenShift Installer Platform API V2

This folder contains the API that platforms need to implement to integrate with the OpenShift installer. The primary integration point is located in [platformv2.go](platformv2.go). Any platform wishing to integrate must implement this interface.

Once you have the basic functions implemented you should also add the platform to [platformv2registry/platforms.go](platformv2registry/platforms.go):

```go
func init()  {
	Register(none.Name, &nonePlatform.Factory{})
	Register(ovirt.Name, &ovirtPlatform.Factory{})
    //<-- add yours here
}
```

## Replacing legacy calls

Legacy API calls usually consist of a large amount of switch-case statements:

```go
switch platform {
case aws.Name:
    //...
case azure.Name:
    //...
case baremetal.Name:
    //...
}
```

These constructs should be wrapped as follows:

```go
p, err := platformv2registry.GetByName(installConfig.Config.Platform.Name())
if err == nil {
    p.CallReplacementMethod()
} else if !errors.Is(err, platformv2errors.ErrPlatformNotRegistered) {
    // handle other error
} else {
    switch (platform) {
        // Legacy calls that have not been ported yet
    }
}
```
