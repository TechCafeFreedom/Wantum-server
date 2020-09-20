const protoUtil = require("../utility/protoUtil");
const userResource = require("../resource/user");

const protoDescriptor = protoUtil.getProtoDescriptor("user.proto");
exports.userProto = protoDescriptor.proto_user;

/**
 * mock service in userService
 * 
 * @param {*} call 
 * @param {*} callback 
 */
exports.createUser = function (call, callback) {
  callback(null, userResource.getDummyUser(call.request));
}

/**
 * mock service in userService
 * 
 * @param {*} call
 * @param {*} callback
 */
exports.getMyProfile = function (call, callback) {
  callback(null, userResource.getDummyUser(call.request));
}

/**
 * mock service in userService
 *
 * @param {*} call
 * @param {*} callback
 */
exports.getUserProfile = function (call, callback) {
  callback(null, userResource.getDummyUser(call.request));
}

/**
 * mock service in userService
 *
 * @param {*} call
 * @param {*} callback
 */
exports.updateUserProfile = function (call, callback) {
  callback(null, userResource.getDummyUser(call.request));
}
