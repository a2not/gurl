/*
Copyright © 2021 Naoki Honda

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var (
	errUrlNotGiven = errors.New("url is not given as the first command line argument")
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "execute HTTP GET request",
	Long:  `gurl get is subcommand for sending HTTP GET request to a given url.`,
	RunE:  httpGet,
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func httpGet(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errUrlNotGiven
	}
	return getURL(args[0], os.Stdout)
}

func getURL(url string, w io.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if _, err := w.Write([]byte(fmt.Sprintf("%v %v\n", resp.Proto, resp.Status))); err != nil {
		return err
	}

	for k, v := range resp.Header {
		if _, err := w.Write([]byte(fmt.Sprintf("%v : %v\n", k, v))); err != nil {
			return err
		}
	}
	return nil
}
