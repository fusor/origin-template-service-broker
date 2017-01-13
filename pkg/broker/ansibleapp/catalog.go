package ansibleapp

import (
	"github.com/fusor/origin-template-service-broker/pkg/broker"
)

func (b Broker) Catalog() (*broker.CatalogResponse, error) {
	return nil, notImplemented
}
