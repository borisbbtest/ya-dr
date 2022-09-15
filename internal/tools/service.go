package tools

import "github.com/borisbbtest/ya-dr/internal/model"

func Equal(db *model.DataUser, r *model.DataUser) bool {
	return db.Login == r.Login && CheckPasswordHash(r.Password, db.Password)
}
