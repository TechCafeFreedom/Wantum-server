const TagResourceBuilder = require("./tag");

module.exports = class WishCardResourceBuilder {
  constructor() {
    this.wishCardId = 1;
    this.activity = "hogehoge";
    this.description = "fugafuga";
    this.date = 1600513200;
    this.doneAt = 1600513200;
    this.place = "tsushima";
    this.tags = [];
  }

  build() {
    return {
      wish_card_id: this.wishCardId,
      activity: this.activity,
      description: this.description,
      date: this.date,
      done_at: this.doneAt,
      place: this.place,
      tags: this.tags
    };
  }

  withWishCardId(id) {
    this.wishCardId = id;
    return this;
  }

  withActivity(activity) {
    this.activity = activity;
    return this;
  }

  withDescription(description) {
    this.description = description;
    return this;
  }

  withDate(date) {
    this.date = date;
    return this;
  }

  withDoneAt(date) {
    this.doneAt = date;
    return this;
  }
  
  withPlace(place) {
    this.place = place;
    return this;
  }
  
  withTags(tags) {
    this.tags = tags;
    return this;
  }

  /**
   * 指定された数だけtagを生成
   * 
   * @param {Number} count 
   */
  withAutoGeneratedTags(count) {
    for (var i = 0; i < count; i++) {
      this.tags.push(new TagResourceBuilder()
        .withTagId(i + 1)
        .withTagName(`tag${i + 1}`)
        .build());
    }
    return this;
  }
};