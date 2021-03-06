package jwk

import (
	"crypto"
	"time"

	"github.com/lestrrat-go/backoff/v2"
	"github.com/lestrrat-go/option"
)

type Option = option.Interface

type identHTTPClient struct{}
type identThumbprintHash struct{}
type identRefreshInterval struct{}
type identMinRefreshInterval struct{}
type identFetchBackoff struct{}

// AutoRefreshOption is a type of Option that can be passed to the
// AutoRefresh object.
type AutoRefreshOption interface {
	Option
	autoRefreshOption()
}

type autoRefreshOption struct {
	Option
}

func (*autoRefreshOption) autoRefreshOption() {}

// FetchOption is a type of Option that can be passed to `jwk.Fetch()`
// This type also implements the `AutoRefreshOption`, and thus can be
// safely passed to `(*jwk.AutoRefresh).Configure()`
type FetchOption interface {
	AutoRefreshOption
	fetchOption()
}

type fetchOption struct {
	Option
}

func (*fetchOption) autoRefreshOption() {}
func (*fetchOption) fetchOption()       {}

// WithHTTPClient allows users to specify the "net/http".Client object that
// is used when fetching jwk.Set objects.
func WithHTTPClient(cl HTTPClient) FetchOption {
	return &fetchOption{option.New(identHTTPClient{}, cl)}
}

// WithFetchBackoff specifies the backoff policy to use when
// refreshing a JWKS from a remote server fails.
//
// This does not have any effect on initial `Fetch()`, or any of the `Refresh()` calls --
// the backoff is applied ONLY on the background refreshing goroutine.
func WithFetchBackoff(v backoff.Policy) FetchOption {
	return &fetchOption{option.New(identFetchBackoff{}, v)}
}

func WithThumbprintHash(h crypto.Hash) Option {
	return option.New(identThumbprintHash{}, h)
}

// WithRefreshInterval specifies the static interval between refreshes
// of jwk.Set objects controlled by jwk.AutoRefresh.
//
// Providing this option overrides the adaptive token refreshing based
// on Cache-Control/Expires header (and jwk.WithMinRefreshInterval),
// and refreshes will *always* happen in this interval.
func WithRefreshInterval(d time.Duration) AutoRefreshOption {
	return &autoRefreshOption{
		option.New(identRefreshInterval{}, d),
	}
}

// WithMinRefreshInterval specifies the minimum refresh interval to be used
// when using AutoRefresh. This value is ONLY used if you did not specify
// a user-supplied static refresh interval via `WithRefreshInterval`.
//
// This value is used as a fallback value when tokens are refreshed.
//
// When we fetch the key from a remote URL, we first look at the max-age
// directive from Cache-Control response header. If this value is present,
// we compare the max-age value and the value specified by this option
// and take the larger one.
//
// Next we check for the Expires header, and similarly if the header is
// present, we compare it against the value specified by this option,
// and take the larger one.
//
// Finally, if neither of the above headers are present, we use the
// value specified by this option as the next refresh timing
//
// If unspecified, the minimum refresh interval is 1 hour
func WithMinRefreshInterval(d time.Duration) AutoRefreshOption {
	return &autoRefreshOption{
		option.New(identMinRefreshInterval{}, d),
	}
}
