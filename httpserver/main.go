package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	port := 8080
	portstr := ":" + strconv.Itoa(port)
	log.Printf("httpserver listend port: %+v", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", HttpHandler)
	mux.HandleFunc("/healthz", HealthZ)
	mux.HandleFunc("/sleep", SleepTest)
	server := &http.Server{
		Addr:         portstr,
		Handler:      mux,
	}
	go server.ListenAndServe()
	listenSignal(context.Background(), server)
}

func HealthZ(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
func HttpHandler(w http.ResponseWriter, r *http.Request) {
	ip := ClientIP(r)
	httpcode := 200
	log.Printf("reqLog: clientIP : %+v, httpcode : %+v", ip, httpcode)
	headers := r.Header
	log.Printf("req headers : %+v", headers)
	for k, v := range headers {
		for _, val := range v {
			w.Header().Set(k, val)
		}
	}
	version := os.Getenv("VERSION")
	if version != "" {
		w.Header().Set("VERSION", version)
	}
	w.WriteHeader(httpcode)
	io.WriteString(w, "ok")
}

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func SleepTest(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	sleeptimeStr := values.Get("sleep")
	sleeptime, err := strconv.Atoi(sleeptimeStr)
	if err != nil {
		sleeptime = 1
	}
	time.Sleep(time.Duration(sleeptime) * time.Second)
	fmt.Fprintln(w, "Hello world, sleep " + strconv.Itoa(sleeptime) + "s")
	log.Printf( "Hello world, sleep %+vs", sleeptime)
}

func listenSignal(ctx context.Context, httpSrv *http.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-sigs:
		fmt.Println("notify sigs")
		httpSrv.Shutdown(ctx)
		fmt.Println("http shutdown gracefully")
	}
}
