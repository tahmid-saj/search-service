package models

import "search-service/trie"

type SearchQueryInput struct {
	SearchQuery string `json:"searchQuery"`
	TableName   string `json:"tableName"`
}

type Response struct {
	Ok       bool `json:"ok"`
	Response interface{} `json:"response"`
}

func Search(searchQueryInput SearchQueryInput) (*Response, error) {
	searchResults, err := trie.Search(searchQueryInput.SearchQuery, searchQueryInput.TableName)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: searchResults,
	}, nil
}

func AddSearchQuery(searchQueryInput SearchQueryInput) (*Response, error) {
	addedSearchQuery, err := trie.AddSearchQuery(searchQueryInput.SearchQuery, searchQueryInput.TableName)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: addedSearchQuery,
	}, nil
}

func DeleteSearchQuery(searchQueryInput SearchQueryInput) (*Response, error) {
	deletedResult, err := trie.Search(searchQueryInput.SearchQuery, searchQueryInput.TableName)
	if err != nil {
		return &Response{
			Ok: false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: deletedResult,
	}, nil
}