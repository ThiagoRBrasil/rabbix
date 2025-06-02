package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todos os casos de teste salvos",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		testDir := filepath.Join(homeDir, ".rabbix", "tests")

		files, err := ioutil.ReadDir(testDir)
		if err != nil {
			fmt.Printf("Erro ao acessar diretório: %v\n", err)
			return
		}

		fmt.Println("📄 Casos de teste salvos:")

		for _, file := range files {
			if filepath.Ext(file.Name()) == ".json" {
				path := filepath.Join(testDir, file.Name())
				data, err := ioutil.ReadFile(path)
				if err != nil {
					continue
				}

				var test map[string]interface{}
				if err := json.Unmarshal(data, &test); err != nil {
					continue
				}

				name := test["name"]
				rk := test["route_key"]
				fmt.Printf("🧪 %s  (routeKey: %s)\n", name, rk)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
