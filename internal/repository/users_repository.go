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
    SELECT users.id as id, fullname, email, picture FROM users LEFT JOIN user_profiles ON users.id = user_profiles.user_id
    `, 
  )

  if err != nil { 
    return &users, nil
  }
  for rows.Next() { 
    var p domain.User
    err = rows.Scan(&p.Id, &p.FullName, &p.Email, &p.Picture)
    users = append(users, p)
  }
  return &users, nil
}

func (r *UserRepository) FindByEmail(email string, ctx context.Context) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRow(
	  ctx,`
			SELECT users.id as id, email, password, fullname, profile FROM users
		  LEFT JOIN user_profiles ON users.id = user_profiles.user_id
			WHERE email = $1
		`, email, 
	).Scan(&user.Id ,&user.Email, &user.Password, &user.FullName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user tidak ditemukan")
		}
		return nil, err
	}
	return user, err
}

func (r *UserRepository) FindByID(id int64, ctx context.Context) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRow(
	  ctx,`
			SELECT id, email FROM users WHERE id = $1
		`, id, 
	).Scan(&user.Id ,&user.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user tidak ditemukan")
		}
		return nil, err
	}
	return user, err
}

func (r *UserRepository) Delete(id int64, ctx context.Context) (error) {
	err := r.db.QueryRow(
	  ctx,`
			DELETE FROM users WHERE id = $1
		`, id, 
	)
	if err != nil { 
	  return nil
	}
	return nil
}

func (r *UserRepository) Patch(id int64, req *domain.PatchUserRequest, ctx context.Context) (*domain.User, error) {
  tx, err := r.db.Begin(ctx) 
  if err != nil { 
    return nil, err
  }
  defer tx.Rollback(ctx)
  user := &domain.User{}

  err = tx.QueryRow(
    ctx, `
      UPDATE users
      SET email = COALESCE($1, email)
      WHERE id = $2
      RETURNING id, email 
    `, req.Email, id, 
  ).Scan(&user.Id, &user.Email)

  if err != nil { 
    return nil, err 
  }

  _, err = tx.Exec(
    ctx, `
      UPDATE user_profiles 
      SET fullname = COALESCE($1, fullname)
      WHERE user_id = $2
    `, req.FullName, id,
  )
  if err != nil { 
    return nil, err
  }
  if err := tx.Commit(ctx); err != nil { 
    return nil, err
  }
  user.FullName = *req.FullName
  return user, nil
}

func (r *UserRepository) UploadPictureProfile(id int64, req *domain.UploadPicturesProfileRequest, ctx context.Context) (*domain.User, error) { 
    var user domain.User
    err := r.db.QueryRow(
      ctx,`
        UPDATE user_profiles
        SET picture = COALESCE($2, picture)
        WHERE id = $1 
        RETURNING id, picture
      `, id, req.Picture,
    ).Scan(&user.Id, &user.Picture)
    if err != nil {
      return nil, err
    }
    return &user , err 
}