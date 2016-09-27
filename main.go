package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	"github.com/codegangsta/cli"
)

// http handler
func serves(server *Server, headers []Header) {
	fs := http.FileServer(http.Dir(server.GetPath()))

	handlers := func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// print access log
			var accessLog = fmt.Sprintf("[%s] %s %s", r.RemoteAddr, r.Method, r.URL.String())
			log.Println(accessLog)

			// Set custom header
			for _, header := range headers {
				w.Header().Add(header.key, header.value) 
			}

			w.Header().Add("X-Server", "serves")
			h.ServeHTTP(w, r) 
		} 
	} 

	var realhost = fmt.Sprintf("%s:%d", server.GetHostname(), server.GetPort())
	s := http.Server{
		Addr: realhost,
		Handler: handlers(fs),
	}

	log.Printf("Listening... %s at %s\n", realhost, server.GetPath())

	panic(s.ListenAndServe())
}

func main() {
  app := cli.NewApp()
  app.Name = "serves"
  app.Usage = "serves your static files."
  app.Flags = []cli.Flag{
	  cli.IntFlag {
		  Name: "port,p",
		  Value: 5555,
		  Usage: "Port to be used (default 5555)",
	  },
	  cli.StringFlag {
		  Name: "host,n",
		  Value: "127.0.0.1",
		  Usage: "Host to be used (default: 127.0.0.1)",
	  },
	  cli.StringFlag {
		  Name: "root,r",
		  Value: "public",
		  Usage: "Host to be used (default: 127.0.0.1)",
	  },
	  cli.StringFlag {
		  Name: "header,H",
		  Value: "",
		  Usage: "headers file",
	  },
  }

  app.Action = func(c *cli.Context) {
	  var host = c.String("host")
	  var filename = c.String("header")
	  var port = c.Int("port")
	  var path = c.String("root")

	  // get directory to serve and check
	  if(len(c.Args()) > 0){
		  path = c.Args()[0]
	  }
	  if ! existsDir(path) {
		  panic("Path not exist: \"" + path + "\"")
	  }

	  var server = new(Server)
	  server.SetHostname(host)
	  server.SetPort(port)
	  server.SetPath(path)

	  var headers []Header
	  if len(filename) > 0 {
		  headers = parseHeader(filename)
	  }

	  serves(server, headers)
  }

  app.Run(os.Args)
}
