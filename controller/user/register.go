package user

import (
	"context"
	"errors"

	"github.com/elotus_hackathon/model"
	"golang.org/x/crypto/bcrypt"
)

// Register create new user
func (i impl) Register(ctx context.Context, input model.User) (model.User, error) {
	// check if username exist in db
	existingUser, err := i.repo.User().GetUserByUsername(ctx, input.Username)
	if err != nil {
		return model.User{}, err
	}
	if existingUser.ID != 0 {
		return model.User{}, errors.New("user already exists")
	}

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	input.Password = string(hashedPassword)
	user, err := i.repo.User().InsertUser(ctx, input)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
