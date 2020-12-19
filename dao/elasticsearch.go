package dao

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/esapi"
	"log"
	"os"
	"strconv"
	"strings"
)

func (s *SearchDao)LoadIntoES(filepath string) error{
	return nil
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line_id := 1

	// Reading line by line
	for scanner.Scan() {
		line := scanner.Text()

		var b strings.Builder
		b.WriteString(`{"line" : "`)
		b.WriteString(line)
		b.WriteString(`"}`)

		if(len(line) > 0){  //Skipping empty lines
			line_id = line_id + 1 // Generating auto-increament ID for the doc
			// Set up the request object.
			req := esapi.IndexRequest {
				Index:      s.index_name,
				DocumentType: s.doc_type,
				DocumentID: strconv.Itoa(line_id),
				Body:       strings.NewReader(b.String()),
				Refresh:    "true",
			}
			// Perform the request with the client.
			res, err := req.Do(context.Background(), s.esClient)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
				return err
			}
			defer res.Body.Close()
			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), line_id)
				return err
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}
	}
	return nil
}




func (s *SearchDao) SearchPrefixBasedMatchFromES(key string) map[string]interface{} {
	var r  map[string]interface{}
	// Perform the search request.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase_prefix": map[string]interface{}{
				"line": key,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := s.esClient.Search(
		s.esClient.Search.WithContext(context.Background()),
		s.esClient.Search.WithIndex(s.index_name),
		s.esClient.Search.WithDocumentType(s.doc_type),
		s.esClient.Search.WithBody(&buf),
		s.esClient.Search.WithTrackTotalHits(true),
		s.esClient.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	return r
}



