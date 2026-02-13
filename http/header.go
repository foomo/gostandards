package http

import (
	"net/http"
)

// Header represents a standard HTTP header.
type Header string

const (
	HeaderAIM                             Header = "A-IM"
	HeaderAccept                          Header = "Accept"
	HeaderAcceptCharset                   Header = "Accept-Charset"
	HeaderAcceptDatetime                  Header = "Accept-Datetime"
	HeaderAcceptEncoding                  Header = "Accept-Encoding"
	HeaderAcceptLanguage                  Header = "Accept-Language"
	HeaderAcceptPatch                     Header = "Accept-Patch"
	HeaderAcceptRanges                    Header = "Accept-Ranges"
	HeaderAccessControlAllowCredentials   Header = "Access-Control-Allow-Credentials"
	HeaderAccessControlAllowHeaders       Header = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowMethods       Header = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowOrigin        Header = "Access-Control-Allow-Origin"
	HeaderAccessControlExposeHeaders      Header = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge             Header = "Access-Control-Max-Age"
	HeaderAccessControlRequestHeaders     Header = "Access-Control-Request-Headers"
	HeaderAccessControlRequestMethod      Header = "Access-Control-Request-Method"
	HeaderAge                             Header = "Age"
	HeaderAllow                           Header = "Allow"
	HeaderAltSvc                          Header = "Alt-Svc"
	HeaderAuthorization                   Header = "Authorization"
	HeaderBaggage                         Header = "Baggage"          // W3C Baggage header for passing contextual information
	HeaderCFAuthorization                 Header = "CF-Authorization" // Used for Cloudflare Access (for securing endpoints)
	HeaderCFBypass                        Header = "CF-Bypass"        // Bypass rules for Cloudflare Access or other services
	HeaderCFCacheStatus                   Header = "CF-Cache-Status"  // Cache status (HIT, MISS, EXPIRED, etc.)
	HeaderCFConnectingIP                  Header = "CF-Connecting-IP" // Original client IP address as seen by Cloudflare
	HeaderCFIPCountry                     Header = "CF-IPCountry"     // Country of the requester's IP address
	HeaderCFRailgun                       Header = "CF-Railgun"       // Railgun performance information (used in Cloudflare Railgun optimization)
	HeaderCFRay                           Header = "CF-RAY"           // Unique request ID (used for tracing through Cloudflare's network)
	HeaderCFRequestID                     Header = "CF-Request-ID"    // Unique ID for the request for tracking purposes
	HeaderCFRocketLoader                  Header = "CF-Rocket-Loader" // Related to Cloudflare's Rocket Loader optimization
	HeaderCFVisitor                       Header = "CF-Visitor"       // Information about the protocol (HTTP/HTTPS) used by the visitor
	HeaderCFWorker                        Header = "CF-Worker"        // Header set by Cloudflare Workers
	HeaderCacheControl                    Header = "Cache-Control"
	HeaderConnection                      Header = "Connection"
	HeaderContentDisposition              Header = "Content-Disposition"
	HeaderContentEncoding                 Header = "Content-Encoding"
	HeaderContentLanguage                 Header = "Content-Language"
	HeaderContentLength                   Header = "Content-Length"
	HeaderContentLocation                 Header = "Content-Location"
	HeaderContentMD5                      Header = "Content-MD5"
	HeaderContentRange                    Header = "Content-Range"
	HeaderContentSecurityPolicy           Header = "Content-Security-Policy"
	HeaderContentSecurityPolicyReportOnly Header = "Content-Security-Policy-Report-Only"
	HeaderContentType                     Header = "Content-Type"
	HeaderCookie                          Header = "Cookie"
	HeaderDAV                             Header = "DAV"
	HeaderDate                            Header = "Date"
	HeaderDepth                           Header = "Depth"
	HeaderDestination                     Header = "Destination"
	HeaderETag                            Header = "ETag"
	HeaderExpect                          Header = "Expect"
	HeaderExpires                         Header = "Expires"
	HeaderForwarded                       Header = "Forwarded"
	HeaderFrom                            Header = "From"
	HeaderHost                            Header = "Host"
	HeaderIf                              Header = "If"
	HeaderIfMatch                         Header = "If-Match"
	HeaderIfModifiedSince                 Header = "If-Modified-Since"
	HeaderIfNoneMatch                     Header = "If-None-Match"
	HeaderIfRange                         Header = "If-Range"
	HeaderIfUnmodifiedSince               Header = "If-Unmodified-Since"
	HeaderLastModified                    Header = "Last-Modified"
	HeaderLink                            Header = "Link"
	HeaderLocation                        Header = "Location"
	HeaderLockToken                       Header = "Lock-Token"
	HeaderMaxForwards                     Header = "Max-Forwards"
	HeaderOrigin                          Header = "Origin"
	HeaderP3P                             Header = "P3P"
	HeaderPragma                          Header = "Pragma"
	HeaderProxyAuthenticate               Header = "Proxy-Authenticate"
	HeaderProxyAuthorization              Header = "Proxy-Authorization"
	HeaderPublicKeyPins                   Header = "Public-Key-Pins"
	HeaderRange                           Header = "Range"
	HeaderReferer                         Header = "Referer"
	HeaderReferrerPolicy                  Header = "Referrer-Policy"
	HeaderRetryAfter                      Header = "Retry-After"
	HeaderServer                          Header = "Server"
	HeaderSetCookie                       Header = "Set-Cookie"
	HeaderStrictTransportSecurity         Header = "Strict-Transport-Security"
	HeaderTE                              Header = "TE"
	HeaderTimeout                         Header = "Timeout"
	HeaderTk                              Header = "Tk"
	HeaderTraceParent                     Header = "traceparent" // W3C Trace Context header for trace propagation
	HeaderTraceState                      Header = "tracestate"  // W3C Trace Context header for vendor-specific trace information
	HeaderTrailer                         Header = "Trailer"
	HeaderTransferEncoding                Header = "Transfer-Encoding"
	HeaderUpgrade                         Header = "Upgrade"
	HeaderUpgradeInsecureRequests         Header = "Upgrade-Insecure-Requests"
	HeaderUserAgent                       Header = "User-Agent"
	HeaderVary                            Header = "Vary"
	HeaderVia                             Header = "Via"
	HeaderWWWAuthenticate                 Header = "WWW-Authenticate"
	HeaderWarning                         Header = "Warning"
	HeaderXAccelRedirect                  Header = "X-Accel-Redirect"  // Internal redirect by reverse proxy (e.g., Nginx)
	HeaderXAmznTraceID                    Header = "X-Amzn-Trace-Id"   // AWS X-Ray trace header
	HeaderXB3Flags                        Header = "X-B3-Flags"        // B3 Debug flag
	HeaderXB3ParentSpanID                 Header = "X-B3-ParentSpanId" // B3 Parent Span ID used for distributed tracing
	HeaderXB3Sampled                      Header = "X-B3-Sampled"      // B3 Sampling flag (whether or not the request is sampled)
	HeaderXB3SpanID                       Header = "X-B3-SpanId"       // B3 Span ID used for distributed tracing
	HeaderXB3TraceID                      Header = "X-B3-TraceId"      // B3 Trace ID used for distributed tracing
	HeaderXCSRFToken                      Header = "X-CSRF-Token"
	HeaderXCache                          Header = "X-Cache"               // Cache status of the request (e.g., HIT or MISS)
	HeaderXCloudTraceContext              Header = "X-Cloud-Trace-Context" // Google Cloud Trace context header
	HeaderXContentDuration                Header = "X-Content-Duration"    // Duration of the content (e.g., media file)
	HeaderXContentTypeOptions             Header = "X-Content-Type-Options"
	HeaderXCorrelationID                  Header = "X-Correlation-ID"       // Correlation ID for tracing requests
	HeaderXDNSPrefetchControl             Header = "X-DNS-Prefetch-Control" //
	HeaderXDownloadOptions                Header = "X-Download-Options"     // Prevents file download behavior in IE
	HeaderXForwardedFor                   Header = "X-Forwarded-For"        // Original client IP behind proxy
	HeaderXForwardedHost                  Header = "X-Forwarded-Host"       // Original host requested by the client
	HeaderXForwardedProto                 Header = "X-Forwarded-Proto"      // Original protocol (HTTP or HTTPS) used by the client
	HeaderXFrameOptions                   Header = "X-Frame-Options"
	HeaderXOTSpanContext                  Header = "X-OT-Span-Context"                 // OpenTracing Span Context (legacy, pre-OpenTelemetry)
	HeaderXPermittedCrossDomainPolicies   Header = "X-Permitted-Cross-Domain-Policies" // Adobe Flash Player policy control
	HeaderXPoweredBy                      Header = "X-Powered-By"                      //
	HeaderXRateLimitLimit                 Header = "X-RateLimit-Limit"                 // Limit for the number of requests in a given time
	HeaderXRateLimitRemaining             Header = "X-RateLimit-Remaining"             // Number of remaining requests in the current rate limit window
	HeaderXRateLimitReset                 Header = "X-RateLimit-Reset"                 // Time at which the rate limit resets
	HeaderXRealIP                         Header = "X-Real-IP"                         // The original client IP address
	HeaderXRequestID                      Header = "X-Request-ID"                      // Unique request ID for tracing
	HeaderXRobotsTag                      Header = "X-Robots-Tag"                      // Search engine control
	HeaderXUACompatible                   Header = "X-UA-Compatible"                   // Compatibility mode for Internet Explorer
	HeaderXXSSProtection                  Header = "X-XSS-Protection"
)

// String returns the string representation of the Header type.
func (h Header) String() string {
	return string(h)
}

// Add adds a value to the specified header using the Header's string representation as the key.
func (h Header) Add(header http.Header, value string) {
	header.Add(h.String(), value)
}

// Get returns the value of the specified header using the Header's string representation as the key.'
func (h Header) Get(header http.Header) string {
	return header.Get(h.String())
}
