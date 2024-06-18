package main

import (
	stdctx "context"
	"fmt"

	"github.com/iamDushu/lenslocked/context"
	"github.com/iamDushu/lenslocked/models"
)

func main() {
	ctx := stdctx.Background()

	user := models.User {
		Email: "dushyanth@gmail.com",
	}
	ctx = context.WithUser(ctx, &user)

	retrivedUser := context.User(ctx)
	fmt.Println(retrivedUser.Email)
}
