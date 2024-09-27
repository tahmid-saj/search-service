package trie

type SearchResult struct {
	Result string
	ResultFrequency int
}

type TrieNode struct {
	Prefix string
	FrequentQueries map[string]int
	ChildNodes []string
	LeafNode bool
}

func Search(searchQuery string) ([]SearchResult, error) {

}

func AddTrieNode(searchQuery string) (TrieNode, error) {

}

func UpdateTrieNodes(rootTrieNode TrieNode) (TrieNode, error) {

}