package search

import (
	"fmt"
	"github.com/blevesearch/bleve/v2"
)

var (
	subject   Subject
	indexName = "subject"
	servers   = []string{"http://localhost:9200/"}
)


type Subject struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Genres []string `json:"genres"`
}

func e8()  {

}
//func e1()  {
//	const mapping = `
//{
//    "mappings": {
//        "properties": {
//            "id": {
//                "type": "long"
//            },
//            "title": {
//                "type": "text"
//            },
//            "genres": {
//                "type": "keyword"
//            }
//        }
//    }
//}`
//
//	ctx := context.Background()
//	client, err := elastic.NewClient(elastic.url(servers...))
//	if err != nil {
//		panic(err)
//	}
//
//	exists, err := client.IndexExists(indexName).Do(ctx)
//	if err != nil {
//		panic(err)
//	}
//	if !exists {
//		_, err := client.CreateIndex(indexName).BodyString(mapping).Do(ctx)
//		if err != nil {
//			panic(err)
//		}
//	}
//}
func x() {
	// open a new index
	//mapping := bleve.NewIndexMapping()
	//index, err := bleve.New("example.bleve1", mapping)
	index, err := bleve.Open("example.bleve1")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Name string
	}{
		Name: "text",
	}

	// index some data
	index.Index("id", data)

	// search for some text
	query := bleve.NewMatchQuery("text123")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
}
