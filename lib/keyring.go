package lib

import (
	"os"

	"github.com/99designs/keyring"
)

func keyringPrompt(prompt string) (string, error) {
	password := os.Getenv("AWS_OKTA_FILE_PASSPHRASE")
	if len(password) > 0 {
		return password, nil
	}
	
	return PromptWithOutput(prompt, true, os.Stderr)
}

func OpenKeyring(allowedBackends []keyring.BackendType) (kr keyring.Keyring, err error) {
	kr, err = keyring.Open(keyring.Config{
		AllowedBackends:          allowedBackends,
		KeychainTrustApplication: true,
		// this keychain name is for backwards compatibility
		ServiceName:             "aws-okta-login",
		LibSecretCollectionName: "awsvault",
		FileDir:                 "~/.aws-okta/",
		FilePasswordFunc:        keyringPrompt,
	})

	return
}
