package google

import (
	"fmt"
	"testing"

	"github.com/mhmdiamd/go-social-service/internal/config"
	"github.com/stretchr/testify/require"
)

func init() {
  filename := "../../cmd/api/config.yaml"
  err := config.LoadConfig(filename)

  if err != nil {
    panic(err)
  }
}

func Test_GoogleEmailSMTP(t *testing.T) {
    t.Run("success connect smtp", func(t *testing.T) {

    header := SendEmailHeaderParams{
      Sender_email: config.Cfg.App.External.Google.Smtp_sender_email,
      Password: config.Cfg.App.External.Google.Smtp_password,
      Recipent_email: "milham0141@gmail.com",
      Mail_subject: "Sign up OTP",
      Mail_body: "Sign in with OTP",
    }

    newSmtpHeader := NewSendEmailHeader(header)
    err := newSmtpHeader.SendGmailWithSMTP()

    if err != nil {
      fmt.Println(err)
    }

    require.Nil(t, err)
  })

    t.Run("fail connect smtp", func(t *testing.T) {

    header := SendEmailHeaderParams{
      Sender_email: "worngemial@gmail.com",
      Password: config.Cfg.App.External.Google.Smtp_password,
      Recipent_email: "milham0141@gmail.com",
      Mail_subject: "Sign up OTP",
      Mail_body: "Sign in with OTP",
    }

    newSmtpHeader := NewSendEmailHeader(header)
    err := newSmtpHeader.SendGmailWithSMTP()

    if err != nil {
      fmt.Println(err)
    }

    require.NotNil(t, err)
  })
}

