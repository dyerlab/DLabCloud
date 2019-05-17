package manuscript

import (
	"github.com/jinzhu/gorm"
	// importing dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Article is the basic structure for holding research documents
type Article struct {
	ID       int
	Authors  string
	Year     int
	Title    string
	Journal  string
	Volume   int
	Pages    string
	Doi      string
	Abstract string
}

// Model handles all the DB related stuff
type Model struct {
	DB *gorm.DB
}

// AutoMigrate puts the Manuscript struct into the database if it is not alreday there
func (m *Model) AutoMigrate() {
	m.DB.AutoMigrate(Article{})

	var papers []Article
	m.DB.Find(&papers)
	if len(papers) == 0 {
		m.DB.Create(Article{
			Authors:  "Dyer RJ",
			Year:     2007,
			Title:    "The Evolution of Genetic Topologies",
			Journal:  "Theoretical Population Biology",
			Volume:   71,
			Pages:    "71-79",
			Doi:      "http://dx.doi.org/10.1016/j.tpb.2006.07.001",
			Abstract: "This manuscript explores the simultaneous evolution of population genetic parameters and topological features within a population graph through a series of Monte Carlo simulations. I show that node centrality and graph breadth are significantly correlated to population genetic parameters F<sub>ST</sub> and M (r 1⁄4 0:95; r 1⁄4 0:98, respectively), which are commonly used in quantifying among population genetic structure and isolation by distance. Next, the topological consequences of migration patterns are examined by contrasting N-island and stepping stone models of gene movement. Finally, I show how variation in migration rate influences the rate of formation of specific topological features with particular emphasis to the phase transition that occurs when populations begin to become fixed due to restricted movement of genes among populations. I close by discussing the utility of this method for the analysis of intraspecific genetic variation."})
	}
}

// Get returns the article with the given id
func (m *Model) Get(id int) (*Article, error) {
	var article Article
	res := m.DB.First(&article, id)
	if res.RecordNotFound() {
		return nil, m.DB.Error
	}
	return &article, nil
}

// List returns all the entries in the database
func (m *Model) List() ([]Article, error) {
	var ret []Article

	res := m.DB.Find(&ret)
	if res.RecordNotFound() {
		return nil, m.DB.Error
	}
	return ret, nil
}
