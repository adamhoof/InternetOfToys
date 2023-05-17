package database

import (
	"database/sql"
	"github.com/adamhoof/InternetOfToys/pkg/database"
	"github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(connString string) (*Postgres, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (p *Postgres) GetDevice(macAddress string) (database.Device, error) {
	var device database.Device
	query := `SELECT mac_address, functions FROM devices WHERE mac_address = $1`
	row := p.db.QueryRow(query, macAddress)
	err := row.Scan(&device.MacAddress, pq.Array(&device.Functions))
	if err != nil {
		return database.Device{}, err
	}
	return device, nil
}

func (p *Postgres) SaveDevice(device database.Device) error {
	query := `INSERT INTO devices (mac_address, functions) VALUES ($1, $2)`
	_, err := p.db.Exec(query, device.MacAddress, pq.Array(device.Functions))
	return err
}

func (p *Postgres) UpdateDevice(device database.Device) error {
	query := `UPDATE devices SET functions = $1 WHERE mac_address = $2`
	_, err := p.db.Exec(query, pq.Array(device.Functions), device.MacAddress)
	return err
}

func (p *Postgres) DeleteDevice(macAddress string) error {
	query := `DELETE FROM devices WHERE mac_address = $1`
	_, err := p.db.Exec(query, macAddress)
	return err
}

func (p *Postgres) GetAllDevices() ([]database.Device, error) {
	query := `SELECT mac_address, functions FROM devices`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []database.Device
	for rows.Next() {
		var device database.Device
		err = rows.Scan(&device.MacAddress, pq.Array(&device.Functions))
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return devices, nil
}
