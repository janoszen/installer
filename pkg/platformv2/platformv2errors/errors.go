package platformv2errors

import (
	"errors"
)

// ErrPlatformNotRegistered is returned if a platform with the specified name is not registered.
var ErrPlatformNotRegistered = errors.New("platform is not registered")

// ErrPlatformDoesNotSupportIPI is returned when the platform does not support IPI.
var ErrPlatformDoesNotSupportIPI = errors.New("platform does not support IPI")
