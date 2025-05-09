package solcast

import (
	"fmt"
	"net/http"
	"sync"
)

// Solcast API allows only allows a small number of calls per day
const maxAPICallsPerDay = 5 // actually it's 10 but we want to be safe, we should only need 1

var (
    apiCallCount int
    apiCallMutex sync.Mutex
)

// CustomTransport wraps the default transport to add rate-limiting logic
type CustomTransport struct {
    BaseTransport http.RoundTripper
}

func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    // Check and increment the API call count
    apiCallMutex.Lock()
    defer apiCallMutex.Unlock()

    if apiCallCount >= maxAPICallsPerDay {
        return nil, fmt.Errorf("API call limit of %d exceeded", maxAPICallsPerDay)
    }

    apiCallCount++
    return t.BaseTransport.RoundTrip(req)
}

func init() {
    // Replace the default transport with the custom transport
    http.DefaultTransport = &CustomTransport{
        BaseTransport: http.DefaultTransport,
    }
}
