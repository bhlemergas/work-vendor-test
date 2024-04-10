package main

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"flag"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	ff "github.com/peterbourgon/ff"
)

func main() {
	fs := flag.NewFlagSet("example", flag.ExitOnError)
	var s string
	fs.StringVar(&s, "s", "default", "a string flag")

	ff.Parse(fs, []string{"-s", "value"}, ff.WithEnvVarNoPrefix())
	fmt.Println("Value of s:", s)

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id := uuid.NewString()
		buildInfo, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Fprintf(w, "Failed to read build info\n")
			return
		}
		fmt.Fprintf(w, "UUID from module-a: %s\n", id)
		for _, dep := range buildInfo.Deps {
			fmt.Fprintf(w, "%s: %s\n", dep.Path, dep.Version)
		}
	})
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
