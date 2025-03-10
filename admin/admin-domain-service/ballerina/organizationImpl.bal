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

import ballerina/uuid;
import wso2/apk_common_lib as commons;

isolated function createInternalFromOrganization(Organization payload) returns Internal_Organization {
    OrganizationClaim orgClaim = {
        claimKey: "organizationClaimKey",
        claimValue: payload.organizationClaimValue
    };

    Internal_Organization internalOrganization = {
        id: payload.id.toString(),
        name: payload.name,
        displayName: payload.displayName,
        enabled: payload.enabled,
        serviceNamespaces: payload.serviceNamespaces,
        production: payload.production,
        sandbox: payload.sandbox,
        claimList: [ orgClaim ]
    };
    return internalOrganization;
}

isolated function createOrganizationFromInternal(Internal_Organization payload) returns Organization {
    Organization organization = {
        id: payload.id,
        name: payload.name,
        displayName: payload.displayName,
        enabled: payload.enabled,
        serviceNamespaces: payload.serviceNamespaces,
        production: payload.production,
        sandbox: payload.sandbox,
        organizationClaimValue: payload.claimList[0].claimValue
    };
    return organization;
}

isolated function addOrganization(Organization payload) returns Organization|commons:APKError {
    boolean validateOrganization = check validateOrganizationByNameDAO(payload.name);
    if validateOrganization is true {
        return e909427(payload.name);
    }
    payload.id = uuid:createType1AsString();
    Internal_Organization|commons:APKError organization = addOrganizationDAO(createInternalFromOrganization(payload));
    if organization is Internal_Organization {
        Organization createdOrganization = createOrganizationFromInternal(organization);
        return createdOrganization;
    } else {
        return organization;
    } 
}

isolated function updatedOrganization(string id, Organization payload) returns Organization|commons:APKError {
    boolean validateOrganizationId = check validateOrganizationById(id);
    if validateOrganizationId is false {
        return e909428(id);
    }
    payload.id = id;
    Internal_Organization|commons:APKError organization = updateOrganizationDAO(id, createInternalFromOrganization(payload));
    if organization is Internal_Organization {
       return createOrganizationFromInternal(organization);
    } else {
        return organization;
    } 
}

isolated function getAllOrganization() returns OrganizationList|commons:APKError {
    Internal_Organization[]|commons:APKError getOrgnizations = getAllOrganizationDAO();
    if getOrgnizations is Internal_Organization[] {
        int count = getOrgnizations.length();
        Organization[] organizations = [];
        foreach var organization in getOrgnizations {
            organizations.push(createOrganizationFromInternal(organization));
        }
        OrganizationList getOrgnizationsList = {count: count, list: organizations};
        return getOrgnizationsList;
    } else {
       return getOrgnizations;
    }
}

isolated function getOrganizationById(string id) returns Organization|commons:APKError {
    Internal_Organization|commons:APKError organization = getOrganizationByIdDAO(id);
    if organization is Internal_Organization {
        return createOrganizationFromInternal(organization);
    } else {
        return organization;
    }
}

isolated function removeOrganization(string id) returns boolean|commons:APKError {
    boolean validateOrganizationId = check validateOrganizationById(id);
    if validateOrganizationId is false {
        return e909428(id);
    }
    boolean|commons:APKError organization = removeOrganizationDAO(id);
    return organization;
}

isolated function getOrganizationByOrganizationClaim() returns Organization|commons:APKError{
    //TO DO: Get organization claim from JWT
    string organizationClaimValue = "organizationClaimValue";
    Internal_Organization|commons:APKError organization = getOrganizationByOrganizationClaimDAO(organizationClaimValue);
    if organization is Internal_Organization {
        return createOrganizationFromInternal(organization);
    } else {
        return organization;
    }
}
