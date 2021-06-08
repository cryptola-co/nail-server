package cmd

import (
	"log"
	"os"
	"os/user"
	"runtime"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doctorCmd)
}

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check if the host system meets the requirements to run nail server",
	Run:   runDoctor,
}

func runDoctor(cmd *cobra.Command, args []string) {
	current, err := user.Current()
	if err != nil {
		log.Fatalf("unable to get current user: %s\n", err)
	}

	groups, err := current.GroupIds()
	if err != nil {
		log.Fatalf("unable to get current user groups: %s\n", err)
	}

	log.Println(groups)

	platformErrors := []error{}

	switch runtime.GOOS {
	case "freebsd":
		platformErrors = runDoctorFreeBSD()
	case "linux":
		platformErrors = runDoctorLinux()
	case "windows":
		platformErrors = runDoctorWindows()
	default:
		log.Fatalf("unsopported platform \"%s\". Please, create an issue at https://github.com/cryptola-co/nail-server to implement support.", runtime.GOOS)
	}

	if platformErrors != nil {
		log.Printf("Your host system (%s) does not meet the following requirements:\n", runtime.GOOS)
		for _, err := range platformErrors {
			log.Printf("\t - %s\n", err)
		}
		log.Println("Please correct them and try again.")
		os.Exit(1)
	}
}

func runDoctorWindows() []error {
	return nil
}

func runDoctorFreeBSD() []error {
	return nil
}

func runDoctorLinux() []error {
	return nil
}
