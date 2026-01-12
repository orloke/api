package repositories

import (
	"api/src/models"
	"database/sql"
)

type Followers struct {
	db *sql.DB
}

func NewFollowersRepository(db *sql.DB) *Followers {
	return &Followers{db}
}

func (f Followers) Follow(userID, followerID uint64) error {
	statement, err := f.db.Prepare(
		"insert into followers (user_id, follower_id) values (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (f Followers) Unfollow(userID, followerID uint64) error {
	statement, err := f.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (f Followers) SearchFollowers(userID uint64) ([]models.UserResponse, error) {
	rows, err := f.db.Query(`
		select u.id, u.name, u.nick, u.email, u.created_at 
		from users u inner join followers s on u.id = s.follower_id 
		where s.user_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserResponse
	for rows.Next() {
		var user models.UserResponse
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (f Followers) SearchFollowing(userID uint64) ([]models.UserResponse, error) {
	rows, err := f.db.Query(`
		select u.id, u.name, u.nick, u.email, u.created_at 
		from users u inner join followers s on u.id = s.user_id 
		where s.follower_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserResponse
	for rows.Next() {
		var user models.UserResponse
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
