package bots

import (
	"errors"
	"fmt"
	tgapi "github.com/Syfaro/telegram-bot-api"
	"regexp"
	"server/database"
	"server/models"
	"strings"
	"time"
)

type TelegramBot struct {
	bot        *tgapi.BotAPI
	repository database.ReviewerRepository
	password   string
	quit       chan bool
}

func NewTelegramBot(repository database.ReviewerRepository, token string, password string) (*TelegramBot, error) {
	if len(token) == 0 {
		return nil, errors.New("TelegramBot token can not be empty")
	}

	if len(password) == 0 {
		return nil, errors.New("TelegramBot password can not be empty")
	}

	bot, err := tgapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TelegramBot{
		bot:        bot,
		password:   password,
		repository: repository,
		quit:       make(chan bool),
	}, nil
}

func (this *TelegramBot) ServeEvents() {
	update := tgapi.NewUpdate(0)
	update.Timeout = 60

	updates, err := this.bot.GetUpdatesChan(update)
	if err != nil {
		fmt.Println(err)
		return
	}

	for event := range updates {
		select {
		case <-this.quit:
			return
		default:
		}

		if event.Message == nil {
			continue
		}

		text := event.Message.Text

		if len(text) == 0 {
			continue
		}

		matched, err := regexp.MatchString("/connect(.*)", text)
		if err != nil {
			fmt.Println(err)
			this.bot.Send(tgapi.NewMessage(event.Message.Chat.ID, "Шо ты делаешь, мне плохо :("))
			continue
		}
		if !matched {
			this.bot.Send(tgapi.NewMessage(event.Message.Chat.ID, "Неа, не поняль"))
			continue
		}

		var responseText string
		userPassword := text[strings.LastIndex(text, " ")+1:]
		if userPassword == this.password {
			err = this.repository.Add(models.NewTelegramReviewer(event.Message.Chat.UserName, event.Message.Chat.ID))
			if err == nil {
				responseText = "Хм, кажется, Вам можно доверять. Теперь Вы будете получать отчеты в этот чат"
			} else {
				fmt.Println(err)
				responseText = "Хм, кажется, Вам можно доверять, но при подключении чата что-то пошло не так :("
			}
		} else {
			responseText = "Я тебе не доверяю -_-"
		}

		this.bot.Send(tgapi.NewMessage(event.Message.Chat.ID, responseText))
	}
}

func (this *TelegramBot) SendReport(report models.Report) {
	reviewers, err := this.repository.All(models.TelegramInfoMapper)
	if err != nil {
		fmt.Println(err)
		return
	}

	msk, _ := time.LoadLocation("Europe/Moscow")
	userTime := report.SendTime.In(msk).Format("2006-01-02 15:04:05")
	serverTime := report.ServedTime.In(msk).Format("2006-01-02 15:04:05")

	message := fmt.Sprintf("Отчет от %s (IP %s) по приложению %s\nВремя формирования: %s\nВремя получения: %s\nСообщение: %s",
		report.Author.Name, report.Author.IP, report.AppName, userTime, serverTime, report.Message)

	if report.Error != nil {
		message += fmt.Sprintf("\nКласс: %s\nМетод: %s\nСтрока: %s\nСтек вызовов: %s",
			report.Error.Class, report.Error.Method, report.Error.Line, report.Error.StackTrace)
	}

	for _, reviewer := range reviewers {
		info := reviewer.Info.(models.TelegramInfo)
		this.bot.Send(tgapi.NewMessage(info.ChatId, message))
	}
}

func (this *TelegramBot) Quit() {
	this.quit <- true
}
