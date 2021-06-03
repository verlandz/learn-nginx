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
	home_hit         = 0
	name_hit         = 0
	cache_hit        = 0
	client_cache_hit = 0
	rate_limit_hit   = 0
	test_a, test_b   = 0, 0
	http_port        = ""
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
	test_a++
	s := fmt.Sprintf("TestA HIT - %v", test_a)
	log.Println(s)
	fmt.Fprintf(w, s)
}

func TestB(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	test_b++
	s := fmt.Sprintf("TestB HIT - %v", test_b)
	log.Println(s)
	fmt.Fprintf(w, s)
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

	log.Println("Listening to", http_port)
	log.Fatal(http.ListenAndServe(http_port, router))
}
