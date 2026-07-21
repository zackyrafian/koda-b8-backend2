package repository

import (
	"context"
	"errors"
	"koda-b8-backend1/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)


type UserRepository struct {
    db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Create(req *domain.CreateUserRequest, ctx context.Context) (*domain.User, error) {
    tx, err := r.db.Begin(ctx) 
    if err != nil { 
      return nil, err 
    }
    defer tx.Rollback(ctx) 
    var UserID int64 

    err = tx.QueryRow(
      ctx, `
        INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id
      `, req.Email, req.Password,
    ).Scan(&UserID)

    _, err = tx.Exec(
      ctx, `
        INSERT INTO user_profiles (fullname, user_id) VALUES ($1, $2)
      `,req.FullName, UserID,
    )

    if err != nil { 
      return nil, err 
    }

    if err := tx.Commit(ctx); err != nil {
      return nil, err
    }

    return nil, err
    // var users domain.User
    // err := r.db.QueryRow(
    //   ctx, `
    //     INSERT INTO users (email, password) VALUES ($1, $2) RETURNING users.email, id
    //   `, req.Email, req.Password, 
    // ).Scan(&users.Email, &users.Id)
    // if err != nil { 
    //   return nil, err
    // }
    // return &users, nil
}

func (r *UserRepository) FindAll(ctx context.Context) (*[]domain.User, error) { 
  var users []domain.User
  rows, err := r.db.Query(
    ctx, `
    SELECT users.id as id, fullname, email FROM users LEFT JOIN user_profiles ON users.id = user_profiles.user_id
    `, 
  )

  if err != nil { 
    return &users, nil
  }
  for rows.Next() { 
    var p domain.User
    err = rows.Scan(&p.Id, &p.FullName, &p.Email)
    users = append(users, p)
  }
  return &users, nil
}

func (r *UserRepository) FindByEmail(email string, ctx context.Context) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRow(
	  ctx,`
			SELECT id, email, password FROM users WHERE email = $1
		`, email, 
	).Scan(&user.Id ,&user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user tidak ditemukan")
		}
		return nil, err
	}
	return user, err
}