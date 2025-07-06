// #include <drogon/drogon.h>
// #include <drogon/drogon_test.h>
// #include <drogon/drogon.h>
// #include <drogon/drogon_test.h>
// #include <drogon/HttpAppFramework.h>

// DROGON_TEST(BasicTest)
// {
//     // A simple test that should always pass
//     CHECK(1 + 1 == 2);
    
//     // Test that we can create a JSON object
//     Json::Value json;
//     json["test"] = "value";
//     CHECK(json["test"].asString() == "value");
// }





// DROGON_TEST(RemoteAPITest2)
// {
//     std::cout << "Starting RemoteAPITest2..." << std::endl;
    
//     auto client = drogon::HttpClient::newHttpClient("http://localhost:3000");
//     auto req = drogon::HttpRequest::newHttpRequest();
//     req->setPath("/departments");
    
//     std::cout << "Sending request to /departments..." << std::endl;
    
//     client->sendRequest(req, [](drogon::ReqResult res, const drogon::HttpResponsePtr& resp) {
//         std::cout << "Received response with status: " << static_cast<int>(res) << std::endl;
        
//         if (res != drogon::ReqResult::Ok) {
//             std::cout << "Request failed!" << std::endl;
//             return;
//         } else {
//             std::cout << "Request failed!" << std::endl;
//             return;
//         }
        
//         if (!resp) {
//             std::cout << "Response is null!" << std::endl;
//             return;
//         }
        
//         std::cout << "Response status code: " << resp->getStatusCode() << std::endl;
//         std::cout << "Response content type: " << resp->contentType() << std::endl;
//         std::cout << "Response body: " << resp->getBody() << std::endl;
//     });
    
//     std::cout << "Request sent, waiting for response..." << std::endl;
// }




// #include <drogon/drogon.h>
// #include <drogon/drogon_test.h>
// #include "../controllers/DepartmentsController.h"
// #include "../models/Department.h"

// DROGON_TEST(DepartmentsControllerGetTest)
// {
//     // Create an instance of the controller
//     DepartmentsController controller;

//     // Create a mock HTTP request
//     auto req = drogon::HttpRequest::newHttpRequest();
//     req->setPath("/departments");
//     req->setMethod(drogon::Get);

//     // Call the get method with a callback
//     controller.get(req, [TEST_CTX](const drogon::HttpResponsePtr& resp) {
//         // Check the response
//         REQUIRE(resp != nullptr);
//         CHECK(resp->getStatusCode() == drogon::k200OK);

//         // You can add more checks here, for example, to validate the JSON response
//         auto json = resp->getJsonObject();
//         CHECK(json != nullptr);
//     });
// }


// DROGON_TEST(HttpRequestTest)
// {
//     // Test creating an HTTP request
//     auto req = drogon::HttpRequest::newHttpRequest();
//     req->setPath("/departments");
//     req->setMethod(drogon::Get);
//     req->addHeader("Content-Type", "application/json");
    
//     // Verify request properties
//     CHECK(req->path() == "/departments");
//     CHECK(req->method() == drogon::Get);
//     CHECK(req->getHeader("Content-Type") == "application/json");
    
//     std::cout << "HTTP request test passed!" << std::endl;
// }

// // #include <drogon/test/HttpTestClient.h>


// DROGON_TEST(RemoteAPITest)
// {
//     std::cout << "Starting RemoteAPITest..." << std::endl;
//     auto client = drogon::HttpClient::newHttpClient("http://localhost:3000");
//     auto req = drogon::HttpRequest::newHttpRequest();
//     req->setPath("/departments");
//     client->sendRequest(req, [TEST_CTX](drogon::ReqResult res, const drogon::HttpResponsePtr& resp) {
//         // There's nothing we can do if the request didn't reach the server
//         // or the server generated garbage.
//         REQUIRE(res == drogon::ReqResult::Ok);
//         REQUIRE(resp != nullptr);
        
//         CHECK(resp->getStatusCode() == drogon::k200OK);
//         CHECK(resp->contentType() == drogon::CT_APPLICATION_JSON);
//     });
//     std::cout << "RemoteAPITest END" << std::endl;
// }

// #include <drogon/drogon.h>
// #include <drogon/drogon_test.h>
// #include "../controllers/DepartmentsController.h"
// #include "../models/Department.h"
// #include "../models/Person.h"
// #include "../models/Job.h"
// #include "../utils/utils.h"
// #include "../filters/LoginFilter.h"

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