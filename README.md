# ShakeSearch

 In this repository,you'll find a simple web app that allows a user to search for a text string in
the complete works of Shakespeare.

Behavior of search result will be like below

    1. not case sensitive("Ham" and "ham" will result same)(Status : Done)
    2. Exclude any extra whitespaces("Ham " and "    ham"  will give same search result)(Status: Done)
    3. Prefex based and exact search (Status : Done) 
    4. suffix based search(TODO) 
    5. substring search (TODO, approach : using ngram tokenizer of elasticsearch feature)
    (Note : 3rd will have  more precedence followed by 4 and then 5. example : "ham" will first published result of)
    ham.* then .*ham and then .*ham.* )
    6. Only top 10 Result will be published instead of all the results(status: Done) . Todo : Result should be in paginated way(based on click next), example : top 10 results in first go 
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
