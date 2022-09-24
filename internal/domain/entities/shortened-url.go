package entities

type ShortenedUrl struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	OriginalURL     string `json:"originalURL"`
	RecoveriesCount int    `json:"recoveriesCount"`
	CreateBy        string `json:"createBy"`
	CreateDate      int    `json:"createDate"`
	UpdateDate      int    `json:"updateDate"`
}
