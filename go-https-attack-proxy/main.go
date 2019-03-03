package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	certCache = make(map[string]*tls.Certificate)
)

func main() {
	catls, ca, err := loadCA()
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    ":9000",
		Handler: http.HandlerFunc(Handler(catls, ca)),
	}

	fmt.Println("start listen")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func loadCA() (*tls.Certificate, *x509.Certificate, error) {
	catls, err := tls.LoadX509KeyPair("ca.crt", "ca.key")
	if err != nil {
		return nil, nil, err
	}

	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		return nil, nil, err
	}

	return &catls, ca, nil
}

func getCertificate(catls *tls.Certificate, ca *x509.Certificate) func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
		if cert, ok := certCache[info.ServerName]; ok {
			return cert, nil
		}

		// Prepare certificate
		cert := &x509.Certificate{
			SerialNumber: big.NewInt(1658),
			Subject: pkix.Name{
				Organization:       []string{"Smotrov Root CA"},
				Country:            []string{"xx"},
				OrganizationalUnit: []string{"Smotrov Root CA"},
				CommonName:         info.ServerName,
			},
			NotBefore:    time.Now(),
			NotAfter:     time.Now().AddDate(10, 0, 0),
			SubjectKeyId: []byte{1, 2, 3, 4, 6},
			DNSNames:     []string{info.ServerName},
			ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
			KeyUsage:     x509.KeyUsageDigitalSignature,
		}
		priv, _ := rsa.GenerateKey(rand.Reader, 2048)
		pub := &priv.PublicKey

		// Sign the certificate
		cert_b, err := x509.CreateCertificate(rand.Reader, cert, ca, pub, catls.PrivateKey)
		if err != nil {
			return nil, err
		}

		// Public key
		certPEMBlock := bytes.NewBufferString("")
		if err != nil {
			return nil, err
		}
		pem.Encode(certPEMBlock, &pem.Block{Type: "CERTIFICATE", Bytes: cert_b})

		// Private key
		keyPEMBlock := bytes.NewBufferString("")
		if err != nil {
			return nil, err
		}

		pem.Encode(keyPEMBlock, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

		certtls, err := tls.X509KeyPair(certPEMBlock.Bytes(), keyPEMBlock.Bytes())

		certCache[info.ServerName] = &certtls

		return &certtls, err
	}
}

func Handler(catls *tls.Certificate, ca *x509.Certificate) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodConnect {
			// tunnel raw http
			return
		}

		w.WriteHeader(http.StatusOK)

		connSrv, _, err := w.(http.Hijacker).Hijack()
		if err != nil {
			log.Println(err)
			return
		}
		defer connSrv.Close()

		tlsSrv := tls.Server(connSrv, &tls.Config{
			GetCertificate: getCertificate(catls, ca),
		})
		defer tlsSrv.Close()

		err = tlsSrv.Handshake()
		if err != nil {
			log.Println(err)
			return
		}

		connCli, err := net.Dial("tcp", r.Host)
		if err != nil {
			log.Println(err)
			return
		}

		tlsCli := tls.Client(connCli, &tls.Config{
			ServerName: r.URL.Hostname(),
		})
		defer tlsCli.Close()

		err = tlsCli.Handshake()
		if err != nil {
			log.Println(err)
			return
		}

		wg := sync.WaitGroup{}
		wg.Add(2)

		reqReader, reqWriter := io.Pipe()
		respReader, respWriter := io.Pipe()

		go func() {
			defer wg.Done()
			defer reqWriter.Close()

			io.Copy(tlsCli, io.TeeReader(tlsSrv, reqWriter))
		}()

		go func() {
			defer wg.Done()
			defer respWriter.Close()

			io.Copy(tlsSrv, io.TeeReader(tlsCli, respWriter))
		}()

		go func() {
			io.Copy(os.Stderr, reqReader)
		}()

		go func() {
			io.Copy(os.Stdout, respReader)
		}()

		wg.Wait()
	}
}
