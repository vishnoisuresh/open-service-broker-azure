package search

import (
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (s *serviceManager) ValidateBindingParameters(
	service.BindingParameters,
	service.SecureBindingParameters,
) error {
	// There are no parameters for binding to,
	// so there is nothing to validate
	return nil
}

func (s *serviceManager) Bind(
	service.Instance,
	service.BindingParameters,
	service.SecureBindingParameters,
) (service.BindingDetails, service.SecureBindingDetails, error) {
	return &searchBindingDetails{}, &searchSecureBindingDetails{}, nil
}

func (s *serviceManager) GetCredentials(
	instance service.Instance,
	_ service.Binding,
) (service.Credentials, error) {
	dt, ok := instance.Details.(*searchInstanceDetails)
	if !ok {
		return nil, fmt.Errorf(
			"error casting instance.Details as *searchInstanceDetails",
		)
	}
	sdt, ok := instance.SecureDetails.(*searchSecureInstanceDetails)
	if !ok {
		return nil, fmt.Errorf(
			"error casting instance.SecureDetails as *searchSecureInstanceDetails",
		)
	}
	return &searchCredentials{
		ServiceName: dt.ServiceName,
		APIKey:      sdt.APIKey,
	}, nil
}
