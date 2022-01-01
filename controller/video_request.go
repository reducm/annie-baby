package controller

type VideoInfoRO struct {
	Link           string `json:"link"`
	Filename       string `json:"filename binding:"-"`
	ExtractorProxy string `json:"extractor_proxy" binding:"-"`
	OutputDir      string `json:"output_dir"`
	ID             int    `json:"id" binding:"-"`
	Cookie         string `json:"cookie" binding:"-"` //目录
}
