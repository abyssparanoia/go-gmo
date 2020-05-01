package sftp

import (
	"context"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

// Client ... ssh client
type Client struct {
	sshConfig *ssh.ClientConfig
	addr      string
}

// NewClient ... new client for sftp
func NewClient(
	ctx context.Context,
	user,
	host,
	port,
	privateKeyPath string,
) (*Client, error) {
	authMethod := []ssh.AuthMethod{}
	key, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}
	authMethod = append(authMethod, ssh.PublicKeys(signer))

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: authMethod,
		// TOOD: change right callback
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := fmt.Sprintf("%s:%s", host, port)
	return &Client{sshConfig, addr}, nil
}
