package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "index",
		Short: "index knowledge",
		Long:  `index knowledge`,
		Run: func(cmd *cobra.Command, args []string) {
			//ctx := context.Background()
			fmt.Println("Hello Cobra")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

	for {
		fmt.Println("Hello")
		time.Sleep(5 * time.Second)
	}
}
