<template>
  <div>
    <v-card>
      <!-- Action Button to Open Dialog -->
      <div class="d-flex justify-space-between align-center">
        <v-card-title class="text-h5">User Details</v-card-title>
        <v-btn color="primary" class="mr-5" @click="dialog = true">Add User</v-btn>
      </div>
      <v-card-text>
        <v-data-table :headers="headers" :items="users" class="elevation-1" fixed-header>
          <!-- S.No. Column -->
          <template v-slot:body="{ items }">
            <tbody>
              <tr v-for="(user, index) in items" :key="user.id">
                <td>{{ index + 1 }}</td> <!-- S.No Column -->
                <td>{{ user.user_name }}</td>
                <td>{{ user.role }}</td>
                <td>{{ user.password }}</td>
                <td>{{ user.status }}</td>
                <td>
                  <v-btn color="red" small @click="editUser(user)">
                    <v-icon>mdi-pencil</v-icon>
                  </v-btn>
                </td>
              </tr>
            </tbody>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>

    <!-- Dialog to Upload User Details -->
    <v-dialog v-model="dialog" persistent max-width="600px">
      <template v-slot:default>
        <v-card>
          <v-card-title class="text-h5">Add User Details</v-card-title>
          <v-form>
            <v-card-text>
              <v-text-field label="User Name" v-model="newUser.user_name" required></v-text-field>
              <v-select label="Role" v-model="newUser.role" :items="['Back Officer', 'Biller', 'Approver']"
                required></v-select>
              <v-text-field label="Password" v-model="newUser.password" required></v-text-field>
              <v-select label="Status" v-model="newUser.status" :items="['ACTIVE', 'INACTIVE']" required></v-select>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn text @click="cancelDialog">Cancel</v-btn>
              <v-btn @click="addUser" color="primary">Save</v-btn>
            </v-card-actions>
          </v-form>
        </v-card>
      </template>
    </v-dialog>
  </div>
</template>

<script>
import EventService from '../../../Services/EventService';

export default {
  data() {
    return {
      dialog: false,
      newUser: {
        id: "",
        user_name: '',
        role: '',
        // createdAt: '',
        password: '',
        status: '',
      },
      headers: [
        { text: 'S.No.', value: 'sno', sortable: false, align: 'start' },
        { text: 'User Name', value: 'user_name' },
        { text: 'Role', value: 'role' },
        { text: 'password', value: 'password' },
        { text: 'Status', value: 'status' },
        { text: "Edit", value: "edit", sortable: false, align: "center" },

      ],
      users: [],
      edit: false,
    };
  },
  methods: {

    addUser() {
      if (this.edit) {
        EventService.UpdateUser(this.newUser)
          .then((res) => {
            if (res.data.status == "S") {
              console.log("Success Edit");

            } else {
              console.log(res.data.errMsg)
            }
          }).catch((err) => console.log(err)
          );
        this.users.find((item) => {
          if (this.newUser.id === item.id) {
            item.user_name = this.newUser.user_name;
            item.password = this.newUser.password;
            item.role = this.newUser.role;
            item.status = this.newUser.status;
          }
        });
        this.edit = false
      } else {
        EventService.CreateUser(this.newUser)
          .then((res) => {
            if (res.data.status == "S") {
              console.log("User Created");
              this.users.push(this.newUser)

            } else {
              console.log("error", res.data.errMsg);
            }
          }).catch((err) => console.log(err))
      }
      this.clearUser();
      this.dialog = false;
    },

    cancelDialog() {
      this.dialog = false;
      this.clearUser()
      if (this.edit) {
        this.edit = !this.edit;
      }
    },

    clearUser() {
      this.newUser = { id: "", user_name: '', role: '', password: '', status: '' };
    },

    editUser(User) {
      this.edit = true;
      this.newUser.id = User.id;
      this.newUser.user_name = User.user_name;
      this.newUser.role = User.role;
      this.newUser.password = User.password;
      this.newUser.status = User.status;
      this.dialog = true;
    },
  },
  beforeMount() {
    EventService.GetUsers()
      .then((res) => {
        if (res.data.status == "S") {
          this.users = res.data.user_details
        } else {
          console.log("error", res.data.errMsg);
        }
      }).catch((err) => console.log(err))
  }
};
</script>
