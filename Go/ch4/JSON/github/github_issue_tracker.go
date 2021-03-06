// Parsing Github Open Issues in JSON
// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"samgithub"
)

func main() {
	// issue := []string{"repo:golang/go is:open json decoder"}
	// result, err := github.SearchIssues(os.Args[1:])
	result, err := samgithub.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
