const grpc = require("grpc");
const grpcServer = require("./server/server");

// start server
var mockServer = grpcServer.getServer();
mockServer.bind("0.0.0.0:50051", grpc.ServerCredentials.createInsecure());
mockServer.start();
console.log("mock server lintening 0.0.0.0:50051...");