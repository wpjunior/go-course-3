package main

import (
	"context"
	"fmt"
)

// seria um pacote user
type User struct {
	Email string
}

func WithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, "user", user)
}

func UserFromCtx(ctx context.Context) (user *User) {
	user, _ = ctx.Value("user").(*User)
	return user
}

// fim do pacote user

func main() {
	u := &User{Email: "wilsonpjunior@gmail.com"}
	ctx := context.TODO()
	ctx = WithUser(ctx, u)

	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	doSomething2(ctx)
}

func doSomething2(ctx context.Context) {
	u := UserFromCtx(ctx)
	fmt.Println("Usuário da transação", u)
}
