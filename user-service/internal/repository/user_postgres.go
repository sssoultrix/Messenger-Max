package repository

import (
	"context"
	"messenger-max/user-service/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Используем pgxpool для работы с бд
type UserPostgres struct {
	pool *pgxpool.Pool
}

// Метод для создания структуры UserPostgres с pgxpool внутри
func NewUserPostgres(pool *pgxpool.Pool) *UserPostgres {
	return &UserPostgres{pool: pool}
}

// Метод создания User в бд
// Для метода Create используется отдельная структура UserCreateRequest
func (u UserPostgres) Create(ctx context.Context, request domain.UserCreateRequest) error {
	//SQL запрос в бд с целью добавить User с переменными из request
	query := `INSERT INTO users (login, password_hash) VALUES ($1, $2)`

	//Метод Exec позволяет нам сделать SQL запрос к нашей бд
	_, err := u.pool.Exec(ctx, query, request.Login, request.Password)
	return err
}

// Метод обновления полей структуры User уже существующей в бд
// Для метода Update используется отдельная структура UserCreateRequest
func (u *UserPostgres) Update(ctx context.Context, request domain.UserCreateRequest) error {
	//SQL запрос для обновления логина и пароля у User по User.ID
	query := `UPDATE users SET login = $1, password_hash = $2 WHERE id = $3`
	//Метод Exec позволяет нам сделать SQL запрос к нашей бд
	_, err := u.pool.Exec(ctx, query, request.Login, request.Password)
	return err
}

// Метод для удаления уже существующего в бд User
func (u *UserPostgres) Delete(ctx context.Context, id int64) error {
	//SQL запрос для удаления User по User.ID
	query := `DELETE FROM users WHERE id = $1`
	//Метод Exec позволяет нам сделать SQL запрос к нашей бд
	_, err := u.pool.Exec(ctx, query, id)
	return err
}

// Метод для поиска по User.ID самого User в бд
func (u *UserPostgres) GetByID(ctx context.Context, id int64) (*domain.UserResponse, error) {
	//SQL запрос который достает login и name у User по User.ID
	query := `SELECT login, name FROM users WHERE id = $1`
	//C помощью QueryRow получаем одну строку из БД по запросу выше
	//QueryRow(ctx-КОНТЕКСТ, query-ТЕЛО ЗАПРОСА, id-ПАРАМЕТР ДЛЯ ЗАПРОСА)
	row := u.pool.QueryRow(ctx, query, id)
	//Пустышка, в которую мы поместим данные из запроса
	var user domain.UserResponse
	//С помощью row.Scan переносим данные из row в user-ПУСТЫШКА
	if err := row.Scan(&user.Login, &user.Name); err != nil {
		return nil, err
	}
	return &user, nil
}

// Метод возвращающий массив User из бд
func (u *UserPostgres) GetAll(ctx context.Context) ([]domain.UserResponse, error) {
	//SQL запрос, возвращающий все строки из ТАБЛИЦЫ users в нашей ДБ
	query := `SELECT id, login, name FROM users`
	//rows-Список строк из БД
	//Query возвращает ВСЕ строки подходящие по query запросу
	rows, err := u.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	//Отложено закрываем rows
	defer rows.Close()
	//Создаем пустышку для будущего слайса с User
	var users []domain.UserResponse
	//Итерируемся по списку пока он не закончится (rows.Next()-Возвращает bool)
	//"rows.NEXT() сам каждый раз как бы опускается по списку на нижнюю строчку и так до конца".
	for rows.Next() {
		//Создаем пустышку, в которую мы будем класть, при каждой итерации, новые данные из строчки в списке(rows)
		var user domain.UserResponse
		//Помещаем данные из строчки в user выше
		if err := rows.Scan(&user.ID, &user.Login, &user.Name); err != nil {
			return nil, err
		}
		//Добавляем заполненного user в слайс users
		users = append(users, user)
	}
	//возвращаем слайс с user
	return users, nil
}

// Метод для поиска по User.Login самого User в бд
// Такой же принцып как с GetByID
func (u *UserPostgres) GetByLogin(ctx context.Context, login string) (domain.UserResponse, error) {
	query := `SELECT id, login, name FROM users WHERE login = $1`
	row := u.pool.QueryRow(ctx, query, login)
	var user domain.UserResponse
	if err := row.Scan(&user.ID, &user.Login, &user.Name); err != nil {
		return domain.UserResponse{}, err
	}
	return user, nil
}
