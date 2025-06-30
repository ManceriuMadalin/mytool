package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var source, dest string

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Perform incremental backup",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Backing up from %s to %s...\n", source, dest)
		filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			relPath, _ := filepath.Rel(source, path)
			destPath := filepath.Join(dest, relPath)
			if _, err := os.Stat(destPath); os.IsNotExist(err) {
				os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
				copyFile(path, destPath)
				fmt.Println("Copied:", relPath)
			}
			return nil
		})
	},
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

func init() {
	backupCmd.Flags().StringVarP(&source, "source", "s", "", "Source directory")
	backupCmd.Flags().StringVarP(&dest, "dest", "d", "", "Destination directory")
	backupCmd.MarkFlagRequired("source")
	backupCmd.MarkFlagRequired("dest")
	rootCmd.AddCommand(backupCmd)
}
