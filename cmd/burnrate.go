package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	alertWindow    float64
	sloPeriod      float64
	budgetConsumed float64
	alertTime      string
	sloTime        string

	burnRateCmd = &cobra.Command{
		Use:   "burnrate",
		Short: "Calculates the burn rate of the error budget.",
		Long:  `Calculates the burn rate of the error budget to be used in alerts.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			CalculateErrorRate(alertWindow, sloPeriod, budgetConsumed)

			return nil
		},
	}
)

func init() {
	burnRateCmd.Flags().Float64VarP(&alertWindow, "alert-window", "a", 0, "The alert window time (required)")
	burnRateCmd.MarkFlagRequired("alert-window")
	burnRateCmd.Flags().Float64VarP(&sloPeriod, "slo-period", "s", 0, "The slo time. Default is in days (required)")
	burnRateCmd.MarkFlagRequired("slo-period")
	burnRateCmd.Flags().Float64VarP(&budgetConsumed, "budget-consumed", "b", 0, "The percentage of error budget consumed (required)")
	burnRateCmd.MarkFlagRequired("budget-consumed")

	burnRateCmd.Flags().StringVar(&alertTime, "alert-window-time", "", "The unit of time from alert window that should be used to calculate de burn rate. Default is hours.")
	burnRateCmd.Flags().StringVar(&sloTime, "slo-period-time", "", "The unit of time from slo period that should be used to calculate de burn rate. Default is days.")

}

// CalculateErrorRate is responsible for calculating the error rate
// receiving the alert window, slo period and budget consumed
func CalculateErrorRate(alertWindow float64, sloPeriod float64, budgetConsumed float64) {
	burnRate := ((budgetConsumed / 100) * (sloPeriod * 24)) / alertWindow

	fmt.Println(burnRate)
}
