package insecure

import (
	"crypto/tls"
	"crypto/x509"
	"log"
)

const certPEM = `-----BEGIN CERTIFICATE-----
<INSERT CERT.PEM VALUE HERE>
-----END CERTIFICATE-----

`

const keyPEM = `-----BEGIN PRIVATE KEY-----
<INSERT PRIVATE KEY HERE>
-----END PRIVATE KEY-----
`

var (
	// Cert is a self signed certificate
	Cert tls.Certificate
	// CertPool contains the self signed certificate
	CertPool *x509.CertPool
)

func init() {
	var err error
	Cert, err = tls.X509KeyPair([]byte(certPEM), []byte(keyPEM))
	if err != nil {
		log.Fatalln("Failed to parse key pair:", err)
	}
	Cert.Leaf, err = x509.ParseCertificate(Cert.Certificate[0])
	if err != nil {
		log.Fatalln("Failed to parse certificate:", err)
	}

	CertPool = x509.NewCertPool()
	CertPool.AddCert(Cert.Leaf)
}
