package dictionary

type DictionaryResponse struct {
	Result *Dictionary `json:"result"`
}

type Dictionary struct {
	Streets   map[string]string `json:"ulice"`
	StopTypes map[string]string `json:"typy_przystankow"`
	Locations map[string]string `json:"miejsca"`
}
