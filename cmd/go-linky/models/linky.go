package models

type LinkyDataset struct {
	Label    string
	Data     string
	Time     string
	Checksum byte
}

// Enedis-NOI-CPT_54E- 6.1.1 : https://www.enedis.fr/media/2035/download
type TiC struct {
	ADCO     string `json:"adco,omitempty"`
	OPTARIF  string `json:"optarif,omitempty"`
	ISOUSC   uint8  `json:"isousc,omitempty"`
	BASE     uint32 `json:"base,omitempty"`
	HCHC     uint32 `json:"hchc,omitempty"`
	HCHP     uint32 `json:"hchp,omitempty"`
	EJPHN    uint32 `json:"ejphn,omitempty"`
	EJPHPM   uint32 `json:"ejbhpm,omitempty"`
	BBRHCJB  uint32 `json:"bbrhcjb,omitempty"`
	BBRHPJB  uint32 `json:"bbrhpjb,omitempty"`
	BBRHCJW  uint32 `json:"bbrhcjw,omitempty"`
	BBRHPJW  uint32 `json:"bbrhpjw,omitempty"`
	BBRHCJR  uint32 `json:"bbrhcjr,omitempty"`
	BBRHPJR  uint32 `json:"bbrhpjr,omitempty"`
	PEJP     uint8  `json:"pejp,omitempty"`
	PTEC     string `json:"ptec,omitempty"`
	DEMAIN   string `json:"demain,omitempty"`
	IINST    uint8  `json:"iinst"`
	ADPS     uint8  `json:"adps,omitempty"`
	IMAX     uint8  `json:"imax,omitempty"`
	PAPP     uint16 `json:"papp,omitempty"`
	HHPHC    string `json:"hhphc,omitempty"`
	MOTDETAT string `json:"motdetat,omitempty"`
}
