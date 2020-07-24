import Vue from 'vue'
import VueHead from 'vue-head'
import App from './App.vue'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import firebase from 'firebase/app'
import 'firebase/auth'
import router from './router'

Vue.use(BootstrapVue)
Vue.use(VueHead)

Vue.config.productionTip = false

const firebaseConfig = {
  apiKey: "AIzaSyBRrr6yKCKGgS-6BKBojjWnZZ_pAmXSXQo",
  authDomain: "wantum.firebaseapp.com",
  databaseURL: "https://wantum.firebaseio.com",
  projectId: "wantum",
  storageBucket: "wantum.appspot.com",
  messagingSenderId: "1050487292220",
  appId: "1:1050487292220:web:cb377970dc662f04659e39",
  measurementId: "G-ZN4TZ7FF0R"
};

firebase.initializeApp(firebaseConfig);

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')