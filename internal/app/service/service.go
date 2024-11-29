package service

import (
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

// метод возвращает количество дней до 01.01.2025
func (s *Service) DaysLeft() int {
	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d)
	return int(dur.Hours() / 24)
	//s := fmt.Sprintf("Количество дней до 01.01.2025: %d", int(dur.Hours()/24))
}
