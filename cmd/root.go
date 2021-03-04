package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/hum/peep/internal"
	"github.com/spf13/cobra"
)

var (
	domainName string
	domainFile string

	whois = internal.Whois{}

	rootCmd = &cobra.Command{
		Use:   "peep",
		Short: "Peep 0.0.1: 🐥 Search for available domains",
		Run:   getDomains,
	}
)

func getDomains(cmd *cobra.Command, args []string) {
	if len(args) > 0 || len(domainName) == 0 || len(domainFile) == 0 {
		cmd.Help()
		os.Exit(0)
	}

	data, err := os.ReadFile(domainFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	domains := strings.Split(string(data), "\n")
	whois.Domains = domains

	result, err := whois.Search(domainName)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&domainName, "name", "n", "", "domain name to search for")
	rootCmd.PersistentFlags().StringVarP(&domainFile, "file", "f", "", "text file containing all of the domains")
	rootCmd.MarkFlagRequired("name")
	rootCmd.MarkFlagRequired("file")
}
