module.exports = class TagResourceBuilder {
  constructor() {
    this.tagId = 1;
    this.tagName = "hoge";
  }

  /**
   * build message Tag
   */
  build() {
    return {
      tag_id: this.tagId,
      tag_name: this.tagName
    };
  }

  /**
   * 
   * @param {Number} tagId 
   */
  withTagId(tagId) {
    this.tagId = tagId;
    return this;
  }

  /**
   * 
   * @param {String} tagName 
   */
  withTagName(tagName) {
    this.tagName = tagName;
    return this;
  }
};