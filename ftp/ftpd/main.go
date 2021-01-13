// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// Modifications 2019 Elwin Stephan to make it compatible to SCION and
// introduce parts of the GridFTP extension
//
// Copyright 2020-2021 ETH Zurich modifications to add support for SCION
package main

import (
	"crypto/tls"
	"flag"
	"log"

	driver "github.com/netsec-ethz/scion-apps/ftp/ftpd/internal/driver/file"
	"github.com/netsec-ethz/scion-apps/ftp/ftpd/internal/ftp"
)

func main() {
	var (
		root     = flag.String("root", "", "Root directory to serve")
		user     = flag.String("user", "admin", "Username for login")
		pass     = flag.String("pass", "123456", "Password for login")
		port     = flag.Uint("port", 2121, "Port")
		host     = flag.String("host", "", "Host (e.g. 1-ff00:0:110,[127.0.0.1])")
		certFile = flag.String("cert", "", "TLS certificate file")
		keyFile  = flag.String("key", "", "TLS private key file")
		hercules = flag.String("hercules", "", "Enable RETR_HERCULES using the Hercules binary specified")
	)
	flag.Parse()
	if *root == "" {
		log.Fatalf("Please set a root to serve with -root")
	}

	if *host == "" {
		log.Fatalf("Please set the hostaddress with -host")
	}

	factory := &driver.FileDriverFactory{
		RootPath: *root,
		Perm:     ftp.NewSimplePerm("user", "group"),
	}

	if *certFile == "" || *keyFile == "" {
		log.Fatalf("Please specify public/private key files to use with -cert and -key")
	}

	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatalf("could not load key-pair: %s", err.Error())
	}

	opts := &ftp.Opts{
		Factory:        factory,
		Port:           uint16(*port),
		Hostname:       *host,
		Auth:           &ftp.SimpleAuth{Name: *user, Password: *pass},
		Certificate:    &cert,
		HerculesBinary: *hercules,
		RootPath:       *root,
	}

	log.Printf("Starting ftp server on %v:%v", opts.Hostname, opts.Port)
	log.Printf("Username %v, Password %v", *user, *pass)
	srv := ftp.NewServer(opts)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
