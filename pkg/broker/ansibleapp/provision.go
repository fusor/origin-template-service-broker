package ansibleapp

import (
	"github.com/fusor/origin-template-service-broker/pkg/broker"
	"github.com/pborman/uuid"
)

func (b Broker) Provision(instanceUUID uuid.UUID, req *broker.ProvisionRequest) (*broker.ProvisionResponse, error) {
	return nil, notImplemented
}

func (b Broker) Update(instanceUUID uuid.UUID, req *broker.UpdateRequest) (*broker.UpdateResponse, error) {
	return nil, notImplemented // TODO
}

func (b Broker) Deprovision(instanceUUID uuid.UUID) (*broker.DeprovisionResponse, error) {
	return nil, notImplemented // TODO
}
