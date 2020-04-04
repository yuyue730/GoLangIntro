package persist

import (
	"GoLangIntro/SingleThreadCrawler/engine"
	"GoLangIntro/SingleThreadCrawler/persist"
	"log"

	"github.com/olivere/elastic/v7"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	if err == nil {
		log.Printf("Item %v saved.", item)
		*result = "okay"
	} else {
		log.Printf("Error saving item %v: %v.", item, err)
	}
	return err
}
