package lingua

type componentUnavailable interface {
	error
	Component() string
}
