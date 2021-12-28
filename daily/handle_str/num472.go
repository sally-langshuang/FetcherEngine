package handle_str

import (
	"sort"
)

type trie struct {
	isEnd    bool
	children [26]*trie
}

func (root *trie) isConcat(word string) bool {
	if word == "" {
		return true
	}
	node := root
	for i, ch := range word {
		node = node.children[ch-'a']
		if node == nil {
			return false
		}
		if node.isEnd && root.isConcat(word[i+1:]) {
			return true
		}

	}
	return false
}

func (root *trie) insertToTrie(word string) {
	node := root
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}
func findAllConcatenatedWordsInADict(words []string) (ans []string) {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	root := &trie{}
	for _, w := range words {
		if w == "" {
			continue
		}
		if root.isConcat(w) {
			ans = append(ans, w)
		} else {
			root.insertToTrie(w)
		}
	}
	return ans
}

func (root *trie) insert(word string) {
	node := root
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (root *trie) dfs(vis []bool, word string) bool {
	if word == "" {
		return true
	}
	if vis[len(word)-1] {
		return false
	}
	vis[len(word)-1] = true
	node := root
	for i, ch := range word {
		node = node.children[ch-'a']
		if node == nil {
			return false
		}
		if node.isEnd && root.dfs(vis, word[i+1:]) {
			return true
		}
	}
	return false
}

func findAllConcatenatedWordsInADict0(words []string) (ans []string) {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	root := &trie{}
	for _, word := range words {
		if word == "" {
			continue
		}
		vis := make([]bool, len(word))
		if root.dfs(vis, word) {
			ans = append(ans, word)
		} else {
			root.insert(word)
		}
	}
	return
}
