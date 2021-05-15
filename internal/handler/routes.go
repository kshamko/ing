package handler

import (
	"context"
	"net/http"
	"sort"

	"github.com/go-openapi/runtime/middleware"
	"github.com/kshamko/ing/internal/datasource"
	"github.com/kshamko/ing/internal/models"
	"github.com/kshamko/ing/internal/restapi/operations/routes"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

// RoutesDatasource interface for routing data.
// Needed in case we want to mock it and cover handler with tests.
type RoutesDatasource interface {
	GetRoutes(ctx context.Context, src, dst datasource.Point) (datasource.RoutingData, error)
}

// Routes struct to define Handle func on.
type Routes struct {
	ds RoutesDatasource
}

// NewRoutes return Routes object.
func NewRoutes(ds RoutesDatasource) *Routes {
	return &Routes{
		ds: ds,
	}
}

// Handle function processes http request, needed by swagger generated code.
func (rt *Routes) Handle(in routes.RoutesParams) middleware.Responder {
	srcPoint := datasource.Point{
		Lat: in.Src[0],
		Lon: in.Src[1],
	}

	g, ctx := errgroup.WithContext(context.Background())

	rdChan, resChan := combineRoutingData(ctx)

	for _, dst := range in.Dst {
		dst := dst

		g.Go(func() error {
			dstPoint := datasource.Point{
				Lat: dst[0],
				Lon: dst[1],
			}

			routingData, err := rt.ds.GetRoutes(ctx, srcPoint, dstPoint)
			for i := range routingData.Routes {
				routingData.Routes[i].Destination = dstPoint.String()
			}
			rdChan <- routingData

			if err != nil {
				log.WithFields(log.Fields{
					"src": srcPoint,
					"dst": dstPoint,
				}).Error(err.Error())
			}

			return err
		})
	}

	if err := g.Wait(); err != nil {
		return routes.NewRoutesInternalServerError().WithPayload(
			&models.APIInvalidResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		)
	}

	result := <-resChan
	close(resChan)

	result.Source = srcPoint.String()

	return routes.NewRoutesOK().WithPayload(
		&result,
	)
}

func combineRoutingData(ctx context.Context) (chan datasource.RoutingData, chan models.Routes) {
	c := make(chan datasource.RoutingData)
	res := make(chan models.Routes)

	go func() {
		result := models.Routes{}

		for {
			select {
			case <-ctx.Done():
				res <- result

				return

			case data := <-c:
				for _, dst := range data.Routes {
					result.Routes = insertSort(
						result.Routes,
						&models.RoutesRoutesItems0{
							Route: models.Route{
								Distance:    dst.Distance,
								Duration:    dst.Duration,
								Destination: dst.Destination,
							},
						},
					)
				}
			}
		}
	}()

	return c, res
}

// once we still iterate over the data to remap it, so lets get this data sorted.
func insertSort(data []*models.RoutesRoutesItems0, el *models.RoutesRoutesItems0) []*models.RoutesRoutesItems0 {
	index := sort.Search(len(data), func(i int) bool {
		if data[i].Duration == el.Duration {
			return data[i].Distance > el.Distance
		}

		return data[i].Duration > el.Duration
	})

	data = append(data, &models.RoutesRoutesItems0{})
	copy(data[index+1:], data[index:])
	data[index] = el

	return data
}
