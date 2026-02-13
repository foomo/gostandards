# HTTP

The `http` package provides typed constants for standard HTTP headers, content encodings, and authentication prefixes.

## Types

### Header

Standard HTTP header name.

```go
type Header string
```

**Methods:**

| Method | Description |
| --- | --- |
| `String() string` | Returns the string representation |
| `Add(header http.Header, value string)` | Adds a value to the given `http.Header` |
| `Get(header http.Header) string` | Gets the first value from the given `http.Header` |

**Usage:**

```go
package main

import (
	"net/http"

	gohttp "github.com/foomo/gostandards/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Read a header from the request
	contentType := gohttp.HeaderContentType.Get(r.Header)

	// Set a header on the response
	gohttp.HeaderContentType.Add(w.Header(), "application/json")
}
```

**Common constants:**

```go
gohttp.HeaderContentType              // "Content-Type"
gohttp.HeaderAuthorization            // "Authorization"
gohttp.HeaderAccept                   // "Accept"
gohttp.HeaderCacheControl             // "Cache-Control"
gohttp.HeaderAccessControlAllowOrigin // "Access-Control-Allow-Origin"
gohttp.HeaderXRequestID               // "X-Request-Id"
```

The full list contains 134 header constants covering standard, CORS, security, caching, Cloudflare, and tracing headers.

---

### Encoding

Data encoding format for content negotiation.

```go
type Encoding string
```

**Methods:**

| Method | Description |
| --- | --- |
| `String() string` | Returns the string representation |

**Constants:**

```go
gohttp.EncodingGzip     // "gzip"
gohttp.EncodingDeflate  // "deflate"
gohttp.EncodingBr       // "br" (Brotli)
gohttp.EncodingIdentity // "identity"
gohttp.EncodingCompress // "compress"
```

**Usage:**

```go
gohttp.HeaderAcceptEncoding.Add(req.Header, gohttp.EncodingGzip.String())
```

---

### AuthPrefix

Authentication scheme prefix for the `Authorization` header.

```go
type AuthPrefix string
```

**Methods:**

| Method | Description |
| --- | --- |
| `String() string` | Returns the string representation |
| `Set(value string) string` | Returns the prefix combined with the given value (e.g., `"Bearer mytoken"`) |

**Constants:**

```go
gohttp.AuthPrefixBearer    // "Bearer"
gohttp.AuthPrefixBasic     // "Basic"
gohttp.AuthPrefixDigest    // "Digest"
gohttp.AuthPrefixNegotiate // "Negotiate"
gohttp.AuthPrefixAWS4HMAC  // "AWS4-HMAC-SHA256"
gohttp.AuthPrefixAPIKey    // "Api-Key"
```

**Usage:**

```go
// Set a Bearer token on a request
token := "eyJhbGciOi..."
gohttp.HeaderAuthorization.Add(
	req.Header,
	gohttp.AuthPrefixBearer.Set(token),
)
// Adds: "Authorization: Bearer eyJhbGciOi..."
```
