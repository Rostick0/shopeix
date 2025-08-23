package pagination

import (
	"gorm.io/gorm"
)

type Paginator struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}

func (p *Paginator) Paginate(db *gorm.DB, out interface{}) (*Paginator, error) {
	var total int64
	db.Count(&total)

	if p.Page < 1 {
		p.Page = 1
	}
	if p.PerPage <= 0 {
		p.PerPage = 10
	}

	offset := (p.Page - 1) * p.PerPage
	err := db.Limit(p.PerPage).Offset(offset).Find(out).Error
	if err != nil {
		return nil, err
	}

	p.Total = total
	p.TotalPages = (total + int64(p.PerPage) - 1) / int64(p.PerPage)

	return p, nil
}
