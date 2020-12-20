package searcher

import "search-serv/dao"

type Searcher struct {
	searchDao *dao.SearchDao
}


func (s *Searcher)Init(){
	s.searchDao = new(dao.SearchDao)
	s.searchDao.Init()
}

/*


 */

func (s *Searcher)Search(query string) []string {
	res :=  s.searchDao.SearchPrefixBasedMatch(query)
	/*
	If given more time
	Further if the size is less than 10 then match suffix based search and append in the result set
	If the result size is still small then try substring search either using regexp or tokenizer/ngram
	 */
	return res
}
