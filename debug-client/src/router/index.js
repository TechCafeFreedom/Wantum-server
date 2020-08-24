import Vue from "vue";
import Router from "vue-router";
import Login from "../views/Login.vue";
import Success from "../views/Success.vue";
import firebase from "firebase/app";

Vue.use(Router);

const router = new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "login",
      component: Login
    },
    {
      path: "/success",
      name: "success",
      component: Success
    },
    {
      path: "*",
      name: "login",
      component: Login
    }
  ]
});

// 未認証の場合はログイン画面へ
router.beforeResolve((to, from, next) => {
  if (to.path == "/") {
    next();
  } else {
    firebase.auth().onAuthStateChanged(user => {
      if (user) {
        console.log("認証中");
        next();
      } else {
        console.log("未認証");
        next({ path: "/" });
      }
    });
  }
});

export default router;
