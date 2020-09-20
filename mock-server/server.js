// TODO: 以下、userに関わる部分はどこかに切り出し
var PROTO_PARH = __dirname + "/../Wantum-ProtocolBuffer/user.proto";
var grpc = require("grpc");
var protoLoader = require("@grpc/proto-loader");

var packageDefinition = protoLoader.loadSync(
  PROTO_PARH,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
  }
);
var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
var userProto = protoDescriptor.proto_user;

/**
 * mock service in userService
 * 
 * @param {*} call 
 * @param {*} callback 
 */
function createUser(call, callback) {
  callback(null, getDummyUser(call.request));
}

function getMyProfile(call, callback) {
  callback(null, getDummyUser(call.request));
}

function getUserProfile(call, callback) {
  callback(null, getDummyUser(call.request));
}

function updateUserProfile(call, callback) {
  callback(null, getDummyUser(call.request));
}

// TODO: 以下、dummyデータの生成部分はどこかに切り出し

/**
 * return dummy data
 * 
 * @param {Object} request 
 * @param {Number} request.user_id
 * @param {String} request.name
 * @param {String} request.user_name
 * @param {Object} request.thumbnail
 * @param {String} request.bio
 * @param {Enumerator} request.gender
 * @param {String} request.place
 * @param {Number} request.birth
 */
function getDummyUser(request) {
  return {
    user_id: request.user_id ? request.user_id : 1,
    name: request.name ? request.name : "hogehoge",
    user_name: request.user_name ? request.user_name : "fugafuga",
    thumbnail: "hogehoge.png",
    bio: request.bio ? request.bio : "hogefuga",
    gender: request.gender ? request.gender : "MAN",
    place: request.place ? request.place : "tushima",
    birth: request.birth ? request.birth : 1600513200,
  };
}

// NOTE: ここから下がこのファイル分

/**
 * get server instance
 * 
 * @return {grpc.Server}
 */
function getServer() {
  var server = new grpc.Server();
  // userService
  server.addProtoService(userProto.UserService.service, {
    createUser: createUser,
    getMyProfile: getMyProfile,
    getUserProfile: getUserProfile,
    updateUserProfile: updateUserProfile,
  });
  return server;
}

// start server
var mockServer = getServer();
mockServer.bind("0.0.0.0:50051", grpc.ServerCredentials.createInsecure());
mockServer.start();
console.log("mock server linten 0.0.0.0:50051...");
