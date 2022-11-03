package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Define a new command-line flag with the name 'addr', a default value
	// of ":4000" and some short help text explaining what the flag controls.
	// The value of the flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Use log.New() to create a logger for writting information messages. This
	// takes three parameters: the destination to write the logs to (os.Stdout),
	// a string prefix for message (INFO followed by a tab), and flags to indicate
	// what additional information to include (local date and time). Note that the
	// flags are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messages in the same way, but use stderr as
	// the destination and use log.Lshortfile flag to include the relevant file name
	// and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static"
	// directory. Note that the path given to the http.Dir function is
	// relative to the project directory root.
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Use the mux.Handle() method to register the file server as the
	// handler for all URL paths that start with "/static/". For matching
	// paths, we strip the "/static" prefix before the request reaches the
	// file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Register application routes.
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Initialize a new http.Server struct. We set the Addr and Handler fields so
	// that the server uses the same network address and routes as before, and set
	// the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
