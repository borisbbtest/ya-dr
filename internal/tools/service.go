package tools

import "github.com/borisbbtest/ya-dr/internal/model"

func Equal(f *model.DataUser, s *model.DataUser) bool {
	return f.Login == s.Login && f.Password == s.Password
}
