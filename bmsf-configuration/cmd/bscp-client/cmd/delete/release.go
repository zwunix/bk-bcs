/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package delete

import (
	"bk-bscp/cmd/bscp-client/option"
	"bk-bscp/cmd/bscp-client/service"
	"context"

	"github.com/spf13/cobra"
)

//deleteStrategyCmd: client update commit
func deleteStrategyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "strategy",
		Short: "delete release strategy",
		Long:  "delete release strategy content",
		Example: `
	bscp-client delete strategy --Id xxxx
		`,
		RunE: handleDeleteStrategy,
	}
	//command line flags
	cmd.Flags().String("Id", "", "strategy ID")
	cmd.MarkFlagRequired("Id")
	return cmd
}

func handleDeleteStrategy(cmd *cobra.Command, args []string) error {
	//get global command info and create business operator
	operator := service.NewOperator(option.GlobalOptions)
	if err := operator.Init(option.GlobalOptions.ConfigFile); err != nil {
		return err
	}
	ID, _ := cmd.Flags().GetString("Id")
	//update Commit and check result
	if err := operator.DeleteStrategy(context.TODO(), ID); err != nil {
		return err
	}
	cmd.Printf("Delete Strategy successfully: %s\n", ID)
	return nil
}
