package utils

import (
	"context"

	"github.com/TeamEvie/Backend/prisma/db"
	"github.com/fatih/color"
)

func GetEvieUserFromGHToken(token string) *db.UserModel {
	GitHubUser := GetGHUser(token)

	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		color.Red("[ERROR1] %s", err.Error())
		return nil
	}

	user, err := client.User.FindFirst(
		db.User.GithubID.Equals(GitHubUser.ID),
	).Exec(context.Background())

	if err == db.ErrNotFound {
		color.Yellow("User Doesn't Exist, creating...")

		if DoesUserSponsorTristan(token) {
			user, err = client.User.CreateOne(
				db.User.Name.Set(GitHubUser.Name),
				db.User.Email.Set(GitHubUser.Email),
				db.User.GithubID.Set(GitHubUser.ID),
			).Exec(context.Background())

			if err != nil {
				color.Red("[ERROR2] %s", err.Error())
				return nil
			}

			color.Green("[SUCCESS1] Created user %s", user.Name)
			return user
		} else {
			color.Red("[ERROR3] User not sponsored")
			return nil
		}
	} else if err != nil {
		color.Red("[ERROR4] %s", err.Error())
		return nil
	} else {
		color.Green("[SUCCESS2] Found user %s", user.Name)
		return user
	}
}
