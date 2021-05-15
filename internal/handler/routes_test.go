package handler

import (
	"testing"

	"github.com/kshamko/ing/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestInsertSort(t *testing.T) {
	data := []*models.RoutesRoutesItems0{
		{
			Route: models.Route{Distance: 15, Duration: 7},
		},
		{
			Route: models.Route{Distance: 14, Duration: 7},
		},
		{
			Route: models.Route{Distance: 150, Duration: 8},
		},
		{
			Route: models.Route{Distance: 15, Duration: 9},
		},
		{
			Route: models.Route{Distance: 15, Duration: 6},
		},
		{
			Route: models.Route{Distance: 15, Duration: 11},
		},
	}

	dataSorted := []*models.RoutesRoutesItems0{
		{
			Route: models.Route{Distance: 15, Duration: 6},
		},
		{
			Route: models.Route{Distance: 14, Duration: 7},
		},
		{
			Route: models.Route{Distance: 15, Duration: 7},
		},
		{
			Route: models.Route{Distance: 150, Duration: 8},
		},
		{
			Route: models.Route{Distance: 15, Duration: 9},
		},
		{
			Route: models.Route{Distance: 15, Duration: 11},
		},
	}

	res := []*models.RoutesRoutesItems0{}

	for _, d := range data {
		res = insertSort(res, d)
	}

	assert.Equal(t, res, dataSorted)
}
