<template>
  <div class="container">
    <div class="row">
      <div class="col-md-12">
        <h2>ログイン画面</h2>
        <div class="mt-2">
          <b-form-input
            v-model="email"
            type="text"
            placeholder="メールアドレス"
          />
        </div>
        <div class="mt-2">
          <b-form-input
            v-model="password"
            type="text"
            placeholder="パスワード"
          />
        </div>
        <div class="mt-2">
          <b-button block variant="primary" @click="emailLogin"
            >ログイン</b-button
          >
        </div>
        <div class="mt-2">
          <b-button block variant="primary" @click="googleLogin"
            >Google ログイン</b-button
          >
        </div>
        <div class="mt-2">
          <b-alert v-model="showError" dismissible variant="danger">{{
            errorMessage
          }}</b-alert>
        </div>
      </div>
    </div>
  </div>
</template>
<style>
.mt-2 {
  margin-top: 2px;
}
</style>

<script>
import firebase from "firebase/app";
import router from "../router";
import Cookies from 'js-cookie';

export default {
  name: "login",
  data() {
    return {
      email: "",
      password: "",
      errorMessage: "",
      showError: false
    };
  },
  methods: {
    emailLogin() {
      firebase
        .auth()
        .signInWithEmailAndPassword(this.email, this.password)
        .then(result => {
          console.log(result);
          router.push("/success");
        })
        .catch(error => {
          console.log(error);
          this.errorMessage = error.message;
          this.showError = true;
        });
    },
    googleLogin() {
      const provider = new firebase.auth.GoogleAuthProvider();

      firebase
        .auth()
        .signInWithPopup(provider)
        .then(result => {
          Cookies.set('token', result.user.xa);
          console.log(result.user.xa);
          router.push("/success");
        })
        .catch(error => {
          console.log(error);
          this.errorMessage = error.message;
          this.showError = true;
        });
    }
  }
};
</script>
