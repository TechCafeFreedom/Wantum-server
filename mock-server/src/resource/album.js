const MemoryCountResourcBuilder = require("./memoryCount");

module.exports = class AlbumResourceBuilder {
  constructor() {
    this.albumId = 1;
    this.title = "hogehoge";
    this.memories = [];
    this.memoryCount = new MemoryCountResourcBuilder().build();
  }

  build() {
    return {
      album_id: this.albumId,
      title: this.title,
      memories: {
        memories: this.memories,
        memory_count: new MemoryCountResourcBuilder()
          .withMemoriesCount(this.memories.length)
          .withPublichedCount(this.memories.length)
          .build()
      }
    };
  }

  withAlbumId(id) {
    this.albumId = id;
    return this;
  }

  wishTitle(title) {
    this.title = title;
    return this;
  }

  /**
   * buildした配列を引数に取ります
   * @param {*} memories 
   */
  withMemories(memories) {
    this.memories = memories;
    return this;
  }

  appendMemory(memory) {
    this.memories.push(memory);
    return this;
  }
};