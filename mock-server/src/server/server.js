const grpc = require("grpc");
const UserService = require("../service/user");
const MemoryService = require("../service/memory");
const AlbumService = require("../service/album");

/**
 * get server instance
 * 
 * @return {grpc.Server}
 */
exports.getServer = function () {
  var server = new grpc.Server();
  // userService
  const userService = new UserService();
  server.addProtoService(userService.userProto.UserService.service, {
    createUser: userService.createUser,
    getMyProfile: userService.getMyProfile,
    getUserProfile: userService.getUserProfile,
    updateUserProfile: userService.updateUserProfile
  });

  // memoryService
  const memoryService = new MemoryService();
  server.addProtoService(memoryService.memoryProto.MemoryService.service, {
    getMemory: memoryService.getMemory,
    getUserMemories: memoryService.getUserMemories,
    createMemory: memoryService.createMemory,
    deleteMemory: memoryService.deleteMemory,
    updateMemory: memoryService.updateMemory,
    uploadMemoryPhotos: memoryService.uploadMemoryPhotos,
    deleteMemoryPhotos: memoryService.deleteMemoryPhotos,
    addMemoryTags: memoryService.addMemoryTags,
    deleteMemoryTags: memoryService.deleteMemoryTags
  });

  const albumService = new AlbumService();
  server.addProtoService(albumService.albumProto.AlbumService.service, {
    createAlbum: albumService.createAlbum,
    getAlbum: albumService.getAlbum,
    getMyAlbums: albumService.getMyAlbums,
    changeAlbumTitle: albumService.changeAlbumTitle,
    deleteAlbum: albumService.deleteAlbum
  });
  return server;
};