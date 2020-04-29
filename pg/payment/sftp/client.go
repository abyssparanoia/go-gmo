package sftp

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/pkg/sftp"

	"golang.org/x/crypto/ssh"
)

// Client ... ssh client
type Client struct {
	*sftp.Client
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
		User:            user,
		Auth:            authMethod,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), sshConfig)
	if err != nil {
		return nil, err
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &Client{Client: client}, nil
}
