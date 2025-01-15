<template>
  <v-container>
    <v-card class="pa-md-15 mx-lg-auto pa-sm-10" width="500px">
      <v-row>
        <h1 class="mx-auto mb-7">Sign In</h1>
      </v-row>
      <v-row>
        <v-text-field
          v-model="Data.user_id"
          label="User Id"
          outlined
        ></v-text-field>
      </v-row>
      <v-row>
        <v-text-field
          v-model="Data.password"
          label="Password"
          outlined
        ></v-text-field>
      </v-row>
      <v-row>
        <h1 class="body-2">
          If Don't Have an Account
          <router-link to="/register"> Create? </router-link>
        </h1>
      </v-row>

      <v-row>
        <v-col cols="12">
          <div class="d-flex justify-center mx-8">
            <v-btn @click="Login" class="green" dark>Login</v-btn>
          </div>
        </v-col>
      </v-row>
    </v-card>
  </v-container>
</template>

<script>
import EventService from "../../../Services/EventService";

export default {
  data() {
    return {
      Data: { user_id: "", password: "" },
    };
  },
  methods: {
    Login() {
      
      EventService.LoginClient(this.Data)
        .then((res) => {
          if (res.data.status == "E") {
            console.log(res.data.errMsg);
          } else {
            console.log("Logined Successfully");
            this.$store.commit("setRole", "user");
            let currentUser = { role: "client", user_id: this.Data.user_id };
            localStorage.setItem("userRoleAndId", JSON.stringify(currentUser));
            // this.$router.push(`/c/home/${this.Data.user_id}`)
            this.$router.push("/c/home");
          }
        })
        .catch((err) => console.log(err));
    },
  },
};
</script>
