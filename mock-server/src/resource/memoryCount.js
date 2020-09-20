module.exports = class MemoryCountResourceBuilder {
  constructor() {
    this.memoriesCount = 0;
    this.publishedCount = 0;
  }

  build() {
    return {
      memories_count: this.memoriesCount,
      published_count: this.publishedCount
    };
  }

  withMemoriesCount(count) {
    this.memoriesCount = count;
    return this;
  }

  withPublichedCount(count) {
    this.publishedCount = count;
    return this;
  }

  incrementMemoriesCount(incr) {
    this.memoriesCount += incr;
    return this;
  }

  incrementPublishedCount(incr) {
    this.publishedCount += incr;
    return this;
  }
};