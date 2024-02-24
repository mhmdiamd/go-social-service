package google

import (
	"fmt"
	"net/smtp"
)

type SendEmailHeader struct {
  Sender_email string 
  Password string
  Recipent_email string
  Host_address string
  Host_port string
  Mail_subject string
  Mail_body string
}

type SendEmailHeaderParams struct {
  Sender_email string 
  Password string
  Recipent_email string
  Mail_subject string
  Mail_body string
}

func NewSendEmailHeader(header SendEmailHeaderParams) SendEmailHeader {
  return SendEmailHeader{
    Sender_email : header.Sender_email, 
    Password: header.Password,
    Recipent_email: header.Recipent_email,
    Host_address: "smtp.gmail.com",
    Host_port : "587",
    Mail_subject: header.Mail_subject, 
    Mail_body: header.Mail_body,
  }
} 

func (s *SendEmailHeader) SetHostAddress(hs string) {
  s.Host_address = hs
}

func (s *SendEmailHeader) SendGmailWithSMTP() (err error) {

  authenticate := smtp.PlainAuth("", s.Sender_email, s.Password, s.Host_address)
  fullServerAddress := s.Host_address + ":" + s.Host_port;

  err = smtp.SendMail(fullServerAddress, authenticate, s.Sender_email, []string{s.Recipent_email}, []byte(s.Mail_body))

  if err != nil {
    fmt.Println(err)
    return
  }

  return
}


// func (s *SendEmailHeader) SetConfigSMTP() (smtpClient *smtp.Client, err error) {
//   fullServerAddress := s.Host_address + ":" + s.Host_port;
//
//   headerMap := make(map[string]string)
//   headerMap["From"] = s.Sender_email
//   headerMap["To"] = s.Recipent_email
//   headerMap["Subject"] = s.Mail_subject
//
//   mailMessage := ""
//
//   for k, v := range headerMap {
//     mailMessage += fmt.Sprintf("%s: %s\\r", k, v)
//   }
//
//   mailMessage += "\\r" + s.Mail_body
//
//   // authenticate smtp
//   authenticate := smtp.PlainAuth("", s.Sender_email, s.Password, s.Host_address)
//   
//   tlsConfigurations := &tls.Config{
//     InsecureSkipVerify : true,
//     ServerName: s.Host_address,
//   }
//
//   conn, err := tls.Dial("tcp", fullServerAddress, tlsConfigurations)
//   if err != nil {
//     return
//   }
//
//   newClient, err := smtp.NewClient(conn, s.Host_address)
//   if err != nil {
//     return
//   }
//
//   // Auth
//   if err = newClient.Auth(authenticate); err != nil {
//     return
//   }
//
//   // To && Form
//   if err = newClient.Mail(s.Sender_email); err != nil {
//     return
//   }
//
//   if err = newClient.Rcpt(s.Recipent_email); err != nil {
//     return
//   }
//
//   return
// }

// func SendEmailYahooWithSMTP(smtpClient *smtp.Client, message string) (err error) {
//
//   writer, err := smtpClient.Data()
//   if err != nil {
//     return
//   }
//
//   _, err = writer.Write([]byte(message))
//   if err != nil {
//     return
//   }
//
//   err = writer.Close()
//   if err != nil {
//     return
//   }
//
//   err = smtpClient.Quit()
//   if err != nil {
//     fmt.Println("There was an error")
//   }
//
//   fmt.Println("successful, the mail was sent")
//
//   return
// }

