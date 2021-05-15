package datasource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

const osrmBaseRouteURL = "http://router.project-osrm.org/route/v1/driving"

type RoutesOSRM struct {
	uri *url.URL
}

// ErrRoutes error.
var ErrRoutes = errors.New("routes error")

// NewRoutesOSRM creates OSRM client.
func NewRoutesOSRM() (*RoutesOSRM, error) {
	u, err := url.Parse(osrmBaseRouteURL)
	if err != nil {
		return nil, fmt.Errorf("cant parse osrm base url: %w", err)
	}

	return &RoutesOSRM{uri: u}, nil
}

// GetRoutes makes request for the routing data.
func (r *RoutesOSRM) GetRoutes(ctx context.Context, src, dst Point) (RoutingData, error) {
	uri := r.routesURI(src, dst)

	resp, err := r.request(ctx, uri, "GET", nil)
	if err != nil {
		return RoutingData{}, fmt.Errorf("%w:%v", ErrRoutes, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		routingData := RoutingData{}
		err = json.NewDecoder(resp.Body).Decode(&routingData)

		if err != nil {
			return RoutingData{}, fmt.Errorf("%w:%v", ErrRoutes, err)
		}

		if routingData.Code != "Ok" {
			return RoutingData{}, fmt.Errorf("%w:%s", ErrRoutes, "status not ok")
		}

		return routingData, nil
	}

	errResp := RouteErrResponse{}
	err = json.NewDecoder(resp.Body).Decode(&errResp)

	if err != nil {
		return RoutingData{}, fmt.Errorf("%w:%v", ErrRoutes, err)
	}

	return RoutingData{}, fmt.Errorf("%w:%s", ErrRoutes, errResp.Message)
}

func (r *RoutesOSRM) request(ctx context.Context, url string, method string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}

	request, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	return client.Do(request)
}

func (r *RoutesOSRM) routesURI(src, dst Point) string {
	u := *r.uri
	q := u.Query()
	q.Set("overview", "false")
	u.RawQuery = q.Encode()
	u.Path = path.Join(u.Path, fmt.Sprintf("%s;%s", src, dst))

	return u.String()
}
