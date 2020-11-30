package platformv2

// PlatformV2 is an abstraction for supported platforms.
type PlatformV2 interface {
	// Name returns the name of the platform.
	Name() string

	// SupportsIPI returns true if the platform supports IPI.
	SupportsIPI() bool

	// GetIPI returns the handler for IPI-specific calls.
	//
	// Returns platformv2errors.ErrPlatformDoesNotSupportIPI if the platform does not support IPI.
	GetIPI() (IPIPlatform, error)
}
