<template>
  <div class="container">
    <div class="row">
      <div class="col-md-12">
        <h2>ログイン成功</h2>
        <div class="mt-4">
          <b-button block variant="primary" @click="logout"
            >ログアウト</b-button
          >
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.mt-4 {
  margin-top: 4px;
}
</style>

<script>
import firebase from "firebase/app";
import router from "../router";

export default {
  name: "success",
  methods: {
    logout() {
      firebase
        .auth()
        .signOut()
        .then(function() {
          router.push("/");
        })
        .catch(error => {
          console.log(error.message);
          router.push("/");
        });
    }
  }
};
</script>
