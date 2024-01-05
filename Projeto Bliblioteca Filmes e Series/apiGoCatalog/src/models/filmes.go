package models


type Filme struct {
	ID           int         `gorm:"column:id;primary_key"`  // Identificador único do filme
	Titulo       string      `gorm:"column:titulo"`          // Título do filme
	Duracao      int         `gorm:"column:duracao"`         // Duração do filme em minutos
	Categoria    string      `gorm:"column:categoria"`       // Categoria do filme (Ação, Comédia, etc.)
	FaixaEtaria  string      `gorm:"column:faixa_etaria"`    // Faixa etária do filme
	IDProdutora  int         `gorm:"column:id_produtora"`    // Chave estrangeira para a tabela de produtoras
	ImagemURL    string      `gorm:"column:imagem_url"`      // URL da imagem associada ao filme
	Produtora     Produtora   `gorm:"foreignkey:IDProdutora"` // Relacionamento com a produtora
}

// TableName define o nome da tabela no banco de dados
func (f *Filme) TableName() string {
	return "filmes"
}
