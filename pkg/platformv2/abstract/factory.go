package abstract

// PlatformFactory is a factory that provides a platform implementation
type PlatformFactory interface {
	// Create creates a platform API, or returns an error if it is not possible.
	Create() (PlatformV2, error)
}
