package models

type Serie struct {
	ID              int    `gorm:"column:id;primary_key"`
	Titulo          string `gorm:"column:titulo"`
	DuracaoEpisodio int    `gorm:"column:duracao_episodio"`
	Categoria       string `gorm:"column:categoria"`
	FaixaEtaria     string `gorm:"column:faixa_etaria"`
	IDProdutora     int    `gorm:"column:id_produtora"`
	ImagemURL       string `gorm:"column:imagem_url"`
	Produtora       Produtora `gorm:"foreignkey:IDProdutora"`
}

// TableName define o nome da tabela no banco de dados
func (s *Serie) TableName() string {
	return "series"
}
