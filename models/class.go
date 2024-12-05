package models

import "time"

type Class struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  int    `json:"capacity"`
}

func (c *Class) IsValid() bool {
	_, startErr := time.Parse("2006-01-02", c.StartDate)
	_, endErr := time.Parse("2006-01-02", c.EndDate)
	return c.Name != "" && startErr == nil && endErr == nil && c.Capacity > 0
}
