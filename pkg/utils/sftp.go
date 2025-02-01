package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"speakbuddy-be/pkg/config"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func UploadToSftp(localFilePath string, remoteFilePath string) error {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Printf("[ERROR] config init failed, err: %+v", err)
	}

	config := &ssh.ClientConfig{
		User: cfg.SftpUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(cfg.SftpPass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := fmt.Sprintf("%s:%s", cfg.SftpHost, cfg.SftpPort)
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Printf("[ERROR] failed to dial: %+v", err)
		return err
	}
	defer conn.Close()

	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		log.Printf("[ERROR] failed to create SFTP clien: %+v", err)
		return err
	}
	defer sftpClient.Close()

	localFile, err := os.Open(localFilePath)
	if err != nil {
		log.Printf("[ERROR] failed to open local file: %+v, err: %+v", localFilePath, err)
		return err
	}
	defer localFile.Close()

	remoteFile, err := sftpClient.Create(remoteFilePath)
	if err != nil {
		log.Printf("[ERROR] failed to create remote file: %+v, err: %+v", remoteFilePath, err)
		return err
	}
	defer remoteFile.Close()

	_, err = io.Copy(remoteFile, localFile)
	if err != nil {
		log.Printf("[ERROR] failed to copy file: %+v", err)
		return err
	}

	return nil
}

func DownloadFromSftp(remoteFilePath string, localFilePath string) error {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Printf("[ERROR] config init failed, err: %+v", err)
	}

	config := &ssh.ClientConfig{
		User: cfg.SftpUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(cfg.SftpPass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := fmt.Sprintf("%s:%s", cfg.SftpHost, cfg.SftpPort)
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Printf("[ERROR] failed to dial: %+v", err)
		return err
	}
	defer conn.Close()

	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		log.Printf("[ERROR] failed to create SFTP client: %+v", err)
		return err
	}
	defer sftpClient.Close()

	remoteFile, err := sftpClient.Open(remoteFilePath)
	if err != nil {
		log.Printf("[ERROR] failed to open remote file: %+v, err: %+v", remoteFilePath, err)
		return err
	}
	defer remoteFile.Close()

	localFile, err := os.Create(localFilePath)
	if err != nil {
		log.Printf("[ERROR] failed to create local file: %+v, err: %+v", localFilePath, err)
		return err
	}
	defer localFile.Close()

	_, err = io.Copy(localFile, remoteFile)
	if err != nil {
		log.Printf("[ERROR] failed to copy file: %+v", err)
		return err
	}

	return nil
}
