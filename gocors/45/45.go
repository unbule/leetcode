package main

import (
	"fmt"
	"net/http"
)

// type HandlerFunc func(ResponseWriter, *Request)

// type ServeMux struct {
// 	mu    sync.RWMutex
// 	m     map[string]muxEntry
// 	es    []muxEntry // slice of entries sorted from longest to shortest.
// 	hosts bool       // whether any patterns contain hostnames
// }

// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
// 	if handler == nil {
// 		panic("http: nil handler")
// 	}
// 	mux.Handle(pattern, HandlerFunc(handler))
// }

// func (mux *ServeMux) Handle(pattern string, handler Handler) {
// 	mux.mu.Lock()
// 	defer mux.mu.Unlock()

// 	if pattern == "" {
// 		panic("http: invalid pattern")
// 	}
// 	if handler == nil {
// 		panic("http: nil handler")
// 	}
// 	if _, exist := mux.m[pattern]; exist {
// 		panic("http: multiple registrations for " + pattern)
// 	}

// 	if mux.m == nil {
// 		mux.m = make(map[string]muxEntry)
// 	}
// 	e := muxEntry{h: handler, pattern: pattern}
// 	mux.m[pattern] = e
// 	if pattern[len(pattern)-1] == '/' {
// 		mux.es = appendSorted(mux.es, e)
// 	}

// 	if pattern[0] != '/' {
// 		mux.hosts = true
// 	}
// }

// func appendSorted(es []muxEntry, e muxEntry) []muxEntry {
// 	n := len(es)
// 	i := sort.Search(n, func(i int) bool {
// 		return len(es[i].pattern) < len(e.pattern)
// 	})
// 	if i == n {
// 		return append(es, e)
// 	}
// 	// we now know that i points at where we want to insert
// 	es = append(es, muxEntry{}) // try to grow the slice in place, any entry works.
// 	copy(es[i+1:], es[i:])      // Move shorter entries down
// 	es[i] = e
// 	return es
// }

func rootGateWay(w http.ResponseWriter, r *http.Request) {
	println("Welcome to Chris's homepage! ")
}

func defaultGateWay(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello jingjing")
	println("Welcome to Chris's homepage! ")
}

func main() {
	http.HandleFunc("/", defaultGateWay)
	http.ListenAndServe(":8080", nil)
}
