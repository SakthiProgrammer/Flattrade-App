<template>
  <div>
    <v-container>
      <v-card
        class="pa-md-15 elevation-3 rounded-lg mx-lg-auto pa-sm-10"
        width="500px"
      >
        <v-row>
          <h1 class="mx-auto mb-7">Sign In</h1>
        </v-row>
        <v-row>
          <!-- <v-select v-model="name" outlined item-text="name" item-value="name" label="User ID"></v-select> -->
          <v-text-field
            v-model="LoginData.user_id"
            label="Admin Id"
            outlined
          ></v-text-field>
        </v-row>
        <v-row>
          <v-text-field
            v-model="LoginData.password"
            label="Password"
            outlined
          ></v-text-field>
        </v-row>
        <v-row>
          <v-select
            v-model="head.ROLE"
            :items="roles"
            item-text="name"
            item-value="id"
            outlined
            label="Role"
          ></v-select>
        </v-row>

        <v-row>
          <v-col cols="12">
            <div class="d-flex justify-center mx-8">
              <v-btn class="green" @click="Login" dark>Login</v-btn>
            </div>
          </v-col>
        </v-row>
      </v-card>
      <!-- {{ head }} -->
    </v-container>
  </div>
</template>

<script>
import EventService from "../../Services/EventService";
export default {
  data() {
    return {
      LoginData: { user_id: "", password: "" },
      // role: null,
      head: { ROLE: "A" },
      roles: [
        { id: "A", name: "Admin" },
        { id: "BO", name: "Back Officer" },
        { id: "B", name: "Biller" },
        { id: "APPR", name: "Approver" },
      ],
    };
  },
  destroyed() {
    // alert("i");
    this.LoginData.user_id = "";
    this.LoginData.password = "";
  },

  methods: {
    // changeRole(selectedRole) {
    //     this.head.ROLE = selectedRole;
    // },
    Login() {
      if (this.head.ROLE == "A") {
        EventService.LoginAdmin(this.LoginData, this.head)
          .then((res) => {
            if (res.data.status == "E") {
              console.log(res.data.errMsg);
            } else if (res.data.status == "S") {
              console.log("Login Successfully");
              // localStorage.setItem("")
              let user = {
                role: this.head.ROLE,
                user_id: this.LoginData.user_id,
              };
              // this.$store.commit("setRole", "admin");
              localStorage.setItem("userRoleAndId", JSON.stringify(user));
              this.$router.push("/a/home");
            }
            console.log(this.head);
          })
          .catch((err) => console.log(err));
      } else {
        let header = { ROLE: "U", USERROLE: this.head.ROLE };
        EventService.LoginUser(this.LoginData, header)
          .then((res) => {
            if (res.data.status == "E") {
              console.log(res.data.errMsg);
            } else {
              console.log("Logined Successfully");
              // localStorage.setItem("")
              let user = {
                role: this.head.ROLE,
                user_id: this.LoginData.user_id,
              };
              this.$store.commit("setRole", "admin");
              localStorage.setItem("userRoleAndId", JSON.stringify(user));
              this.$router.push("/a/home");
            }
            console.log(this.head);
          })
          .catch((err) => console.log(err));
      }
    },
  },
};
</script>
