const protoUtil = require("../utility/protoUtil");
const MemoryResourceBuilder = require("../resource/memory");
const TagResourceBuilder = require("../resource/tag");
const PhotoResourceBuilder = require("../resource/photo");
const UserResourceBuilder = require("../resource/user");
const MemoryCountResourceBuilder = require("../resource/memoryCount");

module.exports = class MemoryService {
  constructor() {
    const protoDescriptor = protoUtil.getProtoDescriptor("memory.proto");
    this.memoryProto = protoDescriptor.proto_memory;
  }

  getMemory(call, callback) {
    callback(
      null,
      new MemoryResourceBuilder()
        .withMemoryId(call.request.memory_id)
        .withAutoGeneratedTags(2)
        .withAutoGeneratedPhotos(2)
        .build()
    );
  }

  getUserMemories(call, callback) {
    // TODO: 数の調整ができるようにしたい
    var memories = [];
    for(var i = 0; i < 10; i++) {
      memories.push(new MemoryResourceBuilder()
        .withMemoryId(i + 1)
        .withAutoGeneratedTags(2)
        .withAutoGeneratedPhotos(2)
        .withAuthor(new UserResourceBuilder()
          .withUserName(call.request.user_name)
          .build())
        .build());
    }
    callback(null, {
      memories: memories,
      memory_count: new MemoryCountResourceBuilder()
        .withMemoriesCount(memories.length)
        .withPublichedCount(memories.length)
        .build()
    });
  }

  createMemory(call, callback) {
    const tags = call.request.tags
      .map(tag => new TagResourceBuilder()
        .withTagId(tag.tag_id)
        .withTagName(tag.tag_name)
        .build());
    const photos = call.request.photos
      .map(photo => new PhotoResourceBuilder()
        .withPhotoId(photo.photo_id)
        .withPhotoUrl(photo.photo_url)
        .build());
    
    callback(
      null,
      new MemoryResourceBuilder()
        .withAlbumId(call.request.album_id)
        .withActivity(call.request.activity)
        .withDate(call.request.date)
        .withDescription(call.request.description)
        .withPlace(call.request.place)
        .withPhotos(photos)
        .withTags(tags)
        .build()
    );
  }

  deleteMemory(call, callback) {
    callback(null, {});
  }

  updateMemory(call, callback) {
    callback(
      null,
      new MemoryResourceBuilder()
        // QUESTION: albumIDはなし？
        .withMemoryId(call.request.memory_id)
        .withActivity(call.request.activity)
        .withDate(call.request.date)
        .withDescription(call.request.description)
        .withPlace(call.request.place)
        .withAutoGeneratedTags(2)
        .withAutoGeneratedPhotos(2)
        .build()
    );
  }

  uploadMemoryPhotos(call, callback) {
    // TODO: bufferだからなのか、リクエストが取得できない
    const photos = call.request.photo_file
      .map((file, idx) => new PhotoResourceBuilder()
        .withPhotoId(idx)
        .build());
    callback(
      null,
      new MemoryResourceBuilder()
        .withMemoryId(call.request.memory_id)
        .withPhotos(photos)
        .withAutoGeneratedTags(2)
        .build()
    );
  }

  deleteMemoryPhotos(call, callback) {
    callback(
      null,
      new MemoryResourceBuilder()
        .withMemoryId(call.request.memory_id)
        .withAutoGeneratedPhotos(2)
        .withAutoGeneratedTags(2)
        .build()
    );

  }

  addMemoryTags(call, callback) {
    const tags = call.request.tag_names
      .map((tag, idx) => new TagResourceBuilder()
        .withTagId(idx + 1)
        .withTagName(tag)
        .build());
    callback(
      null,
      new MemoryResourceBuilder()
        .withMemoryId(call.request.memory_id)
        .withAutoGeneratedPhotos(2)
        .withTags(tags)
        .build()
    );
  }

  deleteMemoryTags(call, callback) {
    callback(
      null,
      new MemoryResourceBuilder()
        .withMemoryId(call.request.memory_id)
        .withAutoGeneratedPhotos(2)
        .build()
    );
  }
};