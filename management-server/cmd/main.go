/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org).
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/service/apkmgt"
	"github.com/wso2/apk/management-server/internal/database"

	"github.com/wso2/apk/management-server/internal/logger"
)

func main() {
	logger.LoggerServer.Info("Starting Management server ...")
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	// connect to the postgres database
	database.ConnectToDB()
	defer database.CloseDBConn()
	// go xds.InitAPKMgtServer()
	// go xds.FeedData()
	// todo(amaliMatharaarachchi) watch data updates and update snapshot accordingly.
	// temp data
	// var arr = []*internal_types.ApplicationEvent{
	// 	{
	// 		Label:         "dev",
	// 		UUID:          "b9850225-c7db-444d-87fd-4feeb3c6b3cc",
	// 		IsRemoveEvent: false,
	// 	},
	// 	{
	// 		Label:         "stage",
	// 		UUID:          "6e2dc623-1a23-46a3-86cf-389d63bbbc3e",
	// 		IsRemoveEvent: false,
	// 	},
	// }

	api := &apkmgt.API{
		Uuid:           "api-uuid-1",
		Name:           "api-1",
		Context:        "/web-hook/v1.0.0",
		Version:        "v1.0.0",
		Provider:       "webhooke.site",
		OrganizationId: "org-abc",
		CreatedBy:      "Bob",
		CreatedTime:    "2022-11-11T08:36:54.058Z",
		Definition:     "{\"swagger\": \"json\"}",
	}

	apiUpdated := &apkmgt.API{
		Uuid:           "api-uuid-1",
		Name:           "api-1-new-name",
		Context:        "/web-hook/v2.0.0",
		Version:        "v2.0.0",
		Provider:       "webhooke.site",
		OrganizationId: "org-abc-updated",
		UpdatedBy:      "Bob",
		UpdatedTime:    "2022-12-11T08:36:54.058Z",
		Definition:     "{\"swagger\": \"json-updated\"}",
	}

	database.CreateAPI(api)
	time.Sleep(5 * time.Second)
	database.UpdateAPI(apiUpdated)
	time.Sleep(15 * time.Second)
	// database.DeleteAPI(api)
	// go xds.AddMultipleApplications(arr)
	// go server.StartGRPCServer()

OUTER:
	for {
		select {
		case s := <-sig:
			switch s {
			case os.Interrupt:
				break OUTER
			}
		}
	}
}
