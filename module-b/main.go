package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/peterbourgon/ff/v3"
)

func main() {
	fs := flag.NewFlagSet("example", flag.ExitOnError)
	var s string
	fs.StringVar(&s, "s", "default", "a string flag")

	ff.Parse(fs, []string{"-s", "value"}, ff.WithConfigFileFlag("config"), ff.WithEnvVarNoPrefix())
	fmt.Println("Value of s:", s)

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id := uuid.NewString()
		buildInfo, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Fprintf(w, "Failed to read build info\n")
			return
		}
		fmt.Fprintf(w, "UUID from module-b: %s\n", id)
		for _, dep := range buildInfo.Deps {
			fmt.Fprintf(w, "%s: %s\n", dep.Path, dep.Version)
		}
	})
	http.Handle("/", r)
	http.ListenAndServe(":8081", nil)
}
