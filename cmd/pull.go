package cmd

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/emc-advanced-dev/pkg/errors"
	"github.com/emc-advanced-dev/unik/pkg/client"
	"github.com/emc-advanced-dev/unik/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// pushCmd represents the push command
var pullCmd = &cobra.Command{
	Use:   "push",
	Short: "Push an image to a Unik Image Repository",
	Long: `
Example usage:
unik pull --imageName theirImage

Requires that you first authenticate to a unik image repository with 'unik login'
	`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := getHubConfig()
		if err != nil {
			logrus.Fatal(err)
		}
		if imageName == "" {
			logrus.Fatal("--imageName must be set")
		}
		if provider == "" {
			logrus.Fatal("--provider must be set")
		}
		if host == "" {
			host = clientConfig.Host
		}
		if err := client.UnikClient(host).Images().Pull(c, imageName, provider, force); err != nil {
			logrus.Fatal(err)
		}
		fmt.Println(imageName + " pushed")
	},
}

func init() {
	RootCmd.AddCommand(pullCmd)
	pullCmd.Flags().StringVar(&imageName, "imageName", "", "<string,required> image to pull")
	pullCmd.Flags().StringVar(&provider, "provider", "", "<string,required> name of the provider the image is built for")
	pullCmd.Flags().BoolVar(&force, "force", false, "<bool,optional> force overwriting local image of the same name")
}