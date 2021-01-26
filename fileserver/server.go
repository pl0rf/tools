package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var dir string
var addr string

func isDir(p string) bool {
	if stat, err := os.Stat(p); err == nil && stat.IsDir() {
		return true
	}
	return false
}

func logHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()

	fullpath, err := filepath.Abs(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !isDir(fullpath) {
		fmt.Printf("ERR: \"%s\" is not a valid directory\n", fullpath)
		return
	}

	fmt.Printf("Serving %s on %s\n", fullpath, addr)
	log.Fatal(http.ListenAndServe(addr, logHandler(http.FileServer(http.Dir(dir)))))
}

func init() {
	const (
		defaultDir  = "."
		defaultPort = ":80"
		dirUsage    = "define what directory to serve"
		portUsage   = "define what TCP port to bind to e.g. localhost:8080"
	)
	flag.StringVar(&addr, "port", defaultPort, portUsage)
	flag.StringVar(&addr, "p", defaultPort, portUsage+" (shorthand)")
	flag.StringVar(&dir, "dir", defaultDir, dirUsage)
	flag.StringVar(&dir, "d", defaultDir, dirUsage+" (shorthand)")
}
