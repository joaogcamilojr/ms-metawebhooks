package webhook_data_entity

import (
	connection "ms-metawebhooks/connection"
)

func Create(SourceID string, Data string) (id int64, err error) {
  db, err := connection.Connect()

  if err != nil {
    panic(err)
  }

  defer db.Close()

  sql := `
    INSERT INTO webhooks_data
      (source_id, data) VALUES ($1, $2)
        RETURNING id;
  `

  err = db.QueryRow(sql, SourceID, Data).Scan(&id)

  return id, err
}
