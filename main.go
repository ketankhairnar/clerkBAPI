package main

import (
	"fmt"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"log"
)

func main() {

	client, err := clerk.NewClient("")
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}
	// Define pagination parameters
	limit := 25
	offset := 0

	// Retrieve total user count
	count, err := client.Users().Count(clerk.ListAllUsersParams{})
	if err != nil {
		fmt.Println("Error counting users:", err)
		return
	}

	totalCount := count.TotalCount
	fmt.Printf("Total users: %d\n", totalCount)

	// Loop to paginate through the users
	for offset < totalCount {
		// List users for the current page
		users, err := client.Users().ListAll(clerk.ListAllUsersParams{
			Limit:  &limit,
			Offset: &offset,
		})
		if err != nil {
			log.Printf("Error listing users at offset %d: %v", offset, err)
			return
		}

		// Print user info
		for _, user := range users {
			fmt.Printf("%s, %s, %s\n", getFirstName(user), getLastName(user), user.EmailAddresses[0].EmailAddress)
		}

		// Update offset for next page
		offset += limit
	}

	fmt.Println("Completed listing all users.")

}

func getFirstName(user clerk.User) string {
	if user.FirstName != nil {
		return *user.FirstName
	}
	return ""
}

func getLastName(user clerk.User) string {
	if user.LastName != nil {
		return *user.LastName
	}
	return ""
}
