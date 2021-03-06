package servicebus

import "github.com/Azure/open-service-broker-azure/pkg/service"

// ProvisioningParameters encapsulates non-sensitive Service Bus specific
// provisioning options
type ProvisioningParameters struct{}

// SecureProvisioningParameters encapsulates sensitive Service Bus specific
// provisioning options
type SecureProvisioningParameters struct{}

type serviceBusInstanceDetails struct {
	ARMDeploymentName       string `json:"armDeployment"`
	ServiceBusNamespaceName string `json:"serviceBusNamespaceName"`
}

type serviceBusSecureInstanceDetails struct {
	ConnectionString string `json:"connectionString"`
	PrimaryKey       string `json:"primaryKey"`
}

// BindingParameters encapsulates non-sensitive Service Bus specific binding
// options
type BindingParameters struct {
}

// SecureBindingParameters encapsulates sensitive Service Bus specific binding
// options
type SecureBindingParameters struct {
}

type serviceBusBindingDetails struct {
}

type serviceBusSecureBindingDetails struct {
}

// Credentials encapsulates Service Bus-specific coonection details and
// credentials.
type Credentials struct {
	ConnectionString string `json:"connectionString"`
	PrimaryKey       string `json:"primaryKey"`
}

func (
	s *serviceManager,
) GetEmptyProvisioningParameters() service.ProvisioningParameters {
	return &ProvisioningParameters{}
}

func (
	s *serviceManager,
) GetEmptySecureProvisioningParameters() service.SecureProvisioningParameters {
	return &SecureProvisioningParameters{}
}

func (
	s *serviceManager,
) GetEmptyInstanceDetails() service.InstanceDetails {
	return &serviceBusInstanceDetails{}
}

func (
	s *serviceManager,
) GetEmptySecureInstanceDetails() service.SecureInstanceDetails {
	return &serviceBusSecureInstanceDetails{}
}

func (s *serviceManager) GetEmptyBindingParameters() service.BindingParameters {
	return &BindingParameters{}
}

func (
	s *serviceManager,
) GetEmptySecureBindingParameters() service.SecureBindingParameters {
	return &SecureBindingParameters{}
}

func (s *serviceManager) GetEmptyBindingDetails() service.BindingDetails {
	return &serviceBusBindingDetails{}
}

func (
	s *serviceManager,
) GetEmptySecureBindingDetails() service.SecureBindingDetails {
	return &serviceBusSecureBindingDetails{}
}
