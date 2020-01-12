//+build !test

package appy

import (
	"encoding/hex"
	"fmt"
	"os"
)

func newConfigEncCommand(config *Config, logger *Logger, support Supporter) *Command {
	return &Command{
		Use:   "config:enc",
		Short: "Encrypt a config value using the key in `configs/<APPY_ENV>.key` or `APPY_MASTER_KEY`",
		Args:  ExactArgs(1),
		Run: func(cmd *Command, args []string) {
			if len(config.Errors()) > 0 {
				logger.Error(config.Errors()[0])
				os.Exit(-1)
			}

			masterKey := config.MasterKey()
			if masterKey == nil {
				logger.Fatal(config.Errors())
			}

			plaintext := []byte(args[0])
			ciphertext, err := support.AESEncrypt(plaintext, masterKey)
			if err != nil {
				logger.Fatal(err)
			}

			fmt.Println(hex.EncodeToString(ciphertext))
		},
	}
}
