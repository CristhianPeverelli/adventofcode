package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url, filePath string) error {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to download file, status code: %d", resp.StatusCode)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("File downloaded successfully to: %s\n", filePath)
	return nil
}
