<template>
  <v-card elevation="4">
    <v-card-title class="text-h5">Client Details</v-card-title>
    <v-card-text>
      <v-data-table :headers="headers" :items="clients" class="elevation-1" fixed-header style="overflow-x: auto;"
        item-value="client_id">

        <template v-slot:item.s_no="{ index }">
          {{ index + 1 }}
        </template>
        <template v-slot:item.name ="{ item }">
          {{item.first_name +" "+ item.last_name}}
        </template>


        <!-- <template v-slot:item.kyc_is_completed="{ item }">
          <v-chip :color="item.kyc_is_completed ? 'green' : 'red'" dark small>
            {{ item.kyc_is_completed ? 'Completed' : 'Pending' }}
          </v-chip>
        </template> -->


        <template v-slot:item.action="{ item }">
          <div class="d-flex ">
            <v-btn color="green" small @click="approveClient(item)">
              Approve
            </v-btn>
            <v-btn color="red" small class="ml-2" @click="rejectClient(item)">
              Reject
            </v-btn>
          </div>
        </template>
      </v-data-table>
    </v-card-text>
  </v-card>
</template>

<script>
import EventService from '../../../Services/EventService';

export default {
  data() {
    return {
      headers: [
        { text: "S.No.", value: "s_no", align: "start", sortable: false },
        { text: "Client ID", value: "unique_id", align: "start" },
        { text: "Name", value: "name" },
        { text: "Phone Number", value: "phone_number" },
        // { text: "Last Name", value: "last_name" },
        { text: "PAN Number", value: "pan_number" },
        // { text: "Nominee Name", value: "nominee_name" },
        //{ text: "KYC Status", value: "kyc_is_completed" },
        { text: "Bank Account", value: "bank_account" },
        { text: "Email", value: "email" },
        { text: "Action", value: "action", sortable: false },
      ],
      clients: [],
    };

  },
  methods: {
    approveClient(client) {
      client.kyc_is_completed = "Approved"
      EventService.ApproveClient(client)
        .then((res) => {
          if (res.data.status == "S") {
            console.log("Approved");
            this.removeClient(client.client_id)
          } else {
            console.log("Error", res.data.errMsg);
          }
        }).catch((err => console.log(err)))
    },
    rejectClient(client) {
      // Logic for rejecting the client
      // console.log("Rejected client:", client);
      client.kyc_is_completed = "Rejected"; // Example: Update the KYC status
      EventService.RejectClient(client)
        .then((res) => {
          if (res.data.status == "S") {
            console.log("Rejected")
            this.removeClient(client.client_id)
          } else {
            console.log("Error", res.data.errMsg);
          }

        }).catch((err) => console.log(err))
    },

    removeClient(clientId) {
      this.clients = this.clients.filter(currentClient => currentClient.client_id !== clientId)
    }
  },
  beforeMount() {
    EventService.GetClientDetails()
      .then((res) => {
        if (res.data.status == "S") {
          this.clients = res.data.client_details
        } else {
          console.log(res.data.errMsg)
        }
      }).catch((err) => console.log(err))
  }
};
</script>
