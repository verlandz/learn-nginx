// cmd: go run main.go -port 8080
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	home_hit               = 0
	name_hit               = 0
	cache_hit              = 0
	client_cache_hit       = 0
	rate_limit_hit         = 0
	test_a_hit, test_b_hit = 0, 0
	bad_request_hit        = 0
	http_port              = ""
)

func init() {
	// using 8080 as default
	http_port_flag := flag.String("port", "8080", "custom port")
	flag.Parse()

	http_port = ":" + string(*http_port_flag)
}

// HTTP handlers
func Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	home_hit++
	s := fmt.Sprintf("Home HIT - %v", home_hit)
	log.Println(s)
	fmt.Fprintf(w, s)
}

func Name(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name_hit++
	s := fmt.Sprintf("Name HIT - %v | name:%v | [query-string] a:%v & b:%v",
		name_hit,
		ps.ByName("name"),
		r.FormValue("a"), r.FormValue("b"),
	)
	log.Println(s)
	fmt.Fprintf(w, s)
}

func Cache(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cache_hit++
	s := fmt.Sprintf("Cache HIT - %v | Hello:%v", cache_hit, r.Header.Get("Hello"))
	log.Println(s)
	fmt.Fprintf(w, s)
}

func ClientCache(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client_cache_hit++
	s := fmt.Sprintf("ClientCache HIT - %v", client_cache_hit)
	log.Println(s)
	fmt.Fprintf(w, s)
}

func RateLimit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rate_limit_hit++
	s := fmt.Sprintf("RateLimit HIT - %v", rate_limit_hit)
	log.Println(s)
	fmt.Fprintf(w, s)
}

func TestA(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	test_a_hit++
	s := fmt.Sprintf("TestA HIT - %v", test_a_hit)
	log.Println(s)
	fmt.Fprintf(w, s)
}

func TestB(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	test_b_hit++
	s := fmt.Sprintf("TestB HIT - %v", test_b_hit)
	log.Println(s)
	fmt.Fprintf(w, s)
}

func BadRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bad_request_hit++
	s := fmt.Sprintf("BadRequest HIT - %v", bad_request_hit)
	log.Println(s)

	q := r.FormValue("q")
	userID := r.Header.Get("X-USER-ID")
	device := r.Header.Get("X-DEVICE")

	// from query (q)
	if q == "" {
		http.Error(w, "query can't be empty", http.StatusBadRequest)
		return
	}

	// from header (userID)
	if userID == "" {
		http.Error(w, "userID can't be empty", http.StatusBadRequest)
		return
	}

	// from header (device)
	if device == "" {
		http.Error(w, "device can't be empty", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
	return
}

// Main
func main() {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NotFound to %v \n\n- myapp\n", r.URL.Path)
	})

	router.GET("/", Home)
	router.GET("/name/:name", Name)
	router.GET("/cache", Cache)
	router.GET("/client-cache", ClientCache)
	router.GET("/rate-limit", RateLimit)
	router.GET("/test/a", TestA)
	router.GET("/test/b", TestB)
	router.GET("/bad-request", BadRequest)

	log.Println("Listening to", http_port)
	log.Fatal(http.ListenAndServe(http_port, router))
}
