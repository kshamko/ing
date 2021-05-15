package datasource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutesURI(t *testing.T) {
	rds, _ := NewRoutesOSRM()

	src := Point{
		Lat: "12.12",
		Lon: "14.14",
	}

	dst := Point{
		Lat: "15.15",
		Lon: "16.16",
	}
	uri := rds.routesURI(src, dst)
	assert.Equal(t, uri, "http://router.project-osrm.org/route/v1/driving/12.12,14.14;15.15,16.16?overview=false")
}
