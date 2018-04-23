package schemas

import (
	"encoding/json"
	"time"
)

type ScanRecords struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Isbn      string    `json:"isbn"`
	Status    int64     `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *ScanRecords) MarshalJSON() ([]byte, error) {
	type Alias ScanRecords
	return json.Marshal(&struct {
		*Alias
		Stamp string `json:"updated_at"`
	}{
		Alias: (*Alias)(s),
		Stamp: s.UpdatedAt.Format("2006-01-02"),
	})
}

func (ScanRecords) TableName() string {
	return "scan_records"
}
