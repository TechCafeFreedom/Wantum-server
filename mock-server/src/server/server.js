const grpc = require("grpc");
const userService = require("../service/user");

/**
 * get server instance
 * 
 * @return {grpc.Server}
 */
exports.getServer = function () {
  var server = new grpc.Server();
  // userService
  server.addProtoService(userService.userProto.UserService.service, {
    createUser: userService.createUser,
    getMyProfile: userService.getMyProfile,
    getUserProfile: userService.getUserProfile,
    updateUserProfile: userService.updateUserProfile,
  });
  return server;
}