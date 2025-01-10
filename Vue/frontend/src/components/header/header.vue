<template>
    <v-app-bar app elevation="3" elevate-on-scroll color="white" height="80px">
    <v-container fluid>
      <v-row align="center" no-gutters>
        <!-- Logo -->
        <v-col cols="12" md="3" sm="6" class="d-flex align-center">
          <v-img
            class="ml-5"
            lazy-src="https://picsum.photos/id/11/10/6"
            max-height="50"
            max-width="160"
            src="../../assets/logo-blue.png"
          ></v-img>
        </v-col>

        <!-- Desktop Navigation Links -->
        <v-col cols="12" md="6" class="d-none d-md-flex justify-center">
          <v-list class="d-flex justify-center">
            <v-list-item
              v-for="(link, index) in navLinks"
              :key="index"
              :link="true"
              :to="link.path"
            >
              <v-list-item-content class="px-5">
                <v-list-item-title class="h6 font-weight-medium">
                  {{ link.title }}
                </v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-col>

        <!-- Login Button -->
        <v-col cols="12" sm="6" md="3" class="d-flex justify-end">
          <v-btn link to="/login" class="primary rounded">
            Login
          </v-btn>
          <!-- Mobile Menu Icon -->
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
            userType: null, // 'user', 'admin', or null
            navLinks: [] // To hold navigation links dynamically
        };
    },
    mounted() {
        // Example: Fetch the user type from localStorage or API
        this.userType = "admin"// 'user', 'admin', or null

        // Set navigation links based on userType
        if (this.userType === 'user') {
            this.navLinks = [
                { title: 'Home', path: '/u/home' },
                { title: 'Execution', path: '/execution' }
            ];
        } else if (this.userType === 'admin') {
            this.navLinks = [
                { title: 'Home', path: '/a/home' },
                { title: 'Stock', path: '/stock' },
                { title: 'Brokerage', path: '/brokerage' },
                { title: 'Bank', path: '/bank' },
                { title: 'User', path: '/user' }
            ];
        // } else {
        //     this.navLinks = [
        //         { title: 'Login', path: '/login' }
        //     ];
        }
    }
};
</script>

