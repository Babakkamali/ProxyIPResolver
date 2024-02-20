package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"strings"
)

type headerCheck struct {
    Name  string
    CheckIP bool // true to check for real ip address
}

var headersToCheck = []headerCheck{
    {"CF-Connecting-IP", true},				// Cloudflare CDN
    {"True-Client-IP", true},				// Akamai and some others
    {"Ar-Real-Ip", true},					// ArvanCloud CDN
    {"X-Real-IP", true},
    {"X-Forwarded-For", true},
    {"Forwarded-For", true},
    {"X-Client-Ip", false},
    {"Forwarded", false},
    {"X-Forwarded", false},		
    {"X-Forwarded-Port", false},
    {"X-Forwarded-Proto", false},
    {"X-Forwarded-Server", false},
    {"X-Forwarded-Host", false},
    {"X-Country-Code", false},
    {"Ar-Real-Country", false},				// ArvanCloud CDN
    {"Upgrade-Insecure-Requests", false},
    {"User-Agent", false},
    {"Cdn-Loop", false},
    {"Accept", false},
    {"Accept-Encoding", false},
}

// getIP extracts the client's IP address from the request headers or RemoteAddr.
func getIP(r *http.Request) string {
	for _, header := range headersToCheck {
		if header.CheckIP {
			value := r.Header.Get(header.Name)
			if value != "" {
				ips := strings.Split(value, ",")
				trimmedIp := strings.TrimSpace(ips[0])
				if trimmedIp != "" {
					return trimmedIp
				}
			}
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr // Return the full remote address if splitting fails
	}
	return ip
}

func ipHandler(w http.ResponseWriter, r *http.Request, showAll bool) {
	clientIp := getIP(r)
	var response strings.Builder

	// Display the client IP address at the beginning of the response
	response.WriteString(fmt.Sprintf("\n\nYour IP address is %s\n\n\n\n", clientIp))

	// Iterate through the predefined headers and format the output
	for _, header := range headersToCheck {
		values := r.Header.Values(header.Name)
		if len(values) == 0 {
		    // don't print for now
		    //response.WriteString(fmt.Sprintf("%s: null\n\n", header.Name)) 
		} else {
			for _, value := range values {
				response.WriteString(fmt.Sprintf("%s: %s\n\n", header.Name, value))
			}
		}
	}
	

	// Include all headers if the flag is set
	if showAll {
		response.WriteString("\n\nAll Headers:\n")
		for name, values := range r.Header {
			for _, value := range values {
				response.WriteString(fmt.Sprintf("%s: %s\n", name, value))
			}
		}
	}

	fmt.Fprint(w, response.String())
}

func main() {

	showAllHeaders := flag.Bool("show-all-headers", false, "Show all HTTP headers")
	flag.BoolVar(showAllHeaders, "all", false, "Show all HTTP headers (shorthand)")

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	ipHandler(w, r, *showAllHeaders)
	})

	fmt.Println("App listening on port 3000!")
	if err := http.ListenAndServe("0.0.0.0:3000", nil); err != nil {
		panic(err)
	}
}

