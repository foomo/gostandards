package http

type Encoding string

const (
	EncodingGzip     Encoding = "gzip"     // Gzip compression
	EncodingDeflate  Encoding = "deflate"  // Deflate compression
	EncodingBr       Encoding = "br"       // Brotli compression
	EncodingIdentity Encoding = "identity" // No transformation or compression
	EncodingCompress Encoding = "compress" // Deprecated Unix 'compress' algorithm
)

func (e Encoding) String() string {
	return string(e)
}
