package sftp

import (
	"context"
	"io"
	"os"
)

// SendAccountTransfer ... send account transfer file
func (cli *Client) SendAccountTransfer(
	ctx context.Context,
	file *os.File,
) error {
	targetFile, err := cli.Create(file.Name())
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
