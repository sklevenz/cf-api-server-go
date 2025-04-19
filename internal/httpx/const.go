package httpx

// HeaderContentType is the HTTP header key for specifying the media type of the request or response body.
const HeaderContentType = "Content-Type"

// HeaderContentLength is the HTTP header key for specifying the length of the request or response body in bytes.
const HeaderContentLength = "Content-Length"

// HeaderCacheControl represents the HTTP header field "Cache-Control",
// which is used to specify directives for caching mechanisms in both
// requests and responses.
const HeaderCacheControl = "Cache-Control"

// HeaderETag represents the HTTP header field "ETag",
// which is used for web cache validation and conditional requests.
const HeaderETag = "ETag"

// HeaderAccept represents the HTTP header field "Accept",
// which is used by clients to specify the media types they are willing to receive in the response.
const HeaderAccept = "Accept"

// HeaderAcceptEncoding represents the HTTP header key "Accept-Encoding",
// which is used by clients to indicate the content encoding (e.g., gzip, deflate)
// they can understand and accept in the server's response.
const HeaderAcceptEncoding = "Accept-Encoding"

// HeaderAcceptLanguage represents the HTTP header key "Accept-Language",
// which is used by clients to specify the preferred languages for the response.
const HeaderAcceptLanguage = "Accept-Language"

// ContentTypeJSON represents the MIME type for JSON content, typically used in API responses.
const ContentTypeJSON = "application/json"

// ContentTypeXIcon represents the MIME type for icon files in the x-icon format.
// This is an alternative MIME type for favicon.ico files in web applications.
const ContentTypeXIcon = "image/x-icon"
