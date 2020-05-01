package sftp

import (
	"context"
	"io"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// SendAccountTransfer ... send account transfer file
func (c *Client) SendAccountTransfer(
	ctx context.Context,
	file *os.File,
) error {

	conn, err := ssh.Dial("tcp", c.addr, c.sshConfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	client, err := sftp.NewClient(conn)
	if err != nil {
		return err
	}
	defer client.Close()

	targetFile, err := client.Create(file.Name())
	if err != nil {
		return err
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, file)
	if err != nil {
		return err
	}

	return nil
}

// GetAccountTransfer ... get account transfer file
func (c *Client) GetAccountTransfer(
	ctx context.Context,
	file *os.File,
) error {
	conn, err := ssh.Dial("tcp", c.addr, c.sshConfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	client, err := sftp.NewClient(conn)
	if err != nil {
		return err
	}
	defer client.Close()

	targetFile, err := client.OpenFile(file.Name(), os.O_RDONLY)

	_, err = io.Copy(file, targetFile)
	if err != nil {
		return err
	}

	return nil
}
