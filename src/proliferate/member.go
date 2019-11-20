package proliferate

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"path"
	//"math/big"
)

/*
type Permission struct {
	Root big.Int
}
*/

type member struct {
	cert string
	key  string
}

/*
func (node *Node) LoadKeyPair() {
	n := *node
}
*/

// CertificateLoad attaches node certificates to n.member
func (node *Node) CertificateLoad() {
	n := *node
	c := n.Config.Static

	certFile := path.Join(c.IdentityFolder, c.CertFile)
	keyFile := path.Join(c.IdentityFolder, c.KeyFile)

	cert, err := ioutil.ReadFile(certFile)
	if err != nil {
		n.Log(Message{
			Level: 2,
			Text:  err.Error(),
		})
	} else {
		n.member.cert = string(cert)
	}

	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		n.Log(Message{
			Level: 2,
			Text:  err.Error(),
		})
	} else {
		n.member.key = string(key)
	}

	*node = n
}

// ExtractPublicKey returns rsa.PublicKey from root pem
func ExtractPublicKey(pemKey string) rsa.PublicKey {
	block, _ := pem.Decode([]byte(pemKey))
	var cert *x509.Certificate
	cert, _ = x509.ParseCertificate(block.Bytes)
	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)
	return *rsaPublicKey
}
