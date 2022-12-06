package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var optLen int
var numKeys int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-api-keygen",
	Short: "Basic util for creating API keys suitable for HTTP headers",
	Long:  `A basic utility for generating API keys. Base64 encodes cryptographically random bytes`,
	RunE:  genApiKey,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func genApiKey(cmd *cobra.Command, args []string) error {
	data := make([]byte, optLen)
	for i := 0; i < numKeys; i++ {
		_, err := rand.Read(data)
		if err != nil {
			return err
		}
		key := base64.RawURLEncoding.EncodeToString(data)
		fmt.Println(key)
	}

	return nil
}

func init() {
	rootCmd.Flags().IntVarP(&optLen, "length", "l", 8, "API key bytes")
	rootCmd.Flags().IntVarP(&numKeys, "count", "n", 1, "Number of keys to generate")
}
