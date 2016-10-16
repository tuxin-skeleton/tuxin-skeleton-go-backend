# tuxin-skeleton-go-backend

this is a go backend part of a starter kit I created at https://github.com/tuxin-skeleton.
which wasn't that long ago :) so if you have any suggestions to improve it, 
please let me know.

included go features:

- https web server
- go-ini - ini parser
- negroni with the negroni-gorelic package
- mux


## how to build

### building go related packages

welp, first you need to have Go Installed :) 

this project requires the following packages as dependencies:

- github.com/go-ini/ini
- github.com/urfave/negroni
- github.com/phyber/negroni-gzip/gzip
- github.com/jingweno/negroni-gorelic
- github.com/unrolled/secure
- github.com/gorilla/mux


### 3 ways to install packages:

1. (recommended) use glide (https://github.com/Masterminds/glide)
2. just execute `go get <MODULE_NAME>` on each package in the list.
3. execute `go get ./...` on the root of the project

###- config.ini

please copy `config/config.ini.defaults` to `config/config.ini`, open it with your favorite text editor, and change it according to your configuration.

in general you have 3 categories in config.ini

- `Server` - which includes the host name of the server and a flag that indicates if it's in a production environment or not.
- `SslCert` - which contains the pem and key files locations to property start an HTTPS Server.
- `NewRelic` - which includes the license key of your newrelic account and the application name that will appear on the dashboard
- `WebServer` - which includes the https port to open, client project location,read & write timeouts for https requests
                and finally the routes to map to index.html.
- `Auth0` - account domain (to configure secure package to allow login requests)

### compile the go project

execute from the project's root directory the following command: `go build`

# Things to notice

##-  Go Package github.com/unrolled/secure

I've installed and configured for you this awesome security package for your web server.
you should check and understand the ContentSecurityPolicy of the secureMiddleware variable in the main go file (at src/main/WebServer.go).
I configured the content security policies to allow all the wonderful features that this project has, but if you don't quite understand this subject,
I would strongly recommend to checkout `http://content-security-policy.com/` in order for you to quickly resolve future security issue
that are detected by your own code or 3rd party libraries.

##- Routes Configuration

some people redirect all requests that can't match to an existing files to index.html,
I don't quite like this approach cause sometimes I would add link to invalid images, html or javascript locations
and they would return the index.html output instead of 404. it made it harder for me to notice errors, so instead
of that I have all the relevant routes listed in config.ini instead.

##- SSL Certificates

I use GoDaddy SSL Certificate, that comes with a primary crt file, the server key file, and the bundle file sf_bundle-g2-g1.

the docs says the following: `If the certificate is signed by a certificate authority, the certFile should be the concatenation of the server's certificate, any intermediates, and the CA's certificate.`

so in my case all i needed to do is to append the bundle file to the main crt file.

I checked the grade of the SSL Certificate using `https://www.ssllabs.com/ssltest/` and I got grade A.

I did leave the directive `OtherCertificates` under the SslCert category of config.ini, just in case for some reason it needed to be loaded separately.

# Tested

this package was tested on a macbook pro  with macOS Sierra, using go 1.7.1 installed with homebrew.

if for some reason you encounter problems with different versions of go, please let me know.
