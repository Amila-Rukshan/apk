/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
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

package config

// Configuration object which is populated with default values.
var defaultConfig = &Config{
	ManagementServer: managementServer{
		XDSPort:    19000,
		NodeLabels: []string{"default"},
		GRPCPort: 8765,
	},
	Database: database{
		Name:     "apk-db",
		Username: "postgres",
		Password: "amila123",
		Host:     "localhost",
		Port:     5432,
		PoolOptions: dbPool{
			PoolMaxConns:              4,
			PoolMinConns:              0,
			PoolMaxConnLifetime:       "1h",
			PoolMaxConnIdleTime:       "1h",
			PoolHealthCheckPeriod:     "1m",
			PoolMaxConnLifetimeJitter: "1s",
		},
	},
}
