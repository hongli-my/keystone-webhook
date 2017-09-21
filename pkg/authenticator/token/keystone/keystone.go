package keystone

import (
	"errors"

	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/utils"
)

func createIdentityV3Provider(options gophercloud.AuthOptions) (*gophercloud.ProviderClient, error) {
	client, err := openstack.NewClient(options.IdentityEndpoint)
	if err != nil {
		return nil, err
	}

	versions := []*utils.Version{
		{ID: "v3.0", Priority: 30, Suffix: "/v3/"},
	}
	chosen, _, err := utils.ChooseVersion(client, versions)
	if err != nil {
		return nil, fmt.Errorf("Unable to find identity API v3 version : %v", err)
	}

	switch chosen.ID {
	case "v3.0":
		return client, nil
	default:
		return nil, fmt.Errorf("Unsupported identity API version: %s", chosen.ID)
	}
}

func createKeystoneClient(authURL string) (*gophercloud.ServiceClient, error) {

	if authURL == "" {
		return nil, errors.New("Auth URL is empty")
	}

	opts := gophercloud.AuthOptions{IdentityEndpoint: authURL}
	provider, err := createIdentityV3Provider(opts)
	if err != nil {
		return nil, err
	}

	client, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		return nil, errors.New("Failed to authenticate")
	}
	if err != nil {
		return nil, errors.New("Failed to authenticate")
	}

	client.IdentityBase = client.IdentityEndpoint
	client.Endpoint = client.IdentityEndpoint
	return client, nil
}

func NewKeystoneAuthenticator(authURL string) (*KeystoneAuthenticator, error) {
	client, err := createKeystoneClient(authURL)
	if err != nil {
		return nil, err
	}

	return &KeystoneAuthenticator{client: client}, nil
}
