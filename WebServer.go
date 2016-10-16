package main

import (
	"log"
	"strconv"
	"net/http"
	"github.com/urfave/negroni"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/unrolled/secure"
	"time"
	"github.com/gorilla/mux"
	"github.com/jingweno/negroni-gorelic"
)

var secureMiddleware *secure.Secure;

func angular2Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, CfgIni.WebServer.ClientPath + "/index.html")
}

func main() {
	CfgIni = parseIni(configIniPath);

	InitTlsConfig()

	secureMiddleware= secure.New(secure.Options{
	IsDevelopment: !CfgIni.Server.IsProduction,
	SSLRedirect:           true,
	SSLHost:               CfgIni.ServerName,
	SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	STSSeconds:            315360000,
	STSIncludeSubdomains:  true,
	STSPreload:            true,
	FrameDeny:             true,
	ContentTypeNosniff:    true,
	BrowserXssFilter:      true,
	ContentSecurityPolicy: "default-src 'self'; img-src 'self' cdn.auth0.com; connect-src 'self' " + CfgIni.AccountDomain + ";style-src 'self' 'unsafe-inline' fonts.googleapis.com; font-src 'self' fonts.gstatic.com;script-src 'self' 'unsafe-eval' cdn.auth0.com;",
	PublicKey:             `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`,
})
	log.Print("started web server...");
	httpsPortStr := ":" + strconv.FormatUint(CfgIni.HttpsPort, 10)
	log.Printf("starting https web server at port %v", CfgIni.HttpsPort)
	r := mux.NewRouter()
	r.Path("/").HandlerFunc(angular2Handler)
	for _, routeName := range CfgIni.WebServer.Routes {
		r.PathPrefix("/" + routeName).HandlerFunc(angular2Handler)
	}
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(CfgIni.WebServer.ClientPath)))
	n := negroni.New()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(r)
	n.Use(negroni.HandlerFunc(secureMiddleware.HandlerFuncWithNext))
	if CfgIni.IsProduction {
		n.Use(negronigorelic.New(CfgIni.Licensekey, CfgIni.AppName, true))
	}
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	srv := &http.Server{
		Addr: httpsPortStr,
		Handler: n,
		ReadTimeout: time.Duration(CfgIni.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(CfgIni.WriteTimeout) * time.Second,
		TLSConfig: TlsConfig,
	}
	err := srv.ListenAndServeTLS(CfgIni.CertificateFile,CfgIni.PrivateKeyFile)
	if err != nil {
		log.Fatalf("https server stopped with the following error: %v", err)
	} else {
		log.Print("https server stopped with no error")
	}

}
