import Vue from "vue";
import VueRouter from "vue-router";
import Login from "../views/Login.vue";
import Register from "../views/Register.vue";
import UserLandingPage from "../views/UserLandingPage.vue";
import Execution from "../views/Execution.vue";
import AdminLandingPage from "../views/AdminLandingPage.vue";
import Stock from "../views/Stock.vue";
import Bank from "../views/Bank.vue";
import Brokerage from "../views/Brokerage.vue";
import User from "../views/User.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
  {
    path: "/register",
    name: "Register",
    component: Register,
  },
  {
    path: "/u/home",
    name: "UserLandingPage",
    component: UserLandingPage,
  },
  {
    path: "/a/home",
    name: "AdminLandingPage",
    component: AdminLandingPage,
  },
  {
    path: "/execution",
    name: "Execution",
    component: Execution,
  },
  {
    path: "/stock",
    name: "Stock",
    component: Stock,
  },
  {
    path: "/bank",
    name: "Bank",
    component: Bank,
  },
  {
    path: "/brokerage",
    name: "Brokerage",
    component: Brokerage,
  },
  {
    path: "/user",
    name: "User",
    component: User,
  },
];

const router = new VueRouter({
  routes,
  mode: "history",
});

export default router;
