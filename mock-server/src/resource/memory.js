const UserResourceBuilder = require("./user");

module.exports = class MemoryResourceBuilder {
  constructor() {
    this.memoryId = 1;
    this.activity = "hogehoge";
    this.date = 1600513200;
    this.description = "fugafuga";
    this.place = "hogefuga";
    this.albumId = 1;
    this.author = new UserResourceBuilder().build();
    this.photos = [];
    this.tags = [];
  }

  build() {
    return {
      memory_id: this.memoryId,
      activity: this.activity,
      date: this.date,
      description: this.description,
      place: this.place,
      author: this.author,
      photos: this.photos,
      tags: this.tags
    };
  }

  /**
   * 
   * @param {Number} memoryId 
   */
  withMemoryId(memoryId) {
    this.memoryId = memoryId;
    return this;
  }
  
  withActivity(activity) {
    this.activity = activity;
    return this;
  }

  withDate(date) {
    this.date = date;
    return this;
  }
  
  withDescription(description) {
    this.description = description;
    return this;
  }
  
  withPlace(place) {
    this.place = place;
    return this;
  }

  withAlbumId(albumId) {
    this.albumId = albumId;
    return this;
  }
  
  /**
   * buildした後のuserを引数に取ります
   * 
   * @param {*} user
   */
  withAuthor(user) {
    this.author = user;
    return this;
  }
  
  /**
   * 全ての要素をbuildした状態にして下さい
   * 
   * @param {*} photos 
   */
  withPhotos(photos) {
    this.photos = photos;
    return this;
  }
  
  /**
   * 全ての要素をbuildした状態にしてください
   * 
   * @param {*} tags 
   */
  withTags(tags) {
    this.tags = tags;
    return this;
  }

  /**
   * buildした後のtagを引数に取ります
   * @param {*} tag 
   */
  appendTag(tag) {
    this.tags.push(tag);
    return this;
  }

  /**
   * buildした後のphotoを引数に取ります
   * @param {*} photo 
   */
  appendPhoto(photo) {
    this.photos.push(photo);
    return this;
  }
};