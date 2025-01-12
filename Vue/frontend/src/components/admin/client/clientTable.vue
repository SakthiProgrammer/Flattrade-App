<template>
  <v-card dense elevation="2">
    <v-card-title class="text-h5">Client Details</v-card-title>
    <v-card-text>
      <v-data-table
        :headers="headers"
        :items="clients"
        class="elevation-1"
        dense
        fixed-header
        style="overflow-x: auto;"
        item-value="client_id"
      >
        <!-- S.No. Column -->
        <template v-slot:item.s_no="{ index }">
          {{ index + 1 }}
        </template>

        <!-- KYC Status as Chip -->
        <template v-slot:item.kyc_is_completed="{ item }">
          <v-chip
            :color="item.kyc_is_completed ? 'green' : 'red'"
            dark
            small
          >
            {{ item.kyc_is_completed ? 'Completed' : 'Pending' }}
          </v-chip>
        </template>

        <!-- Action Buttons -->
        <template v-slot:item.action="{ item }">
          <v-btn
            color="green"
            small
            @click="approveClient(item)"
          >
            Approve
          </v-btn>
          <v-btn
            color="red"
            small
            class="ml-2"
            @click="rejectClient(item)"
          >
            Reject
          </v-btn>
        </template>
      </v-data-table>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      headers: [
        { text: "S.No.", value: "s_no", align: "start", sortable: false },
        { text: "Client ID", value: "client_id", align: "start" },
        { text: "Phone Number", value: "phone_number" },
        { text: "First Name", value: "first_name" },
        { text: "Last Name", value: "last_name" },
        { text: "PAN Number", value: "pan_number" },
        { text: "Nominee Name", value: "nominee_name" },
        { text: "KYC Status", value: "kyc_is_completed" },
        { text: "Bank Account", value: "bank_account" },
        { text: "Email", value: "email" },
        { text: "Action", value: "action", sortable: false },
      ],
      clients: [
        {
          client_id: 101,
          phone_number: "9876543210",
          first_name: "John",
          last_name: "Doe",
          pan_number: "ABCDE1234F",
          nominee_name: "Jane Doe",
          kyc_is_completed: true,
          bank_account: "1234567890",
          email: "john.doe@example.com",
        },
        {
          client_id: 102,
          phone_number: "8765432109",
          first_name: "Alice",
          last_name: "Smith",
          pan_number: "XYZAB5678C",
          nominee_name: "Bob Smith",
          kyc_is_completed: false,
          bank_account: "0987654321",
          email: "alice.smith@example.com",
        },
      ],
    };
  },
  methods: {
    approveClient(client) {
      // Logic for approving the client
      console.log("Approved client:", client);
      client.kyc_is_completed = true; // Example: Update the KYC status
    },
    rejectClient(client) {
      // Logic for rejecting the client
      console.log("Rejected client:", client);
      client.kyc_is_completed = false; // Example: Update the KYC status
    },
  },
};
</script>

<style scoped>
.v-data-table {
  width: 100%;
}
</style>
