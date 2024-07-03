package controller

import (
	model "telegram_bot/models"
	"telegram_bot/repository"

	"github.com/NicoNex/echotron/v3"
)

type UserController struct {
	userRepository repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) *UserController {
	return &UserController{userRepository: userRepository}
}

func (userController *UserController) SaveUser(user *echotron.User) error {

	_, err := userController.userRepository.FindByTelegramUserID(user.ID)
	if err == nil {
		return nil
		// return exception.ConflictException("User", "email", payload.Email)
	}

	newUser := &model.User{
		TelegramUserID: user.ID,
		IsBot:          user.IsBot,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Username:       user.Username,
		LanguageCode:   user.LanguageCode,
	}

	_, err = userController.userRepository.SaveUser(newUser)

	if err != nil {
		return err
	}

	return nil
}

// // func (userController *UserController) GetAllUser(c echo.Context) error {
// // 	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
// // 	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)

// // 	pagedUser, _ := userController.userRepository.GetAllUser(page, limit)
// // 	return util.Negotiate(c, http.StatusOK, pagedUser)
// // }

// func (userController *UserController) UpdateUser(c echo.Context) error {

// 	id := c.Param("id")

// 	payload := new(model.UserInputUpdate)

// 	if err := util.BindAndValidate(c, payload); err != nil {
// 		return err
// 	}

// 	currentUser, err := userController.userRepository.FindById(id)

// 	if err != nil || util.VerifyPassword(currentUser.Password, payload.CurrentPassword) != nil {
// 		return util.Negotiate(c, http.StatusUnauthorized, nil)
// 	}

// 	updateduser, err := userController.userRepository.UpdateUser(id, &model.UserUpdate{UserInputUpdate: payload})
// 	if err != nil {
// 		return util.Negotiate(c, http.StatusNotFound, updateduser)
// 	}

// 	return util.Negotiate(c, http.StatusOK, updateduser)
// }

// func (userController *UserController) DeleteUser(c echo.Context) error {
// 	id := c.Param("id")

// 	err := userController.userRepository.DeleteUser(id)
// 	if err != nil {
// 		return err
// 	}
// 	return c.NoContent(http.StatusNoContent)
// }

// func beforeSave(user *model.User) (err error) {
// 	hashedPassword, err := util.EncryptPassword(user.Password)
// 	if err != nil {
// 		return err
// 	}
// 	user.Password = string(hashedPassword)
// 	return nil
// }

// func (userController *UserController) HomePage(c echo.Context) error {
// 	// jwt, _ := c.Cookie("jwt")
// 	// log.Println("jwt is:", jwt)
// 	// c.Response().Header().Set("Authorization", "Bearer "+jwt.Value)
// 	// c.Response().WriteHeader(201)
// 	return c.Render(http.StatusOK, "dashboard", nil)

// }

// func (userController *UserController) ChangeUserPassword(c echo.Context) error {

// 	id := c.Param("id")

// 	payload := new(model.ChangeUserPassword)

// 	if err := util.BindAndValidate(c, payload); err != nil {
// 		return err
// 	}

// 	currentUser, err := userController.userRepository.FindById(id)

// 	if err != nil || util.VerifyPassword(currentUser.Password, payload.CurrentPassword) != nil {
// 		return util.Negotiate(c, http.StatusUnauthorized, nil)
// 	}

// 	hashedPassword, err := util.EncryptPassword(payload.Password)
// 	if err != nil {
// 		return err
// 	}
// 	payload.Password = string(hashedPassword)

// 	changeduserpassword, err := userController.userRepository.ChangeUserPassword(id, &model.UserChangePassword{ChangeUserPassword: payload})
// 	if err != nil {
// 		return util.Negotiate(c, http.StatusNotFound, changeduserpassword)
// 	}

// 	return util.Negotiate(c, http.StatusOK, changeduserpassword)
// }

// func (userController *UserController) GetDropdownUsers(c echo.Context) error {
// 	users, _ := userController.userRepository.GetDropdownUsers()
// 	return util.Negotiate(c, http.StatusOK, users)
// }
