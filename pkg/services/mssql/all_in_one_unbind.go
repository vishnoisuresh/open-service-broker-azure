package mssql

import (
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (a *allInOneManager) Unbind(
	instance service.Instance,
	binding service.Binding,
) error {
	dt, ok := instance.Details.(*allInOneInstanceDetails)
	if !ok {
		return fmt.Errorf(
			"error casting instance.Details as *mssql.allInOneInstanceDetails",
		)
	}
	sdt, ok := instance.SecureDetails.(*secureAllInOneInstanceDetails)
	if !ok {
		return fmt.Errorf(
			"error casting instance.SecureDetails as " +
				"*mssql.secureAllInOneInstanceDetails",
		)
	}
	bd, ok := binding.Details.(*bindingDetails)
	if !ok {
		return fmt.Errorf(
			"error casting binding.Details as *mssql.bindingDetails",
		)
	}

	return unbind(
		dt.AdministratorLogin,
		sdt.AdministratorLoginPassword,
		dt.FullyQualifiedDomainName,
		dt.DatabaseName,
		bd,
	)
}
