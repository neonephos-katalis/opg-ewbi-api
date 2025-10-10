package deployment

import (
	"github.com/nbycomp/neonephos-opg-ewbi-api/api/federation/models"
)

type InstallDeployment struct {
	*models.InstallAppJSONBody
	FederationContextID string
}
