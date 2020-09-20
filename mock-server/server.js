// var userService = require("./src/service/user");

// var grpc = require("grpc");

// // NOTE: ここから下がこのファイル分

// /**
//  * get server instance
//  * 
//  * @return {grpc.Server}
//  */
// function getServer() {
//   var server = new grpc.Server();
//   // userService
//   server.addProtoService(userService.userProto.UserService.service, {
//     createUser: userService.createUser,
//     getMyProfile: userService.getMyProfile,
//     getUserProfile: userService.getUserProfile,
//     updateUserProfile: userService.updateUserProfile,
//   });
//   return server;
// }

// // start server
// // var mockServer = getServer();
// // mockServer.bind("0.0.0.0:50051", grpc.ServerCredentials.createInsecure());
// // mockServer.start();
// // console.log("mock server linten 0.0.0.0:50051...");
