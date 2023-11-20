package constants

import "fmt"

func OTPVerificationUsingMail(name, verificationKey string) string {
	emailHeader := fmt.Sprintf("Hello " + name + ",")

	emailBody := fmt.Sprintf("Here is your key to verify your email of Axios CP portalP: " + "*" + verificationKey + "*")

	emailFooter := fmt.Sprintf("Regards,\n AXIOS CP WING")

	return fmt.Sprintf(emailHeader + "\n" + emailBody + "\n" + emailFooter)
}
