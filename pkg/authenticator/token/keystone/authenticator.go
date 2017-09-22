package keystone

import (
	"encoding/json"
	"errors"
	"github.com/gophercloud/gophercloud"
	"io/ioutil"
	"k8s.io/apiserver/pkg/authentication/user"
)

type KeystoneAuthenticator struct {
	client *gophercloud.ServiceClient
}

func (keystoneAuthenticator *KeystoneAuthenticator) AuthenticateToken(token string) (user.Info, bool, error) {
	requestOpts := gophercloud.RequestOpts{
		MoreHeaders: map[string]string{
			"X-Auth-Token":    token,
			"X-Subject-Token": token,
		},
	}
	url := keystoneAuthenticator.client.ServiceURL("auth", "tokens")
	response, err := keystoneAuthenticator.client.Request("GET", url, &requestOpts)
	if err != nil {
		return nil, false, errors.New("Failed to authenticate")
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, false, errors.New("Failed to authenticate")
	}

	// ignore catalog and role
	obj := struct {
		Token struct {
			User struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"user"`
			Project struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"project"`
		} `json:"token"`
	}{}

	err = json.Unmarshal(bodyBytes, &obj)
	if err != nil {
		return nil, false, errors.New("Failed to authenticate")
	}

	authenticatedUser := &user.DefaultInfo{
		Name:   obj.Token.User.Name,
		UID:    obj.Token.User.Id,
		Groups: []string{obj.Token.Project.Id},
	}

	return authenticatedUser, true, nil
}
