package userstorage

import (
	"database/sql"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"user_server/gen/models"
	"user_server/gen/restapi/operations"
	"user_server/gen/restapi/operations/user"
)

type App struct {
	db *sql.DB
}

func NewApp(db *sql.DB) *App {
	return &App{
		db: db,
	}
}

func (a *App) ConfigureAPI(api *operations.UserStorageAPI) {
	api.UserPostUserHandler = user.PostUserHandlerFunc(a.PostUser)
	api.UserGetUserHandler = user.GetUserHandlerFunc(a.GetUser)
	api.UserPatchUserHandler = user.PatchUserHandlerFunc(a.PatchUser)
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(a.DeleteUser)
}

func (a *App) PostUser(params user.PostUserParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	instance := models.User{
		ID:        new(string),
		Name:      params.Body.Name,
		BirthDate: params.Body.BirthDate,
	}

	sqlQuery := `
		INSERT INTO "user" (name, birth_date) 
		VALUES ($1, $2) 
		RETURNING id
	`
	err := a.db.QueryRowContext(ctx, sqlQuery, instance.Name, instance.BirthDate.String()).Scan(instance.ID)
	if err != nil {
		return user.NewPostUserInternalServerError().WithPayload(toModelsError(err))
	}

	return user.NewPostUserCreated().WithPayload(&instance)
}

func (a *App) GetUser(params user.GetUserParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	instance := models.User{
		ID:        &params.ID,
		Name:      new(string),
		BirthDate: new(strfmt.Date),
	}
	var birthDateString string

	sqlQuery := `
		SELECT name, birth_date 
		FROM "user" 
		WHERE id = $1
	`
	err := a.db.QueryRowContext(ctx, sqlQuery, instance.ID).Scan(instance.Name, &birthDateString)
	if err == sql.ErrNoRows {
		return user.NewGetUserNotFound()
	}
	if err != nil {
		return user.NewGetUserInternalServerError().WithPayload(toModelsError(err))
	}

	err = instance.BirthDate.Scan(birthDateString)
	if err != nil {
		errorMessage := "birth_date has invalid format in the database"
		return user.NewGetUserInternalServerError().WithPayload(&models.Error{Error: &errorMessage})
	}

	return user.NewGetUserOK().WithPayload(&instance)
}

func (a *App) PatchUser(params user.PatchUserParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	instance := models.User{
		ID:        &params.ID,
		Name:      new(string),
		BirthDate: new(strfmt.Date),
	}
	var birthDateString string

	sqlQuery := `
		UPDATE "user" 
		SET name = COALESCE($1, name), birth_date = COALESCE($2, birth_date) 
		WHERE id = $3
		RETURNING name, birth_date
	`
	err := a.db.QueryRowContext(ctx, sqlQuery, params.Body.Name, params.Body.BirthDate, params.ID).Scan(instance.Name, &birthDateString)
	if err == sql.ErrNoRows {
		return user.NewPatchUserNotFound()
	}
	if err != nil {
		return user.NewPatchUserInternalServerError().WithPayload(toModelsError(err))
	}

	err = instance.BirthDate.Scan(birthDateString)
	if err != nil {
		errorMessage := "birth_date has invalid format in the database"
		return user.NewPatchUserInternalServerError().WithPayload(&models.Error{Error: &errorMessage})
	}

	return user.NewPatchUserOK().WithPayload(&instance)
}

func (a *App) DeleteUser(params user.DeleteUserParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	sqlQuery := `
		DELETE FROM "user" WHERE id = $1
	`
	result, err := a.db.ExecContext(ctx, sqlQuery, params.ID)
	if err != nil {
		return user.NewDeleteUserInternalServerError().WithPayload(toModelsError(err))
	}

	affectedRowCount, err := result.RowsAffected()
	if err != nil {
		return user.NewDeleteUserInternalServerError().WithPayload(toModelsError(err))
	}

	if affectedRowCount < 1 {
		return user.NewDeleteUserNotFound()
	}

	return user.NewDeleteUserOK()
}

func toModelsError(err error) *models.Error {
	errorMessage := err.Error()
	return &models.Error{
		Error: &errorMessage,
	}
}
