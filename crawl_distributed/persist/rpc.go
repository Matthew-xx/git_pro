package persist

import (
	"../../crawl_double/engine"
	"../../crawl_double/persist"
	"github.com/olivere/elastic"
)

type ItemServerService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemServerService) Save(item engine.Item,result *string) error {
	err := persist.Save(s.Client,s.Index,item)
	if err == nil {
		*result = "ok"
	}
	return err
}

