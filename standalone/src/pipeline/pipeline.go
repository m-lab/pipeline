package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"
	"validators"
)

type Chunk struct {
	Tool string
	Version string
	Data []byte
}

var funcs = map[string]interface{} {
	"fathom": validators.Fathom,
	"ndt": validators.Ndt,
}

func (c *Chunk) validate() error {
	fun := funcs[c.Tool]
	if fun == nil {
		return errors.New("No validator for tool")
	}
	f := reflect.ValueOf(fun)
	in := make([]reflect.Value, 2)
	in[0] = reflect.ValueOf(c.Version)
	in[1] = reflect.ValueOf(c.Data)
	valid := f.Call(in)
	if len(valid) != 1 {
		return errors.New("Unexpected return from validation function")
	}
	if !valid[0].IsNil() {
		return valid[0].Interface().(error)
	}
	return nil
}

func (c *Chunk) save() error {
	// Use tool/version to lookup data format restrictions and apply them.
	// Eg: XML, JSON, content length, etc.
	err := c.validate()
	if err != nil {
		return err
	}

	filename := c.Tool + "-" + c.Version + "." +
		strconv.FormatInt(time.Now().Unix(), 10)
	return ioutil.WriteFile(filename, c.Data, 0600)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", root).Methods("GET")
	r.HandleFunc("/{tool}", tool).Methods("GET")
	r.HandleFunc("/{tool}/{version}", tool_and_version)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	// TODO(dominic): Return information on pipeline and links to tool
	// 		  metrics.
	fmt.Fprint(w, "Info on pipeline and links to tools will be here")
}

func tool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tool := vars["tool"]

	// TODO(dominic): return metrics for tool
	fmt.Fprintf(w, "Metrics for tool '%s' will be here", tool)
}

func tool_and_version(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tool := vars["tool"]
	version := vars["version"]

	if r.Method == "GET" {
		// TODO(dominic): return metrics for tool and version
		fmt.Fprintf(w, "Metrics for tool '%s-%s' will be here",
			    tool, version)
	} else if r.Method == "POST" {
		fmt.Fprintf(w, "Data POSTED for tool '%s-%s'", tool, version)
		// Read body
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		// And save to disk
		chunk := &Chunk{Tool: tool, Version: version, Data: buf.Bytes()}
		err := chunk.save()
		if err != nil {
			http.Error(w, err.Error(),
				   http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusFound)
	}
}
