package http

// Encoding represents a specific data encoding format, typically used for compression or content transformation.
type Encoding string

const (
	EncodingGzip     Encoding = "gzip"     // Gzip compression
	EncodingDeflate  Encoding = "deflate"  // Deflate compression
	EncodingBr       Encoding = "br"       // Brotli compression
	EncodingIdentity Encoding = "identity" // No transformation or compression
	EncodingCompress Encoding = "compress" // Deprecated Unix 'compress' algorithm
)

// String returns the string representation of the Encoding type.
func (e Encoding) String() string {
	return string(e)
}
