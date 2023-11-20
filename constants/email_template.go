package constants

import "fmt"

func OTPVerificationUsingMail(name, verificationKey string) string {
	emailHeader := fmt.Sprintf("Hello " + name + ",\n")

	emailBody := fmt.Sprintf("Here is your key to verify your email of Axios CP portalP: " + verificationKey + "\n")

	emailFooter := fmt.Sprint("Regards,\nAXIOS CP WING")

	return fmt.Sprintf(emailHeader + "\n" + emailBody + "\n" + emailFooter)
}
