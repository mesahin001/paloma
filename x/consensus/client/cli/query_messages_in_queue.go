package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/palomachain/paloma/x/consensus/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMessagesInQueue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "messages-in-queue [queue-type-name]",
		Short: "Query MessagesInQueue",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqQueueTypeName := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryMessagesInQueueRequest{

				QueueTypeName: reqQueueTypeName,
			}

			res, err := queryClient.MessagesInQueue(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
