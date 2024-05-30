package patient

import (
	"backend/graph/model"
	"backend/internal/address"
	db "backend/pkg/database"
	"log"
)

type Patient model.Patient

func (patient *Patient) Get(id *string) (*model.Patient, error) {
	var ret model.Patient
	var err error

	q := `
	SELECT
		id,
		current_address_id,
		name,
		birth_date,
		nik,
		religion
	FROM patients WHERE id = ?
	`
	row := db.Handle.QueryRow(q, id)
	db.LogSQL(q, id)
	err = row.Scan(
		&ret.ID,
		&ret.CurrentAddressID,
		&ret.Name,
		&ret.BirthDate,
		&ret.Nik,
		&ret.Religion,
	)

	if err != nil {
		log.Printf("Error select patient: %v", err)
		return nil, err
	}

	var address address.Address
	addressData, err := address.Get(&ret.CurrentAddressID)

	if err != nil {
		log.Printf("Error select address: %v", err)
		return nil, err
	}

	ret.CurrentAddress = append(ret.CurrentAddress, addressData)

	return &ret, nil
}
