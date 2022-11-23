package db

import "github.com/yevishev/restaurant-customer/models"

func (db Database) GetAllDishes() (*models.DishList, error) {
	dishList := &models.DishList{}
	rows, err := db.Conn.Query("SELECT * FROM dishes")
	if err != nil {
		return dishList, err
	}
	for rows.Next() {
		var dish models.Dish
		err := rows.Scan(&dish.ID, &dish.Name, &dish.Price, &dish.CreatedAt, &dish.UpdatedAt)
		if err != nil {
			return dishList, err
		}
		dishList.Dishes = append(dishList.Dishes, dish)
	}
	return dishList, nil
}

func (db Database) CreateDish(dish *models.Dish) error {
	var id int
	query := "INSERT INTO dishes (name, price) VALUES ($1, $2) RETURNING id"
	err := db.Conn.QueryRow(query, &dish.Name, &dish.Price).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}