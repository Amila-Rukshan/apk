import ballerina/http;

service /commonauth on ep0 {
    
    resource function post login(http:Request request) returns http:Found|UnauthorizedLoginErrorResponse {
        LoginClient loginClient  =new;
        return loginClient.handleLoginRequest(request);
    }
}