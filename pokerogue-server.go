package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/Flashfyre/pokerogue-server/api"
	"github.com/Flashfyre/pokerogue-server/db"
)

func main() {
	network := flag.String("network", "tcp", "network type for server to listen on (tcp, unix)")
	address := flag.String("address", "0.0.0.0", "network address for server to listen on")

	dbuser := flag.String("dbuser", "pokerogue", "database username")
	dbpass := flag.String("dbpass", "", "database password")
	dbproto := flag.String("dbproto", "tcp", "protocol for database connection")
	dbaddr := flag.String("dbaddr", "127.0.0.1", "database address")
	dbname := flag.String("dbname", "pokerogue", "database name")

	flag.Parse()

	err := db.Init(*dbuser, *dbpass, *dbproto, *dbaddr, *dbname)
	if err != nil {
		log.Fatalf("failed to initialize database: %s", err)
	}

	listener, err := net.Listen(*network, *address)
	if err != nil {
		log.Fatalf("failed to create net listener: %s", err)
	}

	// account
	http.HandleFunc("/api/account/info", api.HandleAccountInfo)
	http.HandleFunc("/api/account/register", api.HandleAccountRegister)
	http.HandleFunc("/api/account/login", api.HandleAccountLogin)
	http.HandleFunc("/api/account/logout", api.HandleAccountLogout)

	// savedata
	http.HandleFunc("/api/savedata/get", api.HandleSavedataGet)
	http.HandleFunc("/api/savedata/update", api.HandleSavedataUpdate)
	http.HandleFunc("/api/savedata/delete", api.HandleSavedataDelete)
	
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatalf("failed to create http server or server errored: %s", err)
	}
}
