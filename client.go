package httpclient

import (
	"net/http"
	"time"
)

// Client 也可以擴展 timeout / headers
var defaultClient = &http.Client{Timeout: 10 * time.Second}
