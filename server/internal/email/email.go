package email

import (
	"context"
	"fmt"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mailgun/mailgun-go/v4"
)

func SendVerificationCodeEmail(userEmail, otpCode, userName string) error {
	apiKey := os.Getenv("MAIL_GUN_API_KEY")
	domainName := os.Getenv("MAIL_GUN_DOMAIN_NAME")
	if apiKey == "" || domainName == "" {
		return fmt.Errorf("MAIL_GUN_API_KEY or MAIL_GUN_DOMAIN_NAME is not set")
	}
	body := fmt.Sprintf(`
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Verification Code</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #f6f6f6;
				margin: 0;
				padding: 0;
			}
			.container {
				background-color: #ffffff;
				max-width: 600px;
				margin: 40px auto;
				padding: 20px;
				border-radius: 8px;
				box-shadow: 0 2px 6px rgba(0,0,0,0.1);
			}
			.header {
				text-align: center;
				font-size: 24px;
				color: #333333;
				margin-bottom: 20px;
			}
			.code-box {
				text-align: center;
				font-size: 32px;
				font-weight: bold;
				background-color: #f0f4ff;
				padding: 20px;
				border-radius: 8px;
				letter-spacing: 4px;
				color: #2b55d4;
				margin: 20px 0;
			}
			.footer {
				text-align: center;
				font-size: 14px;
				color: #888888;
				margin-top: 30px;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">Your Verification Code</div>
			<p>Hello,</p>
			<p>Use the following verification code to complete your sign-in or registration:</p>
			<div class="code-box">%s</div>
			<p>If you did not request this code, you can safely ignore this email.</p>
			<div class="footer">
				&copy; 2025 Openpaas.tech. All rights reserved.
			</div>
		</div>
	</body>
</html>
`, otpCode)
	mg := mailgun.NewMailgun(domainName, apiKey)
	mg.SetAPIBase(mailgun.APIBaseEU)
	message := mailgun.NewMessage(domainName, "Verification Code", "", userEmail)
	message.SetHTML(body)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, _, err := mg.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("error sending email: %s", err.Error())
	}
	return nil
}
