module.exports = class PhotoResourceBuilder{
  constructor() {
    this.photoId = 1;
    this.photoUrl = "http://example.com/hogehoge.png";
  }

  build() {
    return {
      photo_id: this.photoId,
      photo_url: this.photoUrl
    };
  }

  withPhotoId(photoId) {
    this.photoId = photoId;
    return this;
  }

  withPhotoUrl(url) {
    this.photoUrl = url;
    return this;
  }
};