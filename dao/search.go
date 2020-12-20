package dao

import (
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"log"
	"reflect"
)

/*
** Pre-requisite of this code
* 1. Index will be created with following mapping
"mappings" : {
      "line" : {
        "properties" : {
          "line" : {
            "type" : "text",
            "fields" : {
              "keyword" : {
                "type" : "keyword",
                "ignore_above" : 256
              }
            }
          }
        }
      }
    }
  }
 */


type SearchDao struct{
	esClient *elasticsearch.Client
	index_name string
}

func (s *SearchDao)Init() error{
	var err error
	s.esClient, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	s.index_name = "test1" //TODO: remove hard-coded index name
	err = s.LoadIntoES("/home/ubuntu/search-serv/dao/completeworks.txt") //TODO:  remove hardcoded values
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func fetchLines(r map[string]interface{}) []string{
	var ListLine []string
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		a := hit.(map[string]interface{})["_source"]
		v := reflect.ValueOf(a)

		if v.Kind() == reflect.Map {
			for _, key := range v.MapKeys() {
				if(key.String() == "line"){
					line := fmt.Sprintf("%v", v.MapIndex(key).Interface())
					ListLine = append(ListLine, line)
				}
			}
		}
	}
	return ListLine
}



func (s *SearchDao) SearchPrefixBasedMatch(query_string string) []string {
	r := s.SearchPrefixBasedMatchFromES(query_string)
	return fetchLines(r)
}

