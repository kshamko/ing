package datasource

import "fmt"

// Point struct represent a coordinate.
type Point struct {
	Lat string
	Lon string
}

// String prints point as "lat,lon".
func (p Point) String() string {
	return fmt.Sprintf("%s,%s", p.Lat, p.Lon)
}

// RoutingData object returned by OSRM server.
type RoutingData struct {
	Code   string  `json:"code"`
	Routes []Route `json:"routes"`
}

// Route object from OSRM response.
type Route struct {
	Destination string
	Duration    float32 `json:"duration"`
	Distance    float32 `json:"distance"`
}

// RouteErrResponse represent an error response from OSRM.
type RouteErrResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
