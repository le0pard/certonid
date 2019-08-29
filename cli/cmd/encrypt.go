package cmd

import (
	"strings"

	"github.com/le0pard/certonid/adapters/awscloud"
	"github.com/le0pard/certonid/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	encryptType          string
	encryptAwsKmsKeyID   string
	encryptAwsKmsProfile string
	encryptAwsKmsRegion  string

	encryptCmd = &cobra.Command{
		Use:   "encrypt [OPTIONS] TEXT",
		Short: "Encrypt text",
		Long:  `Encrypt text with symmetric or kms encryption`,
                Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var (
				err     error
				encText string
			)

			switch strings.ToLower(encryptType) {
			case "aws_kms":
				kmsClient := awscloud.New(encryptAwsKmsProfile).KmsClient(encryptAwsKmsRegion)
				encText, err = kmsClient.KmsEncryptText(encryptAwsKmsKeyID, []byte(args[0]))
			default: // symmetric
				encText, err = utils.SymmetricEncrypt([]byte(args[0]))
			}

			if err != nil {
				er(err)
			}

			log.WithFields(log.Fields{
				"text": encText,
			}).Info("Successfully encrypted")
		},
	}
)

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().StringVarP(&encryptType, "type", "t", "symmetric", "Encryption type (symmetric, aws_kms, gcloud_kms)")
	encryptCmd.Flags().StringVar(&encryptAwsKmsKeyID, "aws-kms-key-id", "", "AWS KMS Key ID")
	encryptCmd.Flags().StringVar(&encryptAwsKmsProfile, "aws-kms-profile", "", "AWS KMS Profile")
	encryptCmd.Flags().StringVar(&encryptAwsKmsRegion, "aws-kms-region", "", "AWS KMS Region")
}
