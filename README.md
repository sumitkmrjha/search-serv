# ShakeSearch

 In this repository,you'll find a simple web app that allows a user to search for a text string in
the complete works of Shakespeare.

In it's current state, The search is

    1. not case sensitive
    2. Exclude any extra whitespaces
    3. Prefex based and exact search (Note : this result will have more precedence and will come before than results of 4 and 5 )
    4. suffix based search(TODO) 
    5. substring search (TODO)
    6. Result will be in paginated way(based on click next), example : top 10 results in first go 
    Note : Result of 3,4 and 5 is not limited to search for one word, user can type multiple keywords and search results     
    will be ranked to its priority order . In layman term, consider searching over google search engine. 

### How to setup
* Setup Elasticsearch
* Create an index with following mapping 

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
           
* Run main.go file which will 
       
       
       1. Load entire book in to elasticsearch index
       2. run a server on port 3001
       
* query results 

            http://54.179.4.29:3001/search?q=<your query string>   
