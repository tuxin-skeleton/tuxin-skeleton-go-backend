package main

import (
	"crypto/tls"
	"log"
	"crypto/x509"
	"io/ioutil"
)

var TlsConfig *tls.Config;

func InitTlsConfig() {
	cert, err := tls.LoadX509KeyPair(CfgIni.CertificateFile,CfgIni.PrivateKeyFile)
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)

	}
	certpool := x509.NewCertPool()
	for _, crFile := range CfgIni.OtherCertificates {
		pem, err := ioutil.ReadFile(crFile)
		if err != nil {
			log.Fatalf("Failed to read client certificate authority: %v", err)
		}
		if !certpool.AppendCertsFromPEM(pem) {
			log.Fatalf("Can't parse client certificate authority")
		}
	}

	TlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    certpool,
	}
	TlsConfig.BuildNameToCertificate()
}
