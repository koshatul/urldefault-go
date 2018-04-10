package urldefault_test

import (
	"net/url"

	"github.com/koshatul/urldefault-go/src/urldefault"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parse", func() {

	Context("supplying defaults", func() {
		url, err := urldefault.Parse("https://localhost/newpath", "http://example.com:8888/oldpath")

		It("should not error", func() {
			Expect(err).To(BeNil())
		})

		It("should return default port and supplied host", func() {
			Expect(url.Host).To(Equal("localhost:8888"))
			Expect(url.Host).To(Not(Equal("localhost")))
			Expect(url.Hostname()).To(Equal("localhost"))
			Expect(url.Port()).To(Equal("8888"))
		})

		It("should return supplied schema", func() {
			Expect(url.Scheme).To(Equal("https"))
		})

		It("should return supplied path", func() {
			Expect(url.Path).To(Equal("/newpath"))
		})

	})

	Context("supplying blank string", func() {
		url, err := urldefault.Parse("", "http://example.com:8888/oldpath")

		It("should not error", func() {
			Expect(err).To(BeNil())
		})

		It("should return default host and port", func() {
			Expect(url.Host).To(Equal("example.com:8888"))
			Expect(url.Hostname()).To(Equal("example.com"))
			Expect(url.Port()).To(Equal("8888"))
		})

		It("should return default schema", func() {
			Expect(url.Scheme).To(Equal("http"))
		})

		It("should return default path", func() {
			Expect(url.Path).To(Equal("/oldpath"))
		})
	})

	Context("invalid schema", func() {
		url, err := urldefault.Parse("localhost:9999/newpath", "http://example.com:8888/oldpath")

		It("should not error", func() {
			Expect(err).To(BeNil())
		})

		It("should return supplied host and port", func() {
			Expect(url.Host).To(Equal("localhost:9999"))
			Expect(url.Hostname()).To(Equal("localhost"))
			Expect(url.Port()).To(Equal("9999"))
		})

		It("should return default schema", func() {
			Expect(url.Scheme).To(Equal("http"))
		})

		It("should return supplied path", func() {
			Expect(url.Path).To(Equal("/newpath"))
		})
	})

	Context("no port", func() {
		url, err := urldefault.Parse("https://localhost/newpath", "http://example.com:8888/oldpath")

		It("should not error", func() {
			Expect(err).To(BeNil())
		})

		It("should return supplied host and default port", func() {
			Expect(url.Host).To(Equal("localhost:8888"))
			Expect(url.Hostname()).To(Equal("localhost"))
			Expect(url.Port()).To(Equal("8888"))
		})

		It("should return supplied schema", func() {
			Expect(url.Scheme).To(Equal("https"))
		})

		It("should return supplied path", func() {
			Expect(url.Path).To(Equal("/newpath"))
		})
	})

	Context("no path", func() {
		url, err := urldefault.Parse("https://localhost:9999/", "http://example.com:8888/oldpath")

		It("should not error", func() {
			Expect(err).To(BeNil())
		})

		It("should return supplied host and default port", func() {
			Expect(url.Host).To(Equal("localhost:9999"))
			Expect(url.Hostname()).To(Equal("localhost"))
			Expect(url.Port()).To(Equal("9999"))
		})

		It("should return supplied schema", func() {
			Expect(url.Scheme).To(Equal("https"))
		})

		It("should return supplied path", func() {
			Expect(url.Path).To(Equal("/oldpath"))
		})
	})

	Context("no user", func() {
		url, err := urldefault.Parse("https://localhost:9999/", "http://foo:bar@example.com:8888/oldpath")

		It("should not error", func() {
			Expect(err).To(BeNil())
		})

		It("should return supplied host and default port", func() {
			Expect(url.Host).To(Equal("localhost:9999"))
			Expect(url.Hostname()).To(Equal("localhost"))
			Expect(url.Port()).To(Equal("9999"))
		})

		It("should return supplied schema", func() {
			Expect(url.Scheme).To(Equal("https"))
		})

		It("should return supplied path", func() {
			Expect(url.Path).To(Equal("/oldpath"))
		})

		It("should return default user", func() {
			Expect(url.User).NotTo(BeNil())
			Expect(url.User.Username()).To(Equal("foo"))
			pass, ok := url.User.Password()
			Expect(pass).To(Equal("bar"))
			Expect(ok).To(BeTrue())
		})

	})

	Context("supplied user", func() {
		url, err := urldefault.Parse("https://faa:ber@localhost:9999/", "http://foo:bar@example.com:8888/oldpath")

		It("should not error", func() {
			Expect(err).To(BeNil())
		})

		It("should return supplied host and default port", func() {
			Expect(url.Host).To(Equal("localhost:9999"))
			Expect(url.Hostname()).To(Equal("localhost"))
			Expect(url.Port()).To(Equal("9999"))
		})

		It("should return supplied schema", func() {
			Expect(url.Scheme).To(Equal("https"))
		})

		It("should return supplied path", func() {
			Expect(url.Path).To(Equal("/oldpath"))
		})

		It("should return supplied user", func() {
			Expect(url.User).NotTo(BeNil())
			Expect(url.User.Username()).To(Equal("faa"))
			pass, ok := url.User.Password()
			Expect(pass).To(Equal("ber"))
			Expect(ok).To(BeTrue())
		})

	})

	Context("supplied user, no pass", func() {
		url, err := urldefault.Parse("https://faa@localhost:9999/", "http://foo:bar@example.com:8888/oldpath")

		It("should not error", func() {
			Expect(err).To(BeNil())
		})

		It("should return supplied host and default port", func() {
			Expect(url.Host).To(Equal("localhost:9999"))
			Expect(url.Hostname()).To(Equal("localhost"))
			Expect(url.Port()).To(Equal("9999"))
		})

		It("should return supplied schema", func() {
			Expect(url.Scheme).To(Equal("https"))
		})

		It("should return supplied path", func() {
			Expect(url.Path).To(Equal("/oldpath"))
		})

		It("should return supplied user and blank password", func() {
			Expect(url.User).NotTo(BeNil())
			Expect(url.User.Username()).To(Equal("faa"))
			pass, ok := url.User.Password()
			Expect(pass).To(Equal(""))
			Expect(ok).To(BeFalse())
		})

	})

	Context("dropin replacement for net/url", func() {
		url, urlerr := url.Parse("https://foo:bar@localhost:9999/path?query=something#fragment")
		It("url.Parse should not error", func() {
			Expect(urlerr).To(BeNil())
		})
		urldef, urldeferr := urldefault.Parse("https://foo:bar@localhost:9999/path?query=something#fragment")
		It("urldefault.Parse should not error", func() {
			Expect(urldeferr).To(BeNil())
		})

		It("should return the same result as the original url.Parse()", func() {
			Expect(urldef).To(Equal(url))
		})

	})

})
