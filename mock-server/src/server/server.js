const grpc = require("grpc");
const UserService = require("../service/user");
const MemoryService = require("../service/memory");
const AlbumService = require("../service/album");
const WishCardService = require("../service/wishCard");
const WishCategoryService = require("../service/wishCategory");
const WishBoardService = require("../service/wishBoard");
const TimelineService = require("../service/timeline");

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

  // album service
  const albumService = new AlbumService();
  server.addProtoService(albumService.albumProto.AlbumService.service, {
    createAlbum: albumService.createAlbum,
    getAlbum: albumService.getAlbum,
    getMyAlbums: albumService.getMyAlbums,
    changeAlbumTitle: albumService.changeAlbumTitle,
    deleteAlbum: albumService.deleteAlbum
  });

  // wish card service
  const wishCardService = new WishCardService();
  server.addProtoService(wishCardService.wishCardProto.WishCardService.service, {
    createWishCardInfo: wishCardService.createWishCardInfo,
    updateWishCardActivity: wishCardService.updateWishCardActivity,
    updateWishCardDescription: wishCardService.updateWishCardDescription,
    updateWishCardDate: wishCardService.updateWishCardDate,
    updateWishCardPlace: wishCardService.updateWishCardPlace,
    addWishCardTags: wishCardService.addWishCardTags,
    deleteWishCardTags: wishCardService.deleteWishCardTags
  });

  // wish category service
  const wishCategoryService = new WishCategoryService();
  server.addProtoService(wishCategoryService.wishCategoryProto.WishCategoryService.service, {
    createWishCategory: wishCategoryService.createWishCategory,
    updateWishCategoryName: wishCategoryService.updateWishCategoryName
  });

  // wish board service
  const wishBoardService = new WishBoardService();
  server.addProtoService(wishBoardService.wishBoardProto.WishBoardService.service, {
    createWishBoard: wishBoardService.createWishBoard,
    getWishBoardList: wishBoardService.getWishBoardList,
    getWishBoard: wishBoardService.getWishBoard,
    updateWishBoardName: wishBoardService.updateWishBoardName,
    updateWishBoardBackGroundImage: wishBoardService.updateWishBoardBackgroundImage,
    updateWishCategoryPriority: wishBoardService.updateWishCategoryPriority,
    deleteWishBoard: wishBoardService.deleteWishBoard
  });

  // timeline service
  const timelineService = new TimelineService();
  server.addProtoService(timelineService.timelineProto.TimeLineService.service, {
    getTimeLine: timelineService.getTimeLine
  });

  return server;
};