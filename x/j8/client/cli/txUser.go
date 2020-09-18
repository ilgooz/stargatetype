package cli

import (
  
	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/j/j8/x/j8/types"
)

func CmdCreateUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-user [name] [email]",
		Short: "Creates a new user",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsName := string(args[0])
      argsEmail := string(args[1])
      
        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgUser(clientCtx.GetFromAddress(), argsName, argsEmail)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}
