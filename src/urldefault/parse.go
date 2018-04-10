package urldefault

import (
	"fmt"
	"net/url"
	"strings"
)

// Parse parses rawurl into a URL structure, taking a second parameter as defaults
// to override the main string with missing properties.
//
// Works as a drop in replacement for `url.Parse()`, as per `url.Parse`
// the urls may be relative (a path, without a host) or absolute (starting with a scheme).
// Trying to parse a hostname and path without a scheme is invalid but may
// not necessarily return an error, due to parsing ambiguities.
func Parse(rawurl ...string) (*url.URL, error) {
	if !strings.Contains(rawurl[0], "://") {
		rawurl[0] = fmt.Sprintf("override://%s", rawurl[0])
	}
	o, err := url.Parse(rawurl[0])
	if err != nil {
		return nil, err
	}

	if len(rawurl) < 2 {
		return o, err
	}

	var d *url.URL
	d, err = url.Parse(rawurl[1])
	if err != nil {
		return nil, err
	}

	if o.Scheme == "" || o.Scheme == "override" {
		o.Scheme = d.Scheme
	}
	if o.Opaque == "" {
		o.Opaque = d.Opaque
	}
	if o.User == nil {
		o.User = d.User
	}
	if o.Host == "" {
		o.Host = d.Host
	}
	if !strings.Contains(o.Host, ":") && strings.Contains(d.Host, ":") {
		if o.Host != d.Host {
			p := strings.SplitN(d.Host, ":", 2)
			o.Host = fmt.Sprintf("%s:%s", o.Host, p[1])
		}
	}
	if o.Path == "" || o.Path == "/" {
		o.Path = d.Path
		o.RawPath = d.RawPath
	}
	if o.Fragment == "" {
		o.Fragment = d.Fragment
	}
	return o, nil
}
