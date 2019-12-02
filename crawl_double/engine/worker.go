package engine

import (
	"log"
	"../fetcher"
)

func Worker(r Request) (ParseResult,error)  {
	log.Printf("fetching %s",r.Url)
	body ,err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetcher :error" + "fetching url %s: %v",r.Url,err)
		return ParseResult{},err
	}
	return r.Parser.Parse(body,r.Url),nil
}
