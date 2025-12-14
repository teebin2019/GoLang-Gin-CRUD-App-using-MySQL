package repositories

import (
	"database/sql"

	"GoCRUDApplicationMySQL/app/models"
)

type VegetableRepository struct {
	db *sql.DB
}

func NewVegetableRepository(db *sql.DB) *VegetableRepository {
	return &VegetableRepository{
		db: db,
	}
}

func (r *VegetableRepository) GetAllVegetables() ([]models.Vegetable, error) {
	query := "SELECT * FROM vegetables"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vegetables := []models.Vegetable{}
	for rows.Next() {
		vegetable := models.Vegetable{}
		err := rows.Scan(&vegetable.ID, &vegetable.Name, &vegetable.Price)
		if err != nil {
			return nil, err
		}
		vegetables = append(vegetables, vegetable)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return vegetables, nil
}

func (r *VegetableRepository) CreateVegetable(vegetable *models.Vegetable) error {
	query := "INSERT INTO vegetables (name, price) VALUES (?, ?)"
	result, err := r.db.Exec(query, vegetable.Name, vegetable.Price)
	if err != nil {
		return err
	}

	vegetableID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	vegetable.ID = int(vegetableID)
	return nil
}

// Implement other database operations like GetUser, UpdateUser, and DeleteUser
