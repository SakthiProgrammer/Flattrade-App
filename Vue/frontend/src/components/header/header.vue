<template>
  <v-app-bar app elevation="3" elevate-on-scroll color="white" height="80px">
    <v-container fluid>
      <v-row align="center" no-gutters>

        <v-col cols="12" md="3" sm="6" class="d-flex align-center">
          <v-img class="ml-5" lazy-src="https://picsum.photos/id/11/10/6" max-height="50" max-width="160"
            src="../../assets/logo-blue.png"></v-img>
        </v-col>


        <v-col cols="12" md="6" class="d-none d-md-flex justify-center">
          <v-list class="d-flex justify-center">
            <v-list-item v-for="(link, index) in navLinks" :key="index" @click="handleComponents(link)"
              class="clickable">
              <v-list-item-content class="px-5">
                <v-list-item-title class="h6 font-weight-medium">
                  {{ link.title }}
                </v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-col>


        <v-col cols="12" sm="6" md="3" class="d-flex justify-end">
          <div v-show="userType == null">
            <v-btn link to="/login" class="primary rounded">
              Login
            </v-btn>
          </div>
          <div v-show="userType != null">
            <v-btn link to="/login" @click="ClearStore()" class="primary rounded">
              Logout
            </v-btn>
          </div>

          <v-app-bar-nav-icon class="d-md-none" @click="drawer = !drawer" />
        </v-col>
      </v-row>
    </v-container>
  </v-app-bar>
</template>

<script>
export default {
  data() {
    return {
      drawer: false,
      userType: "", // 'user', 'admin', or null
      navLinks: [],
    };
  },
  mounted() {

    this.userType = this.$store.state.role;

    const savedComponent = localStorage.getItem("currentComponent");
    this.currentComponent = savedComponent || "Home"; // Default to "Home" if nothing is saved

    if (this.userType === "user") {
      this.navLinks = [
        { title: "Home", path: "/c/home" },
        { title: "Execution" },
      ];
    } else if (this.userType === "admin") {
      this.navLinks = [
        { title: "Home", path: "/a/home" },
        { title: "Stock" },
        { title: "Brokerage" },
        { title: "Bank" },
        { title: "User" },
      ];
    }
  },
  methods: {
    handleComponents(link) {
      this.$emit("handleComponent", link.title); // Emit the current component
    },
    ClearStore() {
      this.$store.commit("setRole", null);
    }
  },

};
</script>
