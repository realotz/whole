package data

import (
	"github.com/realotz/whole/internal/apps/users/biz"
)

func (ar *memberRepo) CreateToken(member *biz.Member, expireTime int64) (string, error) {
	return ar.data.tk.Encode(member.Account, member.ID, member.Role, expireTime)
}
