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
	"flag"
	"github.com/netsec-ethz/scion-apps/pkg/appnet/appquic"
	"log"

	"github.com/netsec-ethz/scion-apps/ftpd/internal/core"
	driver "github.com/netsec-ethz/scion-apps/ftpd/internal/driver/file"
)

func main() {
	var (
		root     = flag.String("root", "", "Root directory to serve")
		user     = flag.String("user", "", "Username for login (omit for public access)")
		pass     = flag.String("pass", "", "Password for login (omit for public access)")
		port     = flag.Uint("port", 2121, "Port")
		host     = flag.String("host", "", "Host (e.g. 1-ff00:0:110,[127.0.0.1])")
		hercules = flag.String("hercules", "", "Enable RETR_HERCULES using the Hercules binary specified\nIn Hercules mode, scionFTP checks the following directories for Hercules config files: ., /etc, /etc/scion-ftp")
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
		Perm:     core.NewSimplePerm("user", "group"),
	}

	certs := appquic.GetDummyTLSCerts()

	log.Printf("Starting FTP server on %v:%v", *host, *port)
	var auth core.Auth
	if *user == "" && *pass == "" {
		log.Printf("Anonymous FTP")
		auth = &core.AnonymousAuth{}
	} else {
		log.Printf("Username %v, Password %v", *user, *pass)
		auth = &core.SimpleAuth{Name: *user, Password: *pass}
	}
	opts := &core.Opts{
		Factory:        factory,
		Port:           uint16(*port),
		Hostname:       *host,
		Auth:           auth,
		Certificate:    &certs[0],
		HerculesBinary: *hercules,
		RootPath:       *root,
	}

	srv := core.NewServer(opts)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}