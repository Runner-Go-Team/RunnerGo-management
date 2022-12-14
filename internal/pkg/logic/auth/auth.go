package auth

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-omnibus/omnibus"

	"kp-management/internal/pkg/biz/consts"
	"kp-management/internal/pkg/dal"
	"kp-management/internal/pkg/dal/model"
	"kp-management/internal/pkg/dal/query"
	"kp-management/internal/pkg/dal/rao"
)

func SignUp(ctx context.Context, email, password, nickname string) (*model.User, error) {
	hashedPassword, err := omnibus.GenerateBcryptFromPassword(password)
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	user := model.User{Email: email, Password: hashedPassword, Nickname: nickname, Avatar: consts.DefaultAvatarMemo[rand.Intn(3)]}
	team := model.Team{Name: fmt.Sprintf("%s 的团队", nickname), Type: consts.TeamTypePrivate}

	err = query.Use(dal.DB()).Transaction(func(tx *query.Query) error {
		if err := tx.User.WithContext(ctx).Create(&user); err != nil {
			return err
		}

		team.CreatedUserID = user.ID
		if err := tx.Team.WithContext(ctx).Create(&team); err != nil {
			return err
		}

		return tx.UserTeam.WithContext(ctx).Create(&model.UserTeam{
			UserID: user.ID,
			TeamID: team.ID,
			RoleID: consts.RoleTypeOwner,
		})
	})

	if err != nil {
		return nil, err
	}

	SetUserSettings(ctx, user.ID, &rao.UserSettings{CurrentTeamID: team.ID})

	return &user, nil
}

func Login(ctx context.Context, email, password string) (*model.User, error) {
	tx := query.Use(dal.DB()).User
	user, err := tx.WithContext(ctx).Where(tx.Email.Eq(email)).First()
	if err != nil {
		return nil, err
	}

	if err := omnibus.CompareBcryptHashAndPassword(user.Password, password); err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateLoginTime(ctx context.Context, userID int64) error {
	tx := query.Use(dal.DB()).User
	_, err := tx.WithContext(ctx).Where(tx.ID.Eq(userID)).UpdateColumn(tx.LastLoginAt, time.Now())
	return err
}
