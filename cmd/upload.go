/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"net/http"

	"encoding/json"

	"github.com/ekremparlak/gokapi_client/models"
	"github.com/mdp/qrterminal/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	maxDownload int
	expiry      int
	password    string
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file to a Gokapi Server",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: upload,
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringVarP(&password, "password", "p", "", "Set a password for the file")
	uploadCmd.Flags().IntVarP(&expiry, "expiry", "e", viper.GetInt("GOKAPI_EXPIRY"), "Set the expiry days for the file")
	uploadCmd.Flags().IntVarP(&maxDownload, "max_download", "m", viper.GetInt("GOKAPI_MAX_DOWNLOAD"), "Set the allowed downloads for the file")
}

func upload(cmd *cobra.Command, args []string) {
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(args[0]))
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.CopyBuffer(part, file, make([]byte, 32*1024))
	if err != nil {
		fmt.Println(err)
		return
	}

	writer.WriteField("allowedDownloads", strconv.Itoa(maxDownload))
	writer.WriteField("expiryDays", strconv.Itoa(expiry))
	writer.WriteField("password", password)

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", apiURL+"/api/files/add", body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("apiKey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Error: " + resp.Status)
		return
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	response := models.FileAddResponse{}
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response.URL + response.FileInfo.Id)
	qrterminal.Generate(response.URL+response.FileInfo.Id, qrterminal.L, os.Stdout)
}
