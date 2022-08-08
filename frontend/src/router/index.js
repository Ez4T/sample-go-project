import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import LoginView from "../views/LoginView.vue";
import RegisterView from "../views/RegisterView.vue";
import LogoutView from "../views/LogoutView.vue";

const routes = [
  {
    path: "/",
    name: "homeView",
    component: HomeView,
  },
  {
    path: "/login",
    name: "LoginView",
    component: LoginView,
  },
  {
    path: "/register",
    name: "RegisterView",
    component: RegisterView,
  },
  {
    path: "/logout",
    name: "LogoutView",
    component: LogoutView,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const username = window.localStorage.getItem("username");
  const token = window.localStorage.getItem("token");
  const publicPages = ["LoginView", "RegisterView"];
  const authRequire = !publicPages.includes(to.name);
  if (authRequire && (token === null || username === null)) {
    alert("please login");
    next("/login");
  } else {
    next();
  }
});

export default router;
