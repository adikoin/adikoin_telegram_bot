package main

import (
	"log"
	"os"
	"path/filepath"
	"telegram_bot/config"
	"telegram_bot/controller"
	"telegram_bot/repository"

	"github.com/NicoNex/echotron/v3"
	"github.com/joho/godotenv"
)

const (
	path_dir = "/home/dexter/go_projects/telegram_bot"
)

var userController *controller.UserController

func main() {

	err := godotenv.Load(filepath.Join(path_dir, ".env"))

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot_token := os.Getenv("TELEGRAM_APITOKEN")

	api := echotron.NewAPI(bot_token)

	for u := range echotron.PollingUpdates(bot_token) {
		userName := u.Message.From.Username
		welcomeMessage := "Hello, " + userName + "! In the dynamic world of business, customer retention and engagement are key. RewardNet is transforming loyalty programs by allowing customers to earn cashback and rewards in one business and spend them in another. Our innovative platform leverages blockchain technology, enabling businesses to issue their own NFT collections with exclusive discounts and benefits. But we don’t stop there. RewardNet also offers a launchpad for entrepreneurs, providing a powerful platform to kickstart and scale their businesses. By integrating a versatile rewards ecosystem with entrepreneurial support, RewardNet creates a seamless and interconnected business network. Join us in reshaping the future of loyalty programs and business growth with RewardNet, where rewards are limitless and opportunities are endless. Press the blue button in the menu to start change the world!"
		if u.Message.Text == "/start" {
			// user := u.Message.From
			err = userController.SaveUser(u.Message.From)
			if err != nil {
				api.SendMessage("Упс, что то пошло не так, попробуйте позднее", u.ChatID(), nil)
				log.Fatal(err)
			}

			// log.Println(u.ChatID())

			api.SendMessage(welcomeMessage, u.ChatID(), nil)
		}

		// if u.Message.Text == "/link" {
		// 	// user := u.Message.From
		// 	if err != nil {
		// 		api.SendMessage("Упс, что то пошло не так, попробуйте позднее", u.ChatID(), nil)
		// 		log.Fatal(err)
		// 	}

		// 	// log.Println(u.ChatID())
		// 	message := "https://192.168.88.2"
		// 	api.SendMessage(message, u.ChatID(), nil)
		// }
	}
}

func init() {
	mongoConnection, errorMongoConn := config.MongoConnection()

	if errorMongoConn != nil {
		log.Println("Error when connect mongo : ", errorMongoConn.Error())
	}

	userRepository := repository.NewUserRepository(mongoConnection)
	userController = controller.NewUserController(userRepository)

	// fileRepository := repository.NewFileRepository(mongoConnection)
	// fileController = controller.NewFileController(fileRepository)

	// postRepository := repository.NewPostRepository(mongoConnection)
	// postController = controller.NewPostController(postRepository)

	// ringRepository := repository.NewRingRepository(mongoConnection)
	// ringController = controller.NewRingController(ringRepository)

	// botRepository := repository.NewBotRepository(mongoConnection)
	// botController = controller.NewBotController(botRepository)
}
