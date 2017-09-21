package options

import (
	"github.com/spf13/pflag"
)

type Config struct {
	ListenAddress string
	TlsCertFile   string
	TlsPrivateKey string
	KeystoneUrl   string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) AddFlag(fs *pflag.FlagSet) {
	fs.StringVar(&c.ListenAddress, "listen", "127.0.0.1:8443", "address:port to listen on")
	fs.StringVar(&c.TlsCertFile, "tls-cert-file", "", "File containing the default x509 Certificate for HTTPS")
	fs.StringVar(&c.TlsPrivateKey, "tls-private-key-file", "", "File containing the default x509 private key matching --tls-cert-file")
	fs.StringVar(&c.KeystoneUrl, "keystone-url", "http://localhost/identity/v3/", "url for openstack keystone")
}
