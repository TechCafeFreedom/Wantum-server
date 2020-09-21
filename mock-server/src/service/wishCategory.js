const protoUtil = require("../utility/protoUtil");

module.exports = class WishCategoryService {
  constructor() {
    const protoDescriptor = protoUtil.getProtoDescriptor("wish_category.proto");
    this.wishCategoryProto = protoDescriptor.proto_wishcategory;
  }

  createWishCategory(call, callback) {
    callback(null, {});
  }

  // QUESTION: Name -> Titleのがよくね
  updateWishCategoryName(call, callback) {
    callback(null, {});
  }

  // TODO: deleteってないの...
};