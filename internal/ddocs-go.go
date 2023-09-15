package internal

import (
	"bytes"
  "os/exec"
  "fmt"
	"compress/flate"
	"encoding/base64"
	"encoding/json"
	"io"
	"os"
)

func Compress(input string) (string, error) {
	var buf bytes.Buffer
	compressor, err := flate.NewWriter(&buf, flate.DefaultCompression)
  if err != nil {
    return "", err
  }
	compressor.Write([]byte(input))
	compressor.Close()
  base := base64.StdEncoding.EncodeToString(buf.Bytes())
	return base, nil
}

func Decompress(compressedData string) (string, error) {
  base, err := base64.StdEncoding.DecodeString(compressedData)
  if err != nil {
    return "", err
  }
  buf := bytes.NewReader([]byte(base))
	decompressor := flate.NewReader(buf)
	decompressedData, err := io.ReadAll(decompressor)
	decompressor.Close()
  if err != nil {
    return "", err
  }
  return string(decompressedData), nil
}


func WriteToFile(data map[string]string) error {
  r, err := json.Marshal(data)
  if err != nil {
    return err
  }

  home, err := os.UserHomeDir()
  if err != nil {
    return err
  }
  home += "/.config/ddocs/data.json"
  err = os.WriteFile(home, r, 0644)
  if err != nil {
    return err
  }
  return nil
}

func ReadFile() (map[string]string, error) {
  home, err := os.UserHomeDir()
  if err != nil {
    return nil, err
  }
  home += "/.config/ddocs/data.json"

  data, err := os.ReadFile(home)
  if err != nil {
    return nil, err
  }

  var datamap  map[string]string
  err = json.Unmarshal(data, &datamap)
  if err != nil {
    return nil, err
  }

  return datamap, nil
}

func Editor(old_text string) (string, error) {
	tmpfile, err := os.CreateTemp("", "default_text.txt")
	if err != nil {
		fmt.Println("Failed to create temporary file:", err)
		return "", err
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(old_text); err != nil {
		fmt.Println("Failed to write to temporary file:", err)
		return "", err
	}
	tmpfile.Close()

	cmd := exec.Command("vim", tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to run Vim:", err)
		return "", err
	}

	modifiedText, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		fmt.Println("Failed to read modified text:", err)
		return "", err
	}

	return string(modifiedText), nil
}
