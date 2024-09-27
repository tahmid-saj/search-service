package main

import (
	"search-service/dynamodb"
	"fmt"
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
}