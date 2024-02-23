# ProxyIPResolver

ProxyIPResolver is a Go application designed to accurately identify the real IP address of a client connecting to your server, even when requests are proxied through various layers like load balancers, CDNs (Content Delivery Networks), and reverse proxies. It extracts IP information from standard and non-standard HTTP headers and provides a comprehensive view of all request headers when needed.

## Features

- Extracts the client's real IP address from common and custom headers used by proxies and CDNs.
- Supports a wide range of headers, including `X-Forwarded-For`, `X-Real-IP`, `CF-Connecting-IP` (Cloudflare), and more.
- Option to display all HTTP request headers for debugging or informational purposes.
- Simple, lightweight, and easy to integrate into existing Go web applications.

## Getting Started

### Prerequisites

- Go (version 1.15 or later recommended)

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/Babakkamali/ProxyIPResolver.git
cd ProxyIPResolver
```

Build the application:
 
```bash
go build -o proxyipresolver
```

### Usage

Run the application:

```bash
./proxyipresolver
```

By default, the application listens on port 3000 and can be accessed via `http://localhost:3000`. 

To display all HTTP request headers, start the application with the `--show-all-headers` or `-all` flag:

```bash
./proxyipresolver --show-all-headers
# or
./proxyipresolver -all
```

### Docker Support

You can also run ProxyIPResolver within a Docker container. Make sure you have docker and docker compose plugin installed And then build and run the Docker container:

```bash
docker compose up -d
```

## Docker All Header Usage

when using Docker or Docker Compose, you can specify the environment variable in your docker-compose.yml file:

```bash
services:
  appserver:
    environment:
      - SHOW_ALL_HEADERS=true

```

set the SHOW_ALL_HEADERS=true to receive all headers

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests with improvements or additional features.

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Your Name - [info@babakkamali.com](mailto:info@babakkamali.com)

Project Link: [https://github.com/Babakkamali/ProxyIPResolver](https://github.com/Babakkamali/ProxyIPResolver)

## Demo Link

Demo Link: [https://ipchecker.babakkamali.com](https://ipchecker.babakkamali.com)