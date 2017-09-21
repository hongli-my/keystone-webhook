package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"k8s.io/apiserver/pkg/authentication/authenticator"
	"keystone-webhook/pkg/authenticator/token/keystone"
	"keystone-webhook/pkg/handler"
	"keystone-webhook/pkg/options"
	"net/http"
	"os"
)

func main() {
	config := options.NewConfig()
	config.AddFlag(pflag.CommandLine)
	pflag.Parse()

	authenticationHandler, err := keystone.NewKeystoneAuthenticator(config.KeystoneUrl)
	if err != nil {
		fmt.Println("New Keystone client error: ", err)
		os.Exit(1)
	}
	http.Handle("/webhook", webHookServer(authenticationHandler))
	fmt.Println("Start Keystone-webhook, Listen :8443")
	http.ListenAndServeTLS(":8443", config.TlsCertFile, config.TlsPrivateKey, nil)
}

func webHookServer(authenticator authenticator.Token) http.Handler {
	return &handler.WebhookHandler{
		Authenticator: authenticator,
	}
}
