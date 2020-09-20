const protoUtil = require("../utility/protoUtil");
const UserResourceBuilder = require("../resource/user");

/**
 * mock service for userService
 */
class UserService {
  constructor() {
    const protoDescriptor = protoUtil.getProtoDescriptor("user.proto");
    this.userProto = protoDescriptor.proto_user;
  }

  createUser(call, callback) {
    callback(
      null,
      new UserResourceBuilder()
        .withName(call.request.name)
        .withUserName(call.request.user_name)
        .withBio(call.request.bio)
        .withGender(call.request.gender)
        .withPhone(call.request.phone)
        .withPlace(call.request.place)
        .withBirth(call.request.birth)
        .build()
    );
  }

  getMyProfile(call, callback) {
    callback(
      null,
      new UserResourceBuilder()
        .build()
    );
  }

  getUserProfile(call, callback) {
    callback(
      null,
      new UserResourceBuilder()
        .withUserName(call.request.user_name)
        .build()
    );
  }

  updateUserProfile(call, callback) {
    callback(
      null,
      new UserResourceBuilder()
        .withUserId(call.request.user_id)
        .withName(call.request.name)
        .withUserName(call.request.user_name)
        .withBio(call.request.bio)
        .withGender(call.request.gender)
        .withPhone(call.request.phone)
        .withPlace(call.request.place)
        .withBirth(call.request.birth)
        .build()
    );
  }
}

module.exports = UserService;