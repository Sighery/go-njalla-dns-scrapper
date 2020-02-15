package provider

import "fmt"

// Provider struct
type Provider struct {
	URL string
}

// Providers interface implemented for providers
type Providers interface {
	getUrl()
	getDomainURL()
	login()
}

// New constructor
func New(url string) *Provider {
	return &Provider{
		URL: url,
	}
}

func (p *Provider) getURL(path string) string {
	return fmt.Sprintf("%s%s", p.URL, path)
}

func (p *Provider) getDomainURL() string {
	return p.URL
}
