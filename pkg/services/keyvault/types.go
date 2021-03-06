package keyvault

import "github.com/Azure/open-service-broker-azure/pkg/service"

// ProvisioningParameters encapsulates non-sensitive keyvault-specific
// provisioning options
type ProvisioningParameters struct {
	ObjectID string `json:"objectId"`
	ClientID string `json:"clientId"`
}

// SecureProvisioningParameters encapsulates sensitive keyvault-specific
// provisioning options
type SecureProvisioningParameters struct {
	ClientSecret string `json:"clientSecret"`
}

type keyvaultInstanceDetails struct {
	ARMDeploymentName string `json:"armDeployment"`
	KeyVaultName      string `json:"keyVaultName"`
	VaultURI          string `json:"vaultUri"`
	ClientID          string `json:"clientId"`
}

type keyvaultSecureInstanceDetails struct {
	ClientSecret string `json:"clientSecret"`
}

// BindingParameters encapsulates non-sensitive keyvault-specific binding
// options
type BindingParameters struct {
}

// SecureBindingParameters encapsulates sensitive keyvault-specific binding
// options
type SecureBindingParameters struct {
}

type keyvaultBindingDetails struct {
}

type keyvaultSecureBindingDetails struct {
}

// Credentials encapsulates Key Vault-specific coonection details and
// credentials.
type Credentials struct {
	VaultURI     string `json:"vaultUri"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
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
	return &keyvaultInstanceDetails{}
}

func (
	s *serviceManager,
) GetEmptySecureInstanceDetails() service.SecureInstanceDetails {
	return &keyvaultSecureInstanceDetails{}
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
	return &keyvaultBindingDetails{}
}

func (
	s *serviceManager,
) GetEmptySecureBindingDetails() service.SecureBindingDetails {
	return &keyvaultSecureBindingDetails{}
}
