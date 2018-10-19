package es使用

import (
	"testing"
	"encoding/json"
	"fmt"
	"context"
	"reflect"
	"time"
	elastic "gopkg.in/olivere/elastic.v5"
)

func TestA(t *testing.T) {
	a()
}

var (
	index_log_order string = "logisitc_order"
	type_log_order  string = "order"
)

// LogOrder is a structure used for serializing/deserializing data in Elasticsearch.
type LogOrder struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Reorders int                   `json:"reorders"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

const mapping = `
{
"settings":{
"number_of_shards": 1,
"number_of_replicas": 0
},
"mappings":{
"order":{
 "properties":{
     "user":{
         "type":"keyword"
     },
     "message":{
         "type":"text",
         "store": true,
         "fielddata": true
     },
     "image":{
         "type":"keyword"
     },
     "created":{
         "type":"date"
     },
     "tags":{
         "type":"keyword"
     },
     "location":{
         "type":"geo_point"
     },
     "suggest_field":{
         "type":"completion"
     }
 }
}
}
}`

func a() {
	// Starting with elastic.v5, you must pass a context to execute each service
	ctx := context.Background()

	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course you can configure your client to connect
	// to other hosts and configure it in various other ways.
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
		panic(err)

	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)

	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)

	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(index_log_order).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)

	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex(index_log_order).BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)

		}
		if !createIndex.Acknowledged {
			// Not acknowledged

		}

	}

	// Index a order (using JSON serialization)
	order1 := LogOrder{User: "olivere", Message: "Take Five", Reorders: 0}
	put1, err := client.Index().
		Index(index_log_order).
		Type(type_log_order).
		Id("1").
		BodyJson(order1).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)

	}
	fmt.Printf("Indexed order %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	// Index a second order (by string)
	order2 := `{"user" : "olivere", "message" : "It's a Raggy Waltz"}`
	put2, err := client.Index().
		Index(index_log_order).
		Type(type_log_order).
		Id("2").
		BodyString(order2).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)

	}
	fmt.Printf("Indexed order %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)

	// Get order with specified ID
	get1, err := client.Get().
		Index(index_log_order).
		Type(type_log_order).
		Id("1").
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)

	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)

	}

	// Flush to make sure the documents got written.
	_, err = client.Flush().Index(index_log_order).Do(ctx)
	if err != nil {
		panic(err)

	}

	// Search with a term query
	termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := client.Search().
		Index(index_log_order). // search in index index_log_order
		Query(termQuery). // specify the query
		Sort("user", true). // sort by "user" field, ascending
		From(0).Size(10). // take documents 0-9
		Pretty(true). // pretty print request and response JSON
		Do(ctx) // execute
	if err != nil {
		// Handle error
		panic(err)

	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Each is a convenience function that iterates over hits in a search result.
	// It makes sure you don't need to check for nil values in the response.
	// However, it ignores errors in serialization. If you want full control
	// over iterating the hits, see below.
	var ttyp LogOrder
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(LogOrder); ok {
			fmt.Printf("LogOrder by %s: %s\n", t.User, t.Message)

		}

	}
	// TotalHits is another convenience function that works even when something goes wrong.
	fmt.Printf("Found a total of %d orders\n", searchResult.TotalHits())

	// Here's how you iterate through results with full control over each step.
	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d orders\n", searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a LogOrder (could also be just a map[string]interface{}).
			var t LogOrder
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				// Deserialization failed

			}

			// Work with order
			fmt.Printf("LogOrder by %s: %s\n", t.User, t.Message)

		}

	} else {
		// No hits
		fmt.Print("Found no orders\n")

	}

	// Update a order by the update API of Elasticsearch.
	// We just increment the number of reorders.
	update, err := client.Update().Index(index_log_order).Type(type_log_order).Id("1").
		Script(elastic.NewScriptInline("ctx._source.reorders += params.num").Lang("painless").Param("num", 1)).
		Upsert(map[string]interface{}{"reorders": 0}).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)

	}
	fmt.Printf("New version of order %q is now %d\n", update.Id, update.Version)

	// ...

	// Delete an index.
	deleteIndex, err := client.DeleteIndex(index_log_order).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)

	}
	if !deleteIndex.Acknowledged {
		// Not acknowledged

	}
}
