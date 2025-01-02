// Package main is a Cobra cli entrypoint.
package main

import (
	"bufio"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/feliux/autoclicker/internal/autoclick"
	"github.com/spf13/cobra"
	"golang.org/x/exp/rand"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "autoclick",
	Short: "AutoClick cli",
	Long:  `Automate your mouse pointer.`,
}

var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move the pointer",
	Long:  `Move the pointer to the current location within several seconds.`,
	Run: func(cmd *cobra.Command, args []string) {
		duration, _ := cmd.Flags().GetInt("duration")
		interval, _ := cmd.Flags().GetInt("interval")
		random, _ := cmd.Flags().GetBool("random")
		handler(duration, interval, 1, random)
	},
}

var clickCmd = &cobra.Command{
	Use:   "click",
	Short: "Click on the pointer",
	Long:  `Just click where the pointer is located every several seconds.`,
	Run: func(cmd *cobra.Command, args []string) {
		duration, _ := cmd.Flags().GetInt("duration")
		interval, _ := cmd.Flags().GetInt("interval")
		random, _ := cmd.Flags().GetBool("random")
		handler(duration, interval, 2, random)
	},
}

var moveAndClickCmd = &cobra.Command{
	Use:   "move-click",
	Short: "Move the pointer and click",
	Long:  `Move the pointer to the current location and click every few seconds.`,
	Run: func(cmd *cobra.Command, args []string) {
		duration, _ := cmd.Flags().GetInt("duration")
		interval, _ := cmd.Flags().GetInt("interval")
		random, _ := cmd.Flags().GetBool("random")
		handler(duration, interval, 3, random)
	},
}

func init() {
	rootCmd.PersistentFlags().IntP("duration", "d", 0, "Time in seconds for running the program.")
	rootCmd.PersistentFlags().IntP("interval", "i", 1, "Time between every action for the pointer.")
	rootCmd.PersistentFlags().BoolP("random", "r", false, "If true, the interval for every action will be a random value between 1 second and the value of interval flag.")
	rootCmd.AddCommand(moveCmd)
	rootCmd.AddCommand(clickCmd)
	rootCmd.AddCommand(moveAndClickCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		slog.Error("error occurred calling root command", "err", err)
		os.Exit(1)
	}
}

func handler(duration, interval, option int, random bool) {
	if ok := askForConfirmation("Move the pointer to a desired location and then confirm with [y/n]..."); ok {
		x, y := autoclick.GetLocation()
		tf := time.Now().Add(time.Second * time.Duration(duration))
		for tf.After(time.Now()) {
			if random {
				autoclick.Interval = 1 + rand.Intn(interval)
			} else {
				autoclick.Interval = interval
			}
			switch option {
			case 1:
				autoclick.Move(x, y)
			case 2:
				autoclick.Click()
			case 3:
				autoclick.MoveAndClick(x, y)
			}
		}
	}
}

// askForConfirmation asks the user for confirmation. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user.
func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		slog.Info(s)
		response, err := reader.ReadString('\n')
		if err != nil {
			slog.Error("error ocurred", "err", err)
		}
		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}
