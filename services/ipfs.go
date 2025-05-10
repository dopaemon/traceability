package services

import (
	"bytes"
	"encoding/json"
	_ "fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const LocalIPFSEndpoint = "http://127.0.0.1:5001/api/v0/add" // Kubo IPFS API endpoint

// IPFSResponse là cấu trúc phản hồi khi upload file lên Kubo.
type IPFSResponse struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
	Size string `json:"Size"`
}

// UploadToIPFS upload file lên Kubo và trả về hash của file.
func UploadToIPFS(filePath string) (string, error) {
	// Mở file cần upload
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Tạo multipart body để gửi file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return "", err
	}

	// Copy nội dung file vào phần của body
	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}

	// Đóng writer
	writer.Close()

	// Tạo HTTP request gửi lên Kubo IPFS
	req, err := http.NewRequest("POST", LocalIPFSEndpoint, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Gửi request đến Kubo
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Parse phản hồi từ IPFS
	var ipfsRes IPFSResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&ipfsRes); err != nil {
		return "", err
	}

	// Trả về hash của file đã upload lên IPFS
	return ipfsRes.Hash, nil
}
