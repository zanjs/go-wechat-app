package schemas

import (
	"encoding/json"
	"time"
)

/*
err := config.LukaWechatDB.
		Select("picture_books.name, picture_books.isbn, scan_records.id, scan_records.status, scan_records.updated_at").
		Join("INNER", "scan_records", "scan_records.book_id = picture_books.id").
		Limit(size, page*size).
		Find(&scanRecordsPictureBooks)
*/

type ScanRecordPictureBook struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Isbn      string    `json:"isbn"`
	Status    int64     `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *ScanRecordPictureBook) MarshalJSON() ([]byte, error) {
	type Alias ScanRecordPictureBook
	return json.Marshal(&struct {
		*Alias
		Stamp string `json:"updated_at"`
	}{
		Alias: (*Alias)(s),
		Stamp: s.UpdatedAt.Format("2006-01-02"),
	})
}

func (ScanRecordPictureBook) TableName() string {
	return "picture_books"
}
