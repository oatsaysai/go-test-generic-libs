package mariaDB

import (
	"database/sql"
	"time"
)

type SampleData struct {
	Name        string     `json:"name"`
	Data_001    string     `json:"data_001"`
	Data_002    string     `json:"data_002"`
	CreatedTime *time.Time `json:"created_time"`
	UpdatedTime *time.Time `json:"updated_time"`
}

func ListSampleData(db *sql.DB) ([]SampleData, error) {

	stmt, err := db.Prepare(
		`
		SELECT
			name,
			data_001,
			data_002,
			created_time,
			updated_time
		FROM sample_data
		`,
	)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]SampleData, 0)
	for rows.Next() {
		var row SampleData
		err = rows.Scan(
			&row.Name,
			&row.Data_001,
			&row.Data_002,
			&row.CreatedTime,
			&row.UpdatedTime,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, row)
	}

	return res, nil
}
