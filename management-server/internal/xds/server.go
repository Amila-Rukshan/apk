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

package xds

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	apkmgt_application "github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/apkmgt"
	apkmgt_service "github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/service/apkmgt"
	wso2_cache "github.com/wso2/apk/adapter/pkg/discovery/protocol/cache/v3"
	wso2_resource "github.com/wso2/apk/adapter/pkg/discovery/protocol/resource/v3"
	wso2_server "github.com/wso2/apk/adapter/pkg/discovery/protocol/server/v3"
	"github.com/wso2/apk/adapter/pkg/logging"
	"github.com/wso2/apk/management-server/internal/config"
	"github.com/wso2/apk/management-server/internal/logger"
	"github.com/wso2/apk/management-server/internal/xds/callbacks"
	"google.golang.org/grpc"
)

var (
	apiCache      wso2_cache.SnapshotCache
	apiCacheMutex sync.Mutex
	Sent          bool = true
)

const (
	maxRandomInt             int    = 999999999
	typeURL                  string = "wso2.discovery.apkmgt.Application"
	grpcMaxConcurrentStreams        = 1000000
)

// IDHash uses ID field as the node hash.
type IDHash struct{}

// ID uses the node ID field
func (IDHash) ID(node *corev3.Node) string {
	if node == nil {
		return "unknown"
	}
	return node.Id
}

var _ wso2_cache.NodeHash = IDHash{}

func init() {
	apiCache = wso2_cache.NewSnapshotCache(false, IDHash{}, nil)
	rand.Seed(time.Now().UnixNano())
}

// FeedData mock data
func FeedData() {
	config := config.ReadConfigs()
	logger.LoggerXdsServer.Debug("adding mock data")
	// version := rand.Intn(maxRandomInt)
	application1 := apkmgt_application.Application{
		Uuid:  "app-uuid-1",
		Name:  "app1",
		Owner: "admin",
		Attributes: map[string]string{
			"attb1": "value1",
			"attb2": "value2",
		},
		ConsumerKeys: []*apkmgt_application.ConsumerKey{
			{Key: "yef14gh8syDvTt56rdtIHYbjF_Ya", KeyManager: "Resident Key Manager"},
		},
		Subscriptions: []*apkmgt_application.Subscription{
			{Uuid: "a0d73fe1-fdfd-4dd2-a10a-6b0057a7491d",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
			{Uuid: "e2ae04e7-b886-4985-8d9f-3da5dbf25e69",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
			{Uuid: "0b49c011-44d0-4698-8184-7caab4b06b4f",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
			{Uuid: "1325b443-4b35-4999-818f-6762c08740fa",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
			{Uuid: "cfb14e32-c3f9-44fe-ba65-caaa90a9bd1e",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
			{Uuid: "c334806b-2493-417b-ad9c-7eae3dab560f",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
			{Uuid: "fb7c4bdb-ccf7-460a-b2ee-7910179b32c4",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
			{Uuid: "2e9965fe-964c-4334-a52d-d00ccca537d9",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
			{Uuid: "1437fdb4-515b-4339-985d-d5b46355de94",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
			{Uuid: "805367fb-1ed7-4791-a823-69bd40e21271",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
			{Uuid: "d7c9e07e-90f8-4269-a0ba-074842cc44ab",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
			{Uuid: "75cb5c25-7268-4398-99d6-e82735517ed1",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
			{Uuid: "be1ff755-b114-4591-93e4-ad7b844a4f6a",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
			{Uuid: "eea2b16c-41dc-4849-8e28-b6d4cea5c09e",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
		},
	}
	application2 := apkmgt_application.Application{
		Uuid:  "app-uuid-2",
		Name:  "app2",
		Owner: "Alice",
		Attributes: map[string]string{
			"attb3": "value3",
			"attb4": "value4",
		},
		ConsumerKeys: []*apkmgt_application.ConsumerKey{
			{Key: "hef14gh8syDvTtvoWYeIHYbjF_Ya", KeyManager: "Resident Key Manager"},
		},
		Subscriptions: []*apkmgt_application.Subscription{
			{Uuid: "ff041d1b-be19-4529-a861-86a79905a1ad",
				PolicyId:           "Unlimited",
				SubscriptionStatus: "ACTIVE",
			},
		},
	}
	newSnapshot, _ := wso2_cache.NewSnapshot(fmt.Sprint(1), map[wso2_resource.Type][]types.Resource{
		wso2_resource.APKMgtApplicationType: {&application1, &application2},
	})
	apiCacheMutex.Lock()
	// TODO (amaliMatharaarachchi) update relevant adapters with snapshot caches
	apiCache.SetSnapshot(context.Background(), config.ManagementServer.NodeLabels[0], newSnapshot)
	apiCacheMutex.Unlock()
	time.Sleep(10 * time.Second)
}

func InitAPKMgtServer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	apkMgtAPIDsSrv := wso2_server.NewServer(ctx, apiCache, &callbacks.Callbacks{})

	var grpcOptions []grpc.ServerOption
	grpcOptions = append(grpcOptions, grpc.MaxConcurrentStreams(grpcMaxConcurrentStreams))
	grpcServer := grpc.NewServer(grpcOptions...)
	apkmgt_service.RegisterAPKMgtDiscoveryServiceServer(grpcServer, apkMgtAPIDsSrv)
	config := config.ReadConfigs()
	port := config.ManagementServer.XDSPort

	//todo (amaliMatharaarachchi) handle error gracefully
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.LoggerServer.ErrorC(logging.ErrorDetails{
			Message:   fmt.Sprintf("Error while listening on port: %v. Error: %v", port, err.Error()),
			Severity:  logging.BLOCKER,
			ErrorCode: 1000,
		})
	}

	logger.LoggerServer.Infof("APK Management server XDS is starting on port %v.", port)
	if err = grpcServer.Serve(listener); err != nil {
		logger.LoggerServer.ErrorC(logging.ErrorDetails{
			Message:   fmt.Sprint("Error while starting APK Management server XDS server."),
			Severity:  logging.BLOCKER,
			ErrorCode: 1001,
		})
	}
}
