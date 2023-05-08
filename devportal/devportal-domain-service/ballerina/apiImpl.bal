//
// Copyright (c) 2022, WSO2 LLC. (http://www.wso2.com).
//
// WSO2 LLC. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

import ballerina/jballerina.java;
import devportal_service.org.wso2.apk.devportal.sdk as sdk;
import devportal_service.java.util as javautil;
import devportal_service.java.lang as javalang;
import ballerina/http;
import wso2/apk_common_lib as commons;

isolated function getAPIByAPIId(string apiId) returns API|NotFoundError|APKError {
    API|APKError|NotFoundError api = getAPIByIdDAO(apiId);
    return api;
}

isolated function getAPIList(int 'limit, int  offset, string? query, commons:Organization organization) returns APIList|APKError {
    if query !is string {
        API[]|APKError apis = getAPIsDAO(organization.uuid);
        if apis is API[] {
            API[] limitSet = [];
            if apis.length() > offset {
                foreach int i in offset ... (apis.length() - 1) {
                    if limitSet.length() < 'limit {
                        limitSet.push(apis[i]);
                    }
                }
            }
            APIList apisList = {count: limitSet.length(), list: limitSet,pagination: {total: apis.length(), 'limit: 'limit, offset: offset}};
            return apisList;
        } else {
            return apis;
        }
    } else {
        boolean hasPrefix = query.startsWith("content");
        if hasPrefix {
            int? index = query.indexOf(":");
            if index is int {
                string modifiedQuery = "%" + query.substring(index+1) +"%";
                API[]|APKError apis = getAPIsByQueryDAO(modifiedQuery, organization.uuid);
                if apis is API[] {
                    API[] limitSet = [];
                    if apis.length() > offset {
                        foreach int i in offset ... (apis.length() - 1) {
                            if limitSet.length() < 'limit {
                                limitSet.push(apis[i]);
                            }
                        }
                    }
                    APIList apisList = {count: limitSet.length(), list: limitSet,pagination: {total: apis.length(), 'limit: 'limit, offset: offset}};
                    return apisList;
                } else {
                    return apis;
                }
            } else {
                string message = "Invalid Content Search Text Provided. Missing :";
                APKError e = error(message, message = message, description = message, code = 90911, statusCode = "400");
                return e;
            }
        } else {
            string message = "Invalid Content Search Text Provided. Missing content keyword";
            APKError e = error(message, message = message, description = message, code = 90911, statusCode = "400");
            return e;
        }
    }
}

isolated function getAPIDefinition(string apiId) returns APIDefinition|NotFoundError|APKError {
    APIDefinition|NotFoundError|APKError apiDefinition = getAPIDefinitionDAO(apiId);
    return apiDefinition;
}

function generateSDKImpl(string apiId, string language) returns http:Response|sdk:APIClientGenerationException|NotFoundError|APKError {
    sdk:APIClientGenerationManager sdkClient = new sdk:APIClientGenerationManager(newSDKClient());
    string apiName;
    string apiVersion;
    API|NotFoundError|APKError api = getAPIByAPIId(apiId);
    if api is API {
        apiName = api.name;
        apiVersion = api.'version;
        APIDefinition|NotFoundError|APKError apiDefinition = getAPIDefinition(apiId);
        if apiDefinition is APIDefinition {
            string? schema = apiDefinition.schemaDefinition;
            if schema is string {
                javautil:Map|sdk:APIClientGenerationException sdkMap = sdkClient.generateSDK(language,apiName,apiVersion,schema,
                sdkConfig.groupId, sdkConfig.artifactId, sdkConfig.modelPackage, sdkConfig.apiPackage);
                if sdkMap is javautil:Map {
                    string path = readMap(sdkMap,"zipFilePath");
                    string fileName = readMap(sdkMap,"zipFileName");
                    http:Response response = new;
                    response.setFileAsPayload(path);
                    response.addHeader("Content-Disposition","attachment; filename=" + fileName);
                    return response;
                } else {
                    return sdkMap;
                }
            } else {
                string message = "Unable to retrieve schema mediation";
                APKError e = error(message, message = message, description = message, code = 90911, statusCode = "500");
                return e;
            }
        }
    } else if api is NotFoundError|APKError {
        return api;
    }
    string message = "Unable to generate SDK";
    APKError e = error(message, message = message, description = message, code = 90911, statusCode = "500");
    return e;
}

isolated function getSDKLanguages() returns string|json|APKError {
    sdk:APIClientGenerationManager sdkClient = new sdk:APIClientGenerationManager(newSDKClient());
    string? sdkLanguages = sdkClient.getSupportedSDKLanguages();
    return sdkLanguages;
}

isolated function newSDKClient() returns handle = @java:Constructor {
    'class: "org.wso2.apk.devportal.sdk.APIClientGenerationManager"
} external;

function readMap(javautil:Map sdkMap, string key) returns string{
    handle keyAsJavaStr = java:fromString(key);
    javalang:Object keyAsObj = new (keyAsJavaStr);
    javalang:Object value = sdkMap.get(keyAsObj);
    // Above simplified to one line
    // javalang:Object value = sdkMap.get(new (java:fromString("zipFilePath")));
    handle valueHandle = value.jObj;
    string? valueStr = java:toString(valueHandle);
    if valueStr is string {
        return valueStr;
    } else {
        return "";
    }
}
