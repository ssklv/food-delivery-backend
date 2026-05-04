package domain

type UpdateUserParams struct {
	ID      string
	Name    *string
	Phone   *string
	Email   *string
	Address *string
}
