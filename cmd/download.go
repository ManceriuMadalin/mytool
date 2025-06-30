package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

var url, downloadDest string

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download an image from URL",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()

		filename := path.Base(url)
		out, err := os.Create(filepath.Join(downloadDest, filename))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Downloaded:", filename)
	},
}

func init() {
	downloadCmd.Flags().StringVarP(&url, "url", "u", "", "URL of the image")
	downloadCmd.Flags().StringVarP(&downloadDest, "dest", "d", ".", "Destination directory")
	downloadCmd.MarkFlagRequired("url")
	rootCmd.AddCommand(downloadCmd)
}
