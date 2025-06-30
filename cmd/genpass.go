package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
)

var length int

var genpassCmd = &cobra.Command{
	Use:   "genpass",
	Short: "Generate a secure password",
	Run: func(cmd *cobra.Command, args []string) {
		chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
		result := ""
		for i := 0; i < length; i++ {
			num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
			result += string(chars[num.Int64()])
		}
		fmt.Println("Generated password:", result)
	},
}

func init() {
	genpassCmd.Flags().IntVarP(&length, "length", "l", 12, "Length of the password")
	rootCmd.AddCommand(genpassCmd)
}
