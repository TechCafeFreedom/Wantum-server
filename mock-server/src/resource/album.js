const memoryResource = require("./memory");

/**
 * return dummy album
 * 
 * @param {Object} request
 * @param {Number} request.album_id
 * @param {String} request.title
 * @param {Array} request.memories
 */
exports.getDummyAlbum = function (request) {
  return {
    album_id: request.album_id ? request.album_id : 1,
    title: request.title ? request.title : "hogehoge",
    memories: request.memories ? request.memories : memoryResource.getDummyMemories() // TODO: set data
  };
};
