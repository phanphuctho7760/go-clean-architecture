package user

import (
	"github.com/phanphuctho7760/go-clean-architecture/app/entities"
)

// User This is for struct only
// or some function constraint data in user, but this is not recommend
type User struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
}

func convertUserRepoToEntity(mysqlUser User) (entitiesUser entities.User) {
	entitiesUser.Id = mysqlUser.Id
	entitiesUser.UserName = mysqlUser.UserName
	return
}

func convertUsersRepoToEntities(usersMysql []User) (entitiesUsers []entities.User) {
	for _, mUser := range usersMysql {
		eUser := convertUserRepoToEntity(mUser)
		entitiesUsers = append(entitiesUsers, eUser)
	}
	return
}
