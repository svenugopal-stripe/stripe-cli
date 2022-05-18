package resource

import (
	"net/http"

	"github.com/spf13/cobra"

	"github.com/stripe/stripe-cli/pkg/config"
)

// OrdersCreateCmd represents the order create API operation command. This
// command is manually defined because of quirkiness with the OpenAPI spec autogenerated command and required parameters.
type OrdersCreateCmd struct {
	opCmd *OperationCmd
}

func (occ *OrdersCreateCmd) runOrdersCreateCmd(cmd *cobra.Command, args []string) error {
	return occ.opCmd.runOperationCmd(cmd, args)
}

// NewOrdersCreateCmd creates a new orders creation sub command.
func NewOrdersCreateCmd(parentCmd *cobra.Command, cfg *config.Config) *OrdersCreateCmd {
	ordersCreateCmd := &OrdersCreateCmd{
		opCmd: NewOperationCmd(parentCmd, "create", "/v1/orders", http.MethodPost, map[string]string{
			"currency":               "string",
			"line_items[][product]":  "string",
			"line_items[][quantity]": "integer",
			"automatic_tax[enabled]": "boolean",
		}, cfg),
	}

	ordersCreateCmd.opCmd.Cmd.RunE = ordersCreateCmd.runOrdersCreateCmd

	return ordersCreateCmd
}
