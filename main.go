package main

import (
	"os"
	"search-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// CreateTable
	// createdTable, err := dynamodb.CreateTable("search-service-trie")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(createdTable)

	// ListTables
	// tableNames, err := dynamodb.ListTables()
	// if err != nil {
	// 	return
	// }
	// fmt.Print(tableNames)

	// AddItem
	// trieNode := dynamodb.TrieNode{
	// 	Prefix: "c",
	// 	FrequentQueries: map[string]int{
	// 		"cat": 5,
	// 		"camera": 4,
	// 		"car": 3,
	// 		"camel": 2,
	// 		"close": 2,
	// 	},
	// 	ChildNodes: []string{"ca", "cl"},
	// }

	// addedItem, err := dynamodb.AddItem(trieNode, "search-service-trie")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(addedItem)

	// ReadItem
	// itemRead, err := dynamodb.ReadItem("c", "search-service-trie")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(itemRead)

	// UpdateItem
	// updatedTrieNode := dynamodb.TrieNode{
	// 	Prefix: "c",
	// 	FrequentQueries: map[string]int{
	// 		"cool": 5,
	// 		"camera": 4,
	// 		"car": 3,
	// 		"camel": 2,
	// 		"close": 2,
	// 	},
	// 	ChildNodes: []string{"ca", "co", "cl"},
	// }
	// updatedItem, err := dynamodb.UpdateItem(updatedTrieNode, "search-service-trie")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(updatedItem)

	// DeleteItem
	// prefix := "c"
	// deletedItem, err := dynamodb.DeleteItem(prefix, "search-service-trie")
	// if err != nil {
	// 	return
	// }
	// fmt.Print(deletedItem)

	// AddSearchQuery
	// searchQueries := []string{
	// 	"ark", "all", "algo", "do", "dag", "daily", "daisy",
	// }
	// tableName := "search-service-trie"
	// for _, searchQuery := range searchQueries {
	// 	res, err := trie.AddSearchQuery(searchQuery, tableName)
	// 	if err != nil {
	// 		return
	// 	}
	// 	fmt.Print(res)
	// }

	// Search
	// searchQuery := "a"
	// tableName := "search-service-trie"
	// searchResult, err := trie.Search(searchQuery, tableName)
	// if err != nil {
	// 	return
	// }
	// fmt.Print(searchResult)

	// DeleteTrieNode
	// searchQuery := "dag"
	// tableName := "search-service-trie"
	// deleteResult, err := trie.DeleteTrieNode(searchQuery, tableName)
	// if err != nil {
	// 	return
	// }
	// fmt.Print(deleteResult)

	godotenv.Load()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(os.Getenv("PORT"))
}