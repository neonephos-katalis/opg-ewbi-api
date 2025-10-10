package deployment

import (
	"context"
	"errors"

	k8scl "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nbycomp/neonephos-opg-ewbi-api/pkg/metastore"
)

var _ Client = &client{}

type Client interface {
	Install(ctx context.Context, app *InstallDeployment) (string, error)
	Uninstall(ctx context.Context, federationContextID, id string) error
}

func NewClient(k8sClient k8scl.Client, namespace string) *client {
	return &client{
		appMetaClient: metastore.NewK8sClient(k8sClient, namespace),
	}
}

type client struct {
	appMetaClient metastore.Client
}

func (c *client) Install(ctx context.Context, dep *InstallDeployment) (string, error) {
	if err := c.appMetaClient.AddApplicationInstance(ctx, &metastore.ApplicationInstance{
		InstallAppJSONBody:  dep.InstallAppJSONBody,
		FederationContextId: dep.FederationContextID,
	}); err != nil {
		return "", err
	}

	return dep.AppInstanceId, nil
}

func (c *client) Uninstall(ctx context.Context, federationContextID, id string) error {
	if err := c.appMetaClient.RemoveApplicationInstance(ctx, federationContextID, id); err != nil && !errors.Is(err, metastore.ErrNotFound) {
		return err
	}

	return nil
}
