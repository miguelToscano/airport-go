package main

import (
	"fmt"
	"math/rand"
)

func main() {
	userStore := NewUserStore()

	fmt.Println("List of users:")
	for _, user := range userStore.users {
		fmt.Println(user)
	}

	fmt.Println("Find user by id:")
	user, err := userStore.Find(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

	fmt.Println("Find user by email:")
	user, err = userStore.FindByEmail("user2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

	fmt.Println("Find user by username:")
	user, err = userStore.FindByUsername("user3")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

	fmt.Println("Create user:")
	user, err = userStore.CreateUser(&CreateUserInput{
		Email:    "user4",
		Username: "user4",
	})
	if err != nil {
		fmt.Println(err)
	}

}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func NewUser(userInput *CreateUserInput) (*User, error) {
	var ID int = rand.Intn(100000000)

	return &User{
		ID:       ID,
		Email:    userInput.Email,
		Username: userInput.Username,
	}, nil
}

type UserStore interface {
	Find(id int) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	CreateUser(user *User) error
}

type CreateUserInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserStoreImpl struct {
	users []*User
}

func NewUserStore() *UserStoreImpl {
	return &UserStoreImpl{
		users: []*User{
			{ID: 1, Email: "",
				Username: "user1"},
			{ID: 2, Email: "	",
				Username: "user2"},
			{ID: 3, Email: "",
				Username: "user3"},
		},
	}
}

func (s *UserStoreImpl) Find(id int) (*User, error) {
	for _, u := range s.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, nil
}

func (s *UserStoreImpl) FindByEmail(email string) (*User, error) {
	for _, u := range s.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, nil
}

func (s *UserStoreImpl) FindByUsername(username string) (*User, error) {
	for _, u := range s.users {
		if u.Username == username {
			return u, nil
		}
	}
	return nil, nil
}

func (s *UserStoreImpl) CreateUser(userInput *CreateUserInput) (*User, error) {
	user, err := NewUser(userInput)

	if err != nil {
		return nil, err
	}

	return user, nil
}
