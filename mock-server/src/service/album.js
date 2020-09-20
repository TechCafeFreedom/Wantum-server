const protoUtil = require("../utility/protoUtil");
const AlbumResourceBuilder = require("../resource/album");
const MemoryResourceBuilder = require("../resource/memory");

module.exports = class AlbumService {
  constructor() {
    const protoDescriptor = protoUtil.getProtoDescriptor("album.proto");
    this.albumProto = protoDescriptor.proto_album;
  }

  createAlbum(call, callback) {
    callback(
      null,
      new AlbumResourceBuilder()
        .wishTitle(call.request.title)
        .build()
    );
  }

  getAlbum(call, callback) {
    var memories = [];
    // TODO: 数固定きしょい
    for (var i = 0; i < 10; i++) {
      memories.push(new MemoryResourceBuilder()
        .withMemoryId(i + 1)
        .withAlbumId(call.request.album_id)
        .build());
    }
    
    callback(
      null,
      new AlbumResourceBuilder()
        .withAlbumId(call.request.album_id)
        .withMemories(memories)
        .build()
    );
  }

  getMyAlbums(call, callback) {
    var memories = [];
    // TODO: 固定値きっしょ
    for (var i = 0; i < 5; i++) {
      memories.push(new MemoryResourceBuilder()
        .withMemoryId(i + 1)
        .withAlbumId(call.request.album_id)
        // TODO: photo, tagを挟んで
        .build());
    }
    // TODO: 固定値きっしょ
    var albums = [];
    const limit = call.request.limit ? call.request.limit : 5;
    for (var i = 0; i < limit; i++) {
      albums.push(new AlbumResourceBuilder()
        .withAlbumId(i + 1)
        .withMemories(memories)
        .build());
    }
    callback(
      null,
      { albums: albums }
    );
  }

  changeAlbumTitle(call, callback) {
    // TODO: memoryがからっぽだよ〜〜〜www
    callback(
      null,
      new AlbumResourceBuilder()
        .withAlbumId(call.request.album_id)
        .wishTitle(call.request.title)
        .build()
    );
  }

  deleteAlbum(call, callback) {
    callback(null, {});
  }
};