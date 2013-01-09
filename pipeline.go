package pipeline

import (
	"appengine"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
       )

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", root).Methods("GET")
	r.HandleFunc("/{tool}", tool).Methods("GET")
	r.HandleFunc("/{tool}/{version}", tool_and_version)
	http.Handle("/", r)
}

func root(w http.ResponseWriter, r *http.Request) {
	// TODO(dominic): Return information on pipeline and links to tool metrics.
	fmt.Fprint(w, "Info on pipeline and links to tools will be here")
}

func tool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tool := vars["tool"]

	// TODO(dominic): return metrics for tool
	c := appengine.NewContext(r)
	c.Infof("Metrics for tool '%s' requested", tool);
	fmt.Fprintf(w, "Metrics for tool '%s' will be here", tool)
}

func tool_and_version(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tool := vars["tool"]
	version := vars["version"]

	c := appengine.NewContext(r)
	if r.Method == "GET" {
		// TODO(dominic): return metrics for tool and version
		c.Infof("Metrics for tool '%s-%s' requested", tool, version);
		fmt.Fprintf(w, "Metrics for tool '%s-%s' will be here", tool, version)
	} else if r.Method == "POST" {
		// TODO(dominic): get data from request and save to datastore for processing
		c.Infof("Data POSTED for tool '%s-%s'", tool, version)
		fmt.Fprintf(w, "Data POSTED for tool '%s-%s'", tool, version)
		w.WriteHeader(http.StatusFound)
	}
}
