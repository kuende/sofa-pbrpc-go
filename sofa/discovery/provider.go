package discovery

// A Provider is a interface that provides a list of peers for a client
type Provider interface {
	Hosts() ([]string, error)
}
