package user

import (
	"DDD/entities"
)

var _ IUserRepository = (*UserRepository)(nil)

type userTable struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Save(user entities.User) error {
	return nil
}

// Create の戻り値に error を追加してエラーを呼び出し元へ伝播する
func (r *UserRepository) Create(id int) (entities.User, error) {
	userData := r.fetchUserData(id)
	user, err := entities.NewUser(userData.ID, userData.Name)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

// get メソッドを fetchUserData に名称変更
func (r *UserRepository) fetchUserData(id int) userTable {
    // 仮のデータ取得処理
	return userTable{
		ID:   id,
		Name: "getJohn",
	}
}