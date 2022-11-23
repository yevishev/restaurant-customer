package db

import (
    "database/sql"
    "github.com/yevishev/restaurant-customer/models"
)

func (db Database) GetAllOrders() (*models.OrderList, error) {
    orderList := &models.OrderList{}
    rows, err := db.Conn.Query("SELECT * FROM orders ORDER BY ID DESC")
    if err != nil {
        return orderList, err
    }

    for rows.Next() {
        var order models.Order
        err := rows.Scan(&order.ID, &order.UsersId, &order.Price, &order.PaymentStatus, &order.CreatedAt, &order.UpdatedAt)
        if err != nil {
            return orderList, err
        }
        orderList.Orders = append(orderList.Orders, order)
    }
    return orderList, nil
}

func (db Database) CreateOrder(order *models.Order) error {
    var id int
    var createdAt string
    query := "INSERT INTO orders (users_id, price) VALUES ($1, $2) RETURNING id, created_at"
    err := db.Conn.QueryRow(query, &order.UsersId, &order.Price).Scan(&id, &createdAt)
    if err != nil {
        return err
    }
    order.ID = id
    order.CreatedAt = createdAt
    return nil

}

func (db Database) GetOrderById(orderId int) (models.Order, error) {
    var order models.Order
    query := "SELECT * FROM orders WHERE id = $1"
    row := db.Conn.QueryRow(query, orderId)
    switch err := row.Scan(&order.ID, &order.UsersId, &order.Price, &order.PaymentStatus, &order.CreatedAt, &order.UpdatedAt); err{
        case sql.ErrNoRows:
            return order, ErrNoMatch
        default:
            return order, err
    }
}

func (db Database) DeleteOrder(orderId int) error {
    query := "DELETE FROM orders WHERE id = $1"
    _, err := db.Conn.Exec(query, orderId)
    switch err {
    case sql.ErrNoRows:
        return ErrNoMatch
    default:
        return err
    }
}

func (db Database) UpdateOrder(itemId int, orderData models.Order) (models.Order, error) {
    order := models.Order{}
    query := "UPDATE orders SET users_id = $1, price = $2, payment_status = $3, updated_at = NOW() WHERE id = $4 RETURNING id, updated_at"
    err := db.Conn.QueryRow(query, orderData.UsersId, orderData.Price, orderData.PaymentStatus, orderData.UpdatedAt).Scan(&order.ID, &order.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return order, ErrNoMatch
        }
        return order, err
    }
    return order, nil
}