package models

type Produtora struct {
	ID   int    `gorm:"column:id;primary key"`
	Nome string `gorm:"column:nome"`
	Username string `gorm:"column:username;unique"`
	Password string `gorm:"column:password"`
}

// TableName define o nome da tabela no banco de dados
func (p *Produtora) TableName() string {
	return "produtoras"
}
