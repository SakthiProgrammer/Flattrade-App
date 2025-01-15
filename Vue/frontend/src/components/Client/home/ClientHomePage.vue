<template>

  <v-container class="row-gap-4">
    <v-row align="center" justify="space-between">
      <!-- Bank Balance Section -->
      <v-col cols="12" md="6" class="d-flex align-center">
        <v-card class="pa-3" outlined color="blue lighten-5">
          <v-card-text class="text-h6 font-weight-bold">
            <a href="#" class="text-decoration-none text-primary">
              Bank Balance: ₹ {{ balance }}
            </a>
          </v-card-text>
        </v-card>
      </v-col>

      <!-- User Details Section -->
      <v-col cols="12" md="6">
        <UserDetails :Client="client" />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <PendingAndHistoryTable />
      </v-col>
    </v-row>

  </v-container>
</template>


<script>
import PendingAndHistoryTable from './components/PendingAndHistoryTable.vue';
import UserDetails from './components/ClientDetails.vue';
import EventService from '../../../Services/EventService';

export default {
  data() {
    return {
      balance: 133231,
      client: []
    };
  },
  components: {
    UserDetails,
    PendingAndHistoryTable,
  },
  methods: {

  },
  beforeMount() {
    // let userId = window.location.pathname.split('/').pop();

    let user = JSON.parse(localStorage.getItem("userRoleAndId"))
    let userId = user.user_id
    console.log('hi', userId);

    let currentUserId = '';
    let flag = false
    for (let i = 2; i < userId.length; i++) {
      if (userId[i] != 0 || flag) {
        flag = true;
        currentUserId += userId[i];
      }
    }

    // let head = { "ID": currentUserId }
    EventService.GetClientById(currentUserId)
      .then((res) => {
        if (res.data.status == "S") {
          this.client = res.data.client_details[0]
        } else {
          console.log(res.data.errMsg)
        }
      })
    console.log(currentUserId);
    // console.log(userId);
  },

}
</script>