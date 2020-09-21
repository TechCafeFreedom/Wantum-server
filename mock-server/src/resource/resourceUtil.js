const TagResourceBuilder = require("./tag");
const PhotoResourceBuilder = require("./photo");
const MemoryResourceBuilder = require("./memory");
const WishCategoryResourceBuilder = require("./wishCategory");
const WishCardResourceBuilder = require("./wishCard");
const WishBoardResourceBuilder = require("./wishBoard");
const UserResourceBuilder = require("./user");

exports.autoGenerateTags = function(count) {
  var tags = [];
  for (var i = 0; i < count; i++) {
    tags.push(new TagResourceBuilder()
      .withTagId(i + 1)
      .withTagName(`tag${i + 1}`)
      .build());
  }
  return tags;
};

exports.autoGeneratePhotos = function(count) {
  var photos = [];
  for (var i = 0; i < count; i++) {
    photos.push(new PhotoResourceBuilder()
      .withPhotoId(i + 1)
      .withPhotoUrl(`http://example.com/photo/${i + 1}`)
      .build());
  }
  return photos;
};

exports.autoGenerateMemories = function(count) {
  var memories = [];
  for (var i = 0; i < count; i++) {
    memories.push(new MemoryResourceBuilder()
      .withMemoryId(i + 1)
      .withAutoGeneratedPhotos(2)
      .withAutoGeneratedTags(2)
      .build());
  }
  return memories;
};

exports.autoGenerateWishBoards = function (count) {
  var boards = [];
  for (var i = 0; i < count; i++) {
    boards.push(new WishBoardResourceBuilder()
      .withWishBoardId(i + 1)
      .withAutoGeneratedAuthors(2)
      .withAutoGeneratedWishCategories(2)
      .build());
  }
  return boards;
};

exports.autoGenerateWishCategories = function(count) {
  var categories = [];
  for (var i = 0; i < count; i++) {
    categories.push(new WishCategoryResourceBuilder()
      .withWishCategoryId(i + 1)
      .withAutoGeneratedWishCards(2)
      .build());
  }
  return categories;
};

exports.autoGenerateWishCards = function (count) {
  var cards = [];
  for (var i = 0; i < count; i++) {
    cards.push(new WishCardResourceBuilder()
      .withWishCardId(i + 1)
      .withAutoGeneratedTags(2)
      .build());
  }
  return cards;
};

exports.autoGenerateUsers = function(count) {
  var users = [];
  for (var i = 0; i < count; i++) {
    users.push(new UserResourceBuilder()
      .withUserId(i + 1)
      .withUserName(`hogehoge${i + 1}`)
      .build());
  }
  return users;
};
