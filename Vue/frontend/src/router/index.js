import Vue from "vue";
import VueRouter from "vue-router";
import Login from "../views/Login.vue";
import Register from "../views/Register.vue";
import ClientHomePage from "../views/ClientHomePage.vue";
import AdminDashboard from "../views/AdminDashboard.vue";
import UserHome from "../views/UserHome.vue";
import PageNotFound from "../components/PageNotFound.vue";

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
    path: "/c/home",
    name: "ClientHomePage",
    component: ClientHomePage,
  },
  {
    path: "/a/home",
    name: "AdminLandingPage",
    component: AdminDashboard,
  },
  {
    path: "/u/home",
    name: "User",
    component: UserHome,
  },
  {
    path: "*",
    name: "NotFoundPage",
    component: PageNotFound,
  },
  /*
  {
    path: "/appr/home",
    name: "Bank",
    // component: Bank,
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
  }, */
];

const router = new VueRouter({
  routes,
  mode: "history",
});

export default router;
