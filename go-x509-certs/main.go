package main

import (
	"crypto/x509"
	"crypto/rand"
	"crypto/x509/pkix"
	"crypto/rsa"
	"math/big"
	"time"
	"fmt"
)

const (
	threeMonthDuration = time.Hour * 24 * 30 * 4
)

var (
	serial int64 = 0
)

func main() {
	var err error
	roots := x509.NewCertPool()

	var caCert *x509.Certificate
	var caKey *rsa.PrivateKey
	{
		caKey, err = newKey()
		if err != nil {
			panic(err)
		}

		caCert, err = newCACertificate(caKey)
		if err != nil {
			panic(err)
		}

	}
	roots.AddCert(caCert)

	var exampleComCertReq *x509.CertificateRequest
	var exampleComKey *rsa.PrivateKey
	{
		exampleComKey, err = newKey()
		if err != nil {
			panic(err)
		}

		exampleComCertReq, err = newCertRequest([]string{"*.example.com", "example.com"}, exampleComKey)
		if err != nil {
			panic(err)
		}
	}

	var exampleComCert *x509.Certificate
	{
		exampleComCert, err = newCert(caCert, exampleComCertReq, caKey)
		if err != nil {
			panic(err)
		}
	}

	_, err = exampleComCert.Verify(x509.VerifyOptions{
		Roots:   roots,
		DNSName: "some-sub-domain.example.com",
	})
	if err != nil {
		panic(err)
	}

	{
		encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, exampleComCert.PublicKey.(*rsa.PublicKey), []byte("test message 1!"))
		if err != nil {
			panic(err)
		}

		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, exampleComKey, encrypted)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(decrypted))
	}
}

func newCert(ca *x509.Certificate, req *x509.CertificateRequest, key *rsa.PrivateKey) (cert *x509.Certificate, err error) {
	serial += 1

	template := x509.Certificate{
		SerialNumber: big.NewInt(serial),
		DNSNames:     req.DNSNames,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(threeMonthDuration),
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, ca, req.PublicKey, key)
	if err != nil {
		return
	}

	return x509.ParseCertificate(certBytes)
}

func newCertRequest(dnsNames []string, key *rsa.PrivateKey) (req *x509.CertificateRequest, err error) {
	template := x509.CertificateRequest{
		DNSNames:  dnsNames,
		PublicKey: &key.PublicKey,
	}

	bytes, err := x509.CreateCertificateRequest(rand.Reader, &template, key)
	if err != nil {
		return
	}

	return x509.ParseCertificateRequest(bytes)
}

func newKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func newCACertificate(key *rsa.PrivateKey) (cert *x509.Certificate, err error) {
	serial += 1

	template := x509.Certificate{
		SerialNumber: big.NewInt(serial),
		Subject: pkix.Name{
			CommonName: "CA Root Certificate",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(threeMonthDuration),
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		return
	}

	return x509.ParseCertificate(certBytes)
}
