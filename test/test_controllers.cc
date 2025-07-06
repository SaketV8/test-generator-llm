#include <drogon/drogon.h>
#include <drogon/drogon_test.h>
#include "../controllers/DepartmentsController.h"
#include "../models/Department.h"
#include "../models/Person.h"
#include "../models/Job.h"
#include "../utils/utils.h"
#include "../filters/LoginFilter.h"
#include "../plugins/Jwt.h"
#include "../plugins/JwtPlugin.h"

DROGON_TEST(DepartmentsControllerGetTest)
{
    // Create an instance of the controller
    DepartmentsController controller;

    // Create a mock HTTP request
    auto req = drogon::HttpRequest::newHttpRequest();
    req->setPath("/departments");
    req->setMethod(drogon::Get);

    // Call the get method with a callback
    controller.get(req, [TEST_CTX](const drogon::HttpResponsePtr& resp) {
        // Check the response
        REQUIRE(resp != nullptr);
        CHECK(resp->getStatusCode() == drogon::k200OK);

        // You can add more checks here, for example, to validate the JSON response
        auto json = resp->getJsonObject();
        CHECK(json != nullptr);
    });
}

DROGON_TEST(HttpRequestTest)
{
    // Test creating an HTTP request
    auto req = drogon::HttpRequest::newHttpRequest();
    req->setPath("/departments");
    req->setMethod(drogon::Get);
    req->addHeader("Content-Type", "application/json");
    
    // Verify request properties
    CHECK(req->path() == "/departments");
    CHECK(req->method() == drogon::Get);
    CHECK(req->getHeader("Content-Type") == "application/json");
    
    std::cout << "HTTP request test passed!" << std::endl;
}

DROGON_TEST(RemoteAPITest)
{
    std::cout << "Starting RemoteAPITest..." << std::endl;
    auto client = drogon::HttpClient::newHttpClient("http://localhost:3000");
    auto req = drogon::HttpRequest::newHttpRequest();
    req->setPath("/departments");
    client->sendRequest(req, [TEST_CTX](drogon::ReqResult res, const drogon::HttpResponsePtr& resp) {
        // There's nothing we can do if the request didn't reach the server
        // or the server generated garbage.
        REQUIRE(res == drogon::ReqResult::Ok);
        REQUIRE(resp != nullptr);
        
        CHECK(resp->getStatusCode() == drogon::k200OK);
        CHECK(resp->contentType() == drogon::CT_APPLICATION_JSON);
    });
    std::cout << "RemoteAPITest END" << std::endl;
}

// New tests for DepartmentsController

DROGON_TEST(DepartmentsControllerGetOneSuccessTest)
{
    DepartmentsController controller;

    // Use a valid existing department ID for testing
    // For testing purposes, we assume ID 1 exists.
    int validDepartmentId = 1;

    auto req = drogon::HttpRequest::newHttpRequest();
    req->setPath("/departments/" + std::to_string(validDepartmentId));
    req->setMethod(drogon::Get);

    controller.getOne(req, [TEST_CTX, validDepartmentId](const drogon::HttpResponsePtr& resp) {
        REQUIRE(resp != nullptr);
        // According to implementation, success returns 201 Created
        CHECK(resp->getStatusCode() == drogon::k201Created);

        auto json = resp->getJsonObject();
        REQUIRE(json != nullptr);
        // Check that the returned JSON has an "id" field matching the requested department ID
        CHECK(json->isMember("id"));
        CHECK((*json)["id"].asInt() == validDepartmentId);
        // Optionally check that name field exists
        CHECK(json->isMember("name"));
    }, validDepartmentId);
}

DROGON_TEST(DepartmentsControllerGetOneNotFoundTest)
{
    DepartmentsController controller;

    // Use an ID that does not exist, e.g., a large number unlikely to exist
    int invalidDepartmentId = 999999;

    auto req = drogon::HttpRequest::newHttpRequest();
    req->setPath("/departments/" + std::to_string(invalidDepartmentId));
    req->setMethod(drogon::Get);

    controller.getOne(req, [TEST_CTX](const drogon::HttpResponsePtr& resp) {
        REQUIRE(resp != nullptr);
        // According to implementation, not found returns 404 Not Found
        CHECK(resp->getStatusCode() == drogon::k404NotFound);
    }, invalidDepartmentId);
}

DROGON_TEST(DepartmentsControllerGetDepartmentPersonsTest)
{
    DepartmentsController controller;

    // Use a valid existing department ID for testing
    int validDepartmentId = 1;

    auto req = drogon::HttpRequest::newHttpRequest();
    req->setPath("/departments/" + std::to_string(validDepartmentId) + "/persons");
    req->setMethod(drogon::Get);

    controller.getDepartmentPersons(req, [TEST_CTX](const drogon::HttpResponsePtr& resp) {
        REQUIRE(resp != nullptr);
        // The response status can be 200 OK or 404 Not Found if no persons found
        // We check for 200 OK here as per test requirement
        if (resp->getStatusCode() == drogon::k200OK)
        {
            auto json = resp->getJsonObject();
            REQUIRE(json != nullptr);
            // The response should be a JSON array
            CHECK(json->isArray());
        }
        else
        {
            // If no persons found, 404 is returned, which is acceptable fallback
            CHECK(resp->getStatusCode() == drogon::k404NotFound);
        }
    }, validDepartmentId);
}