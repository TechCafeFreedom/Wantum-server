/**
 * return dummy user
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
exports.getDummyUser = function(request) {
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