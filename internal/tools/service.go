package tools

import "github.com/borisbbtest/ya-dr/internal/model"

func Equal(f *model.DataUsers, s *model.DataUsers) bool {
	return f.Login == s.Login && f.Password == s.Password
}
