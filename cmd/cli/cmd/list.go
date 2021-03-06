/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Body define the struct taht goes in Response
type Body struct {
	Bytes  []byte
	String string
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all the articles from the server",
	Long:  `list all the articles from the server`,
	Run: func(cmd *cobra.Command, args []string) {
		url := viper.GetString("url")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Id", "Title"})
		a := listArticles(url)

		for i, v := range a {
			// fmt.Printf("Article title %s: %s\n", strconv.Itoa(i), v.Title)
			var l []string
			l = append(l, strconv.Itoa(i), v.Title)
			table.Append(l)
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listArticles(url string) []article {
	c := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	res, err := c.Do(req)
	if err != nil {
		log.Fatalf("can not connect to URL: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("error calling to URL: %v", res.StatusCode)
	}

	var a []article
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &a)
	if err != nil {
		log.Fatalf("can not decode the JSON output: %v", err)
	}

	return a
}
