package models

type Hacienda struct {
	CodHacienda   int    `gorm:"column:CodHacienda" json:"CodHacienda"`
	NomHacienda   string `gorm:"column:NomHacienda" json:"NomHacienda"`
	NomProductor  string `gorm:"column:NomProductor" json:"NomProductor"`
	ProvHacienda  string `gorm:"column:ProvHacienda" json:"ProvHacienda"`
	DirHacienda   string `gorm:"column:DirHacienda" json:"DirHacienda"`
	TelHacienda   string `gorm:"column:TelHacienda" json:"TelHacienda"`
	Tel1Hacienda  string `gorm:"column:Tel1Hacienda" json:"Tel1Hacienda"`
	EstHacAct     int8   `gorm:"column:EstHacAct" json:"EstHacAct"`
	PRODUCTOR_COD string `gorm:"column:PRODUCTOR_COD" json:"PRODUCTOR_COD"`
	TipoHac       int    `gorm:"column:TipoHac" json:"TipoHac"`
}

func (Hacienda) TableName() string {
	return "HACIENDA" // Nombre de la tabla con esquema
}
