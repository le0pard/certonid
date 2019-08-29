package cmd

import (
	"strings"

	"github.com/le0pard/certonid/adapters/awscloud"
	"github.com/le0pard/certonid/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	decryptType          string
	decryptAwsKmsRegion  string
	decryptAwsKmsProfile string

	decryptCmd = &cobra.Command{
		Use:   "decrypt [OPTIONS] TEXT",
		Short: "Decrypt text",
		Long:  `Decrypt text with symmetric or kms encryption`,
                Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				err  error
				text []byte
			)


			switch strings.ToLower(decryptType) {
			case "aws_kms":
				kmsClient := awscloud.New(decryptAwsKmsProfile).KmsClient(decryptAwsKmsRegion)
				text, err = kmsClient.KmsDecryptText(args[0])
			default: // symmetric
				text, err = utils.SymmetricDecrypt(args[0])
			}

			if err != nil {
				er(err)
			}

			log.WithFields(log.Fields{
				"text": string(text),
			}).Info("Successfully decrypted")
		},
	}
)

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().StringVarP(&decryptType, "type", "t", "symmetric", "Decryption type (symmetric, aws_kms, gcloud_kms)")
	decryptCmd.Flags().StringVar(&decryptAwsKmsProfile, "aws-kms-profile", "", "AWS KMS Profile")
	decryptCmd.Flags().StringVar(&decryptAwsKmsRegion, "aws-kms-region", "", "AWS KMS Region")
}
