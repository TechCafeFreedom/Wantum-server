/**
 * dummy builder for user
 * 最後にbuildするとそれっぽいレスポンスができます
 */
module.exports = class UserResourceBuilder {
  constructor() {
    this.userId = 1;
    this.name = "hogehoge";
    this.userName = "fugafuga";
    this.thumbnail = "hogefuga.png";
    this.bio = "hogefuga";
    this.gender = "MAN";
    this.place = "tsushima";
    this.birth = 1600513200;
    this.phone = "01234567890";
  }

  build() {
    return {
      user_id: this.userId,
      name: this.name,
      user_name: this.userName,
      thumbnail: this.thumbnail,
      bio: this.bio,
      gender: this.gender,
      place: this.place,
      birth: this.birth
    };
  }

  withUserId(userId) {
    this.userId = userId;
    return this;
  }

  withName(name) {
    this.name = name;
    return this;
  }

  withUserName(userName) {
    this.userName = userName;
    return this;
  }

  withThumbnail(thumbnail) {
    this.thumbnail = thumbnail;
    return this;
  }
  
  withBio(bio) {
    this.bio = bio;
    return this;
  }
  
  withGender(gender) {
    switch(gender) {
      case GENDER.MAN:
        this.gender = GENDER.MAN;
        break;
      case GENDER.WOMAN:
        this.gender = GENDER.WOMAN;
        break;
      default:
        this.gender = GENDER.UNKNOWN;
    }

    return this;
  }
  
  withPlace(place) {
    this.place = place;
    return this;
  }
  
  withBirth(birth) {
    this.birth = birth;
    return this;
  }

  withPhone(phone) {
    this.phone = phone;
    return this;
  }
};

var GENDER = {
  MAN: "MAN",
  WOMAN: "WOMAN",
  UNKNOWN: "UNKNOWN"
};