package data

import "time"

type TestStore struct {
	TestCreateUser             func(*User) error
	CreateUserCallCount        int
	TestGetUser                func(string) (*User, error)
	GetUserCallCount           int
	TestCreateAccessToken      func(*User) (*AccessToken, error)
	CreateAccessTokenCallCount int
	TestDeleteAccessToken      func(string) error
	DeleteAccessTokenCallCount int
	TestGetAccessToken         func(string) (*AccessToken, error)
	GetAccessTokenCallCount    int
}

func (t *TestStore) CreateUser(user *User) error {
	t.CreateUserCallCount += 1
	return t.TestCreateUser(user)
}

func (t *TestStore) GetUser(username string) (*User, error) {
	t.GetUserCallCount += 1
	return t.TestGetUser(username)
}

func (t *TestStore) CreateAccessToken(user *User) (*AccessToken, error) {
	t.CreateAccessTokenCallCount += 1
	return t.TestCreateAccessToken(user)
}

func (t *TestStore) DeleteAccessToken(id string) error {
	t.DeleteAccessTokenCallCount += 1
	return t.TestDeleteAccessToken(id)
}

func (t *TestStore) GetAccessToken(id string) (*AccessToken, error) {
	t.GetAccessTokenCallCount += 1
	return t.TestGetAccessToken(id)
}

func (t *TestStore) GetAccessTokenByQueryToken(string) (*AccessToken, error) {
	panic("Not implemented")
}

func (t *TestStore) CreateQuery(query *Query) error {
	panic("Not implemented")

}
func (t *TestStore) GetQueryCountForUser(id string, since *time.Time) (int, error) {
	panic("Not implemented")
}
