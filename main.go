package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	sendinblue "github.com/sendinblue/APIv3-go-library/lib"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var ctx context.Context
	cfg := sendinblue.NewConfiguration()

	fmt.Println("start test")
	//Configure API key authorization: api-key
	cfg.AddDefaultHeader("api-key", os.Getenv("sendinblue_api_key"))

	sender := &sendinblue.SendSmtpEmailSender{
		Email: os.Getenv("sendinblue_email_from"),
	}
	to := sendinblue.SendSmtpEmailTo{
		Email: "email_destiny",
	}
	replyTo := &sendinblue.SendSmtpEmailReplyTo{
		Email: os.Getenv("sendinblue_email_from"),
	}
	tag := []string{"test tag"}
	to2 := []sendinblue.SendSmtpEmailTo{to}

	body := sendinblue.SendSmtpEmail{
		Sender:      sender,
		To:          to2,
		Subject:     "test",
		TextContent: "test email",
		Tags:        tag,
		ReplyTo:     replyTo,
	}

	sib := sendinblue.NewAPIClient(cfg)
	result, resp, err := sib.TransactionalEmailsApi.SendTransacEmail(ctx, body)
	if err != nil {
		fmt.Println("Error when calling TransactionalEmailsApi->send_transac_email: ", err.Error())
		return
	}
	fmt.Println("SendTransacEmail Object:", result, " SendTransacEmail Response: ", resp)
	return
}
