package dbdata

import (
	"evelp/log"
	"evelp/model"
	"io/ioutil"
	"sort"

	"gopkg.in/yaml.v2"
)

type corporationsData struct {
	filePath     string
	corporations *model.Corporations
}

func (cd *corporationsData) Refresh() error {
	log.Info("start load corporations", cd.filePath)
	if err := cd.load(); err != nil {
		return err
	}
	log.Info("load ", cd.filePath, " finished")

	log.Info("start save corporations to DB")
	if err := model.SaveCorporations(cd.corporations); err != nil {
		return err
	}
	log.Info("corporations have saved to DB")

	return nil
}

func (cd *corporationsData) load() error {
	file, err := ioutil.ReadFile(cd.filePath)
	if err != nil {
		return err
	}

	data := make(map[int]model.Corporation)
	if err := yaml.Unmarshal(file, &data); err != nil {
		return err
	}

	var corporations model.Corporations
	for id, corporation := range data {
		corporation.CorporationId = id
		corporations = append(corporations, corporation)
	}
	sort.Sort(corporations)
	cd.corporations = &corporations

	return nil
}
