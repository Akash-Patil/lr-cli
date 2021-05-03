package hooks

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/loginradius/lr-cli/api"
	"github.com/spf13/cobra"
)

func NewHooksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hooks",
		Short: "Gets hooks",
		Long: heredoc.Doc(`
		This command fetches the list of webhooks.
		`),
		Example: heredoc.Doc(`
			$ lr get hooks
 
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getHooks()
		},
	}
	return cmd
}

func getHooks() error {
	Hooks, err := api.GetHooks()
	if err != nil {
		return err
	}
	numberOfHooks := len(Hooks.Data)
	for i := 0; i < numberOfHooks; i++ {
		fmt.Println(i + 1)
		fmt.Println("	Name: ", Hooks.Data[i].Name)
		fmt.Println(" Event: ", Hooks.Data[i].Event)
		fmt.Println(" TargetUrl: ", Hooks.Data[i].Targeturl)
		if i != numberOfHooks-1 {
			fmt.Println("-------------------------------------------------")
		}
	}
	return nil

}
