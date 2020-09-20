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
        .appendTag(new TagResourceBuilder().withTagId(1))
        .appendTag(new TagResourceBuilder().withTagId(2))
        .appendPhoto(new PhotoResourceBuilder())
        .build()
    );
  }

  getUserMemories(call, callback) {
    // TODO: 数の調整ができるようにしたい
    var tags = [];
    for (var i = 0; i < 3; i++) {
      tags.push(new TagResourceBuilder()
        .withTagId(i + 1)
        .withTagName("tag" + i)
        .build());
    }
    var photos = [];
    for (var i = 0; i < 5; i++) {
      photos.push(new PhotoResourceBuilder()
        .withPhotoId(i + 1)
        .build());
    }

    var memories = [];
    var memoryCount = new MemoryCountResourceBuilder();
    for(var i = 0; i < 10; i++) {
      memories.push(new MemoryResourceBuilder()
        .withMemoryId(i + 1)
        .withTags(tags)
        .withPhotos(photos)
        .withAuthor(new UserResourceBuilder()
          .withUserName(call.request.user_name)
          .build())
        .build());
      memoryCount
        .incrementMemoriesCount(1)
        .incrementPublishedCount(1);
    }
    callback(null, {
      memory_count: memoryCount.build(),
      memories: memories
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
        .build()
    );
  }

  deleteMemoryPhotos(call, callback) {
    callback(
      null,
      new MemoryResourceBuilder()
        .withMemoryId(call.request.memory_id)
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
        .withTags(tags)
        .build()
    );
  }

  deleteMemoryTags(call, callback) {
    callback(
      null,
      new MemoryResourceBuilder()
        .withMemoryId(call.request.memory_id)
        .build()
    );
  }
};