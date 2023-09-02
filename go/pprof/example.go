package main

import (
	"io"
	"log"
	"net/http"
	"net/http/pprof"
	"sync"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))

	// A sampling of all past memory allocations
	mux.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))

	// Stack traces that led to blocking on synchronization primitives
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))

	// The command line invocation of the current program
	mux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))

	// Stack traces of all current goroutines.
	// Use debug=2 as a query parameter to export in the same format as an unrecovered panic
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))

	// A sampling of memory allocations of live objects.
	// You can specify the gc GET parameter to run GC before taking the heap sample.
	mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))

	// Stack traces of holders of contended mutexes
	mux.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))

	// CPU profile. You can specify the duration in the seconds GET parameter.
	// After you get the profile file, use the go tool pprof command to investigate the profile.
	mux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))

	// Stack traces that led to the creation of new OS threads
	mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

	// A trace of execution of the current program. You can specify the duration in the seconds GET parameter.
	// After you get the trace file, use the go tool trace command to investigate the trace
	mux.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(j int) {
				j++
				j *= j
				j--
				wg.Done()
			}(i)
		}
		wg.Wait()
		_, _ = io.WriteString(w, http.StatusText(http.StatusOK))
	})

	log.Println("HTTP server started on localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal(err)
	}
}
