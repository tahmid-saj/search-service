package trie

import (
	"search-service/dynamodb"
	"search-service/utils"
)

type SearchResult struct {
	Result          string
	ResultFrequency int
}

func Search(searchQuery string, tableName string) ([]SearchResult, error) {
	// ReadItem first from the table
	itemReturned, err := dynamodb.ReadItem(searchQuery, tableName)
	if err != nil {
		AddSearchQuery(searchQuery, tableName)
		return nil, err
	}

	// if the item does exist, update the table's frequent queries, child nodes, leaf node, etc and return the search result
	UpdateTrieNodes(searchQuery, tableName)
	
	searchResults := []SearchResult{}
	for searchResult, searchResultFrequency := range itemReturned.FrequentQueries {
		searchResults = append(searchResults, SearchResult{
			Result: searchResult,
			ResultFrequency: searchResultFrequency,
		})
	}

	return searchResults, nil

	// otherwise, update the trieNode's FrequentQueries (maintain only the top 5 most searched searchQuery)
}

func AddSearchQuery(searchQuery string, tableName string) (bool, error) {
	// loop through each prefix of the searchQuery:
	// 		check if the prefix exists by using ReadItem
	// 				if the prefix doesn't exist, create it using AddItem (add the Prefix, FrequentQueries, ChildNodes to the next Prefix and set LeafNode accordingly)
	// 				if the prefix exists, update the current TrieNode (itemRead) (update it's FrequentQueries by adding the searchQuery to the FrequentQueries if the length is < 5, and leave it's Childnodes, LeafNode unchanged)

	for prefixIndex := range searchQuery {
		// check if the prefix exists by using ReadItem
		currentPrefix := searchQuery[:prefixIndex + 1]

		readItem, err := dynamodb.ReadItem(currentPrefix, tableName)

		if err != nil {
			// if the prefix doesn't exist, create it using AddItem (add the Prefix, FrequentQueries, ChildNodes to the next Prefix and set LeafNode accordingly)
			
			var childNodes []string
			var leafNode bool

			if prefixIndex == len(searchQuery) - 1 {
				childNodes = []string{}
				leafNode = true
			} else if prefixIndex < len(searchQuery) - 1 {
				nextPrefix := searchQuery[:prefixIndex + 2]
				childNodes = []string{nextPrefix}
				leafNode = false
			}

			_, err := dynamodb.AddItem(dynamodb.TrieNode{
				Prefix: currentPrefix,
				FrequentQueries: map[string]int{
					searchQuery: 1,
				},
				ChildNodes: childNodes,
				LeafNode: leafNode,
			}, tableName)
			if err != nil {
				return false, err
			}
		}

		// read the item again to get the updated item after creation
		readItem, err = dynamodb.ReadItem(currentPrefix, tableName)
		if err != nil {
			return false, err
		}

		// if the prefix exists, update the current TrieNode (itemRead) (update it's FrequentQueries by adding the searchQuery to the FrequentQueries if the length is < 5, and leave it's Childnodes, LeafNode unchanged)
		var updatedTrieNode dynamodb.TrieNode
		updatedFrequentQueries := updateFrequentQueries(readItem.FrequentQueries, searchQuery)

		updatedTrieNode = dynamodb.TrieNode{
			Prefix: currentPrefix,
			FrequentQueries: updatedFrequentQueries,
			ChildNodes: readItem.ChildNodes,
			LeafNode: readItem.LeafNode,
		}

		_, err = dynamodb.UpdateItem(updatedTrieNode, tableName)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func UpdateTrieNodes(searchQuery string, tableName string) (bool, error) {
	// loop through each prefix of the searchQuery:
	// 		check if the prefix exists by using ReadItem
	// 				if the prefix doesn't exist, create it using AddItem (add the Prefix, FrequentQueries, ChildNodes to the next Prefix and set LeafNode accordingly)
	// 				if the prefix exists, update the current TrieNode (itemRead) (update it's FrequentQueries by adding the searchQuery to the FrequentQueries if the length is < 5, and leave it's Childnodes, LeafNode unchanged)

	resUpdatedTrieNodes, err := AddSearchQuery(searchQuery, tableName)
	if err != nil {
		return false, err
	}

	return resUpdatedTrieNodes, nil
}

func DeleteTrieNode(searchQuery string, tableName string) (bool, error) {
	_, err := dynamodb.DeleteItem(searchQuery, tableName)
	if err != nil {
		return false, err
	}

	updatedParentNodesAfterDeletedTrieNode, err := updateParentNodesAfterDeletedTrieNode(searchQuery, tableName)
	if err != nil {
		return false, err
	}

	return updatedParentNodesAfterDeletedTrieNode, nil
}


// helper functions

func updateFrequentQueries(frequentQueries map[string]int, searchQuery string) map[string]int {
	// if searchQuery exists in frequentQueries, increase it's frequency by 1 and return the updated frequentQueries
	_, exists := frequentQueries[searchQuery]
	if exists {
		frequentQueries[searchQuery] = frequentQueries[searchQuery] + 1
		return frequentQueries
	}

	// if the length of frequentQueries is < 5, add searchQuery with a frequency of 1 to frequentQueries and return it
	if len(frequentQueries) < 5 {
		frequentQueries[searchQuery] = 1
		return frequentQueries
	}
	
	// if the length is 5, find the smallest frequency value:
	// 		if the frequency value is == 1, replace it with searchQuery, otherwise do nothing
	for result, frequency := range frequentQueries {
		if frequency == 1 {
			delete(frequentQueries, result)
			frequentQueries[searchQuery] = 1
		}
	}
	
	return frequentQueries
}

func updateParentNodesAfterDeletedTrieNode(searchQuery string, tableName string) (bool, error) {
	// loop through each prefix of the searchQuery:
	// 		check if the prefix exists by using ReadItem
	// 			if the prefix doesn't exist, return an error
	// 			if the prefix exists, remove the searchQuery from the FrequentQueries map of the current TrieNode (itemRead), 
	// 			and if the current TrieNode has the searchQuery in it's ChildNodes, remove the searchQuery from the ChildNodes too

	for prefixIndex := range searchQuery {
		// check if the prefix exists by using ReadItem
		currentPrefix := searchQuery[:prefixIndex + 1]

		readItem, err := dynamodb.ReadItem(currentPrefix, tableName)
		if err != nil {
			// if the prefix doesn't exist, return an error
			return false, err
		}

		// if the prefix exists, remove the searchQuery from the FrequentQueries map of the current TrieNode (itemRead), 
		// and if the current TrieNode has the searchQuery in it's ChildNodes, remove the searchQuery from the ChildNodes too
		// if after removing the searchQuery from the ChildNodes, the ChildNodes' length is 0, then set the current TrieNode as the leaf node (LeafNode = true)

		updatedFrequentQueries := deleteFrequentQueriesItem(readItem.FrequentQueries, searchQuery)
		updatedChildNodes := deleteChildNode(readItem.ChildNodes, searchQuery)
		 
		leafNode := false
		if len(updatedChildNodes) == 0 {
			leafNode = true
		}

		updatedTrieNode := dynamodb.TrieNode{
			Prefix: currentPrefix,
			FrequentQueries: updatedFrequentQueries,
			ChildNodes: updatedChildNodes,
			LeafNode: leafNode,
		}

		_, err = dynamodb.UpdateItem(updatedTrieNode, tableName)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func deleteFrequentQueriesItem(frequentQueries map[string]int, searchQuery string) map[string]int {
	_, exists := frequentQueries[searchQuery]
	if exists {
		delete(frequentQueries, searchQuery)
	}

	return frequentQueries
}

func deleteChildNode(childNodes []string, searchQuery string) []string {
	for childNodeIndex, childNode := range childNodes {
		if childNode == searchQuery {
			childNodes = utils.DeleteAtIndexSliceString(childNodes, childNodeIndex)
		}
	}

	return childNodes
}
