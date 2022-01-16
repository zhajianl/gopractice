package main
import (
        "fmt"
        "log"
        //"net"
        "net/http"
        //"net/http/pprof"
        "os"
        "strings"
)
func index(w http.ResponseWriter, r *http.Request) {
        version := os.Getenv("VERSION")
        w.Header().Set("VERSION", version)
        fmt.Printf("os version: %s \n", version)
        for k, v := range r.Header {
                for _, vv := range v {
                        fmt.Printf("Header key: %s, Header value: %s \n", k, v)
                        w.Header().Set(k, vv)
                }
        }
        clientip := getCurrentIP(r)
        log.Printf("Success! Response code: %d", 200)
        log.Printf("Success! clientip: %s", clientip)
}

func healthz(w http.ResponseWriter, r *http.Request) {

        fmt.Fprintf(w, "200")
}


func getCurrentIP( r *http.Request) string {
        ip := r.Header.Get("X-Real-IP")
        if ip == "" {
          ip = strings.Split(r.RemoteAddr,":")[0]
        }
        return ip

}


func main() {
        mux := http.NewServeMux()
        mux.HandleFunc("/", index)
        mux.HandleFunc("/healthz", healthz)
        if err := http.ListenAndServe(":8080", mux); err != nil {
                log.Fatalf("start http server failed, error: %s\n", err.Error())
        }
}
