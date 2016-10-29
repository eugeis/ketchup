package bolt

import (
	"github.com/boltdb/bolt"
	"github.com/golang/protobuf/proto"
	"github.com/octavore/press/proto/press/models"
	"github.com/satori/go.uuid"
)

const ROUTE_BUCKET = "routes"

func (m *Module) UpdateRoute(route *models.Route) error {
	if route.Uuid == nil {
		route.Uuid = proto.String(uuid.NewV4().String())
	}
	return m.Update(ROUTE_BUCKET, route)
}

func (m *Module) GetRoute(key string) (*models.Route, error) {
	route := &models.Route{}
	err := m.Get(ROUTE_BUCKET, key, route)
	if err != nil {
		return nil, err
	}
	return route, nil
}

func (m *Module) ListRoutes() ([]*models.Route, error) {
	lst := []*models.Route{}
	err := m.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(ROUTE_BUCKET))
		return b.ForEach(func(_, v []byte) error {
			pb := &models.Route{}
			err := proto.Unmarshal(v, pb)
			if err != nil {
				return err
			}
			lst = append(lst, pb)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return lst, nil
}
