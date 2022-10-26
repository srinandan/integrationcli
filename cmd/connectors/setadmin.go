// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package connectors

import (
	"github.com/spf13/cobra"
	"github.com/srinandan/integrationcli/apiclient"
	"github.com/srinandan/integrationcli/client/connectors"
)

// SetAdminCmd to set admin role
var SetAdminCmd = &cobra.Command{
	Use:   "setadmin",
	Short: "Set Connection Admin IAM policy on a Connection",
	Long:  "Set Connection Admin IAM policy on a Connection",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		if err = apiclient.SetRegion(region); err != nil {
			return err
		}
		return apiclient.SetProjectID(project)
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return connectors.SetIAM(name, memberName, "admin", memberType)
	},
}

func init() {

	SetAdminCmd.Flags().StringVarP(&memberName, "member", "m",
		"", "Member Name, example Service Account Name")
	SetAdminCmd.Flags().StringVarP(&memberType, "memberType", "p",
		"serviceAccount", "memberType must be serviceAccount, user or group")

	_ = SetAdminCmd.MarkFlagRequired("name")
}