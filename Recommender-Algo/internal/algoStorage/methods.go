package algoStorage

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muesli/regommend"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func (c *Config) Init() {
	file, err := os.Open("./configs/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	d := yaml.NewDecoder(file)

	if err := d.Decode(&c); err != nil {
		log.Fatal(err)
	}
}

func (s *AlgoServer) Init(c Config, u ArtStorable, a *fiber.App) {
	s.ArtStore = u
	s.Cfg = c
	s.App = a
}

func (s *AlgoServer) AlgoSolve(in chan RecommendTuple) {
	bigMap, err := s.ArtStore.GetArtsForRecommend()
	if err != nil {
		log.Fatal(err)
	}

	artsTable := regommend.Table("arts")
	// name - имя пользователя, arts - map[artsName]userPoints
	for name, arts := range bigMap {
		artsVisited := make(map[interface{}]float64)
		for art, point := range arts {
			artsVisited[art] = point
		}
		// добавляем арты и оценки пользователя в таблицу под его именем
		artsTable.Add(name, artsVisited)
		//ClearMap(&artsVisited)
	}

	for t := range in {
		CityArts, err := s.ArtStore.GetArtsByCity(t.Req)
		if err != nil {
			t.Out <- nil
		}

		res := make(map[string]string)
		recs, err := artsTable.Recommend(t.Name)
		if err != nil {
			log.Println(err)
		}
		// проходим по всем рекомендованным артам
		for _, distPair := range recs {
			// проверяем, что арт из нужного города
			inTime, ok := CityArts[distPair.Key.(string)]
			if ok {
				// если есть - добавляем в результат
				res[distPair.Key.(string)] = inTime
			}
		}

		// если результат не пустой - отправляем его
		if len(res) > 0 {
			t.Out <- res
			// иначе отправляем список всех артов города
		} else {
			t.Out <- CityArts
		}

		close(t.Out)
	}
}

//func ClearMap(m *map[interface{}]float64) {
//	for key := range *m {
//		delete(*m, key)
//	}
//}
