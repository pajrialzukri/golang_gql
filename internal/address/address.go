package address

import (
	"backend/graph/model"
	db "backend/pkg/database"
	"log"
)

type Address model.Address

func (address *Address) Get(id *string) (*model.Address, error) {
	var ret model.Address
	var err error

	q := `
	SELECT
		id,
		line,
		province_id,
		city_id,
		district_id,
		village_id,
		postal_code
	FROM address WHERE id = ?
	`
	row := db.Handle.QueryRow(q, id)
	db.LogSQL(q, id)
	err = row.Scan(
		&ret.ID,
		&ret.Line,
		&ret.ProvinceID,
		&ret.CityID,
		&ret.DistrictID,
		&ret.VillageID,
		&ret.PostalCode,
	)

	if err != nil {
		log.Printf("Error select person: %v", err)
		return nil, err
	}

	return &ret, err
}
