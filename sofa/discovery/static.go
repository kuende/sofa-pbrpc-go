package discovery

// StaticProvider implements discovery provider based on static array
type StaticProvider struct {
	hosts []string
}

// NewStaticProvider returns a new static discovery provider
func NewStaticProvider(hosts []string) Provider {
	return &StaticProvider{hosts: hosts}
}

// Hosts returns a list of for a service
func (e *StaticProvider) Hosts() ([]string, error) {
	return e.hosts, nil
}
