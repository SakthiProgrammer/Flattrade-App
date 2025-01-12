<template>
  <div>
    <v-card>
      <!-- Action Button to Open Dialog -->
      <div class="d-flex justify-space-between align-center">
        <v-card-title class="text-h5">User Details</v-card-title>
        <v-btn color="primary" class="mr-5" @click="dialog = true">Add User</v-btn>
      </div>
      <v-card-text>
        <v-data-table
          :headers="headers"
          :items="users"
          class="elevation-1"
          fixed-header
        >
          <!-- S.No. Column -->
          <template v-slot:body="{ items }">
            <tbody>
              <tr v-for="(user, index) in items" :key="user.id">
                <td>{{ index + 1 }}</td> <!-- S.No Column -->
                <td>{{ user.userName }}</td>
                <td>{{ user.role }}</td>
                <td>{{ user.createdAt }}</td>
                <td>{{ user.status }}</td>
              </tr>
            </tbody>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>

    <!-- Dialog to Upload User Details -->
    <v-dialog
      v-model="dialog"
      persistent
      max-width="600px"
    >
      <template v-slot:default>
        <v-card>
          <v-card-title class="text-h5">Add User Details</v-card-title>
          <v-form @submit.prevent="submitUserDetails">
            <v-card-text>
              <v-text-field
                label="User Name"
                v-model="newUser.userName"
                required
              ></v-text-field>
              <v-select
                label="Role"
                v-model="newUser.role"
                :items="['Back Officer', 'Biller', 'Approver']"
                required
              ></v-select>
              <v-text-field
                label="Created At"
                v-model="newUser.createdAt"
                type="date"
                required
              ></v-text-field>
              <v-select
                label="Status"
                v-model="newUser.status"
                :items="['ACTIVE', 'INACTIVE']"
                required
              ></v-select>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn text @click="dialog = false">Cancel</v-btn>
              <v-btn type="submit" color="primary">Save</v-btn>
            </v-card-actions>
          </v-form>
        </v-card>
      </template>
    </v-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      dialog: false,
      newUser: {
        userName: '',
        role: '',
        createdAt: '',
        status: '',
      },
      headers: [
        { text: 'S.No.', value: 'sno', sortable: false, align: 'start' },
        { text: 'User Name', value: 'userName' },
        { text: 'Role', value: 'role' },
        { text: 'Created At', value: 'createdAt' },
        { text: 'Status', value: 'status' },
      ],
      users: [
        {
          id: 1,
          userName: 'John Doe',
          role: 'Back Officer',
          createdAt: '2025-01-01',
          status: 'ACTIVE',
        },
        {
          id: 2,
          userName: 'Jane Smith',
          role: 'Biller',
          createdAt: '2025-01-05',
          status: 'INACTIVE',
        },
        {
          id: 3,
          userName: 'Mike Johnson',
          role: 'Approver',
          createdAt: '2025-01-10',
          status: 'ACTIVE',
        },
      ],
    };
  },
  methods: {
    submitUserDetails() {
      this.users.push({ ...this.newUser, id: Date.now() });
      this.dialog = false;
      this.newUser = { userName: '', role: '', createdAt: '', status: '' };
    },
  },
};
</script>

<style scoped>
.v-data-table {
  width: 100%;
}
</style>
