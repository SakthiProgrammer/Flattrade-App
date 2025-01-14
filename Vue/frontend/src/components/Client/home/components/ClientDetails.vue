<template>

  <v-hover>
    <template v-slot:default="{ hover }">

      <v-card class="pa-4" :elevation="hover ? 3 : 0" outlined width="450px">
        <v-card-title class="text-h5">Profile</v-card-title>
        <v-card-subtitle>{{ Client.first_name }} {{ Client.last_name }}</v-card-subtitle>
        <v-card-text>
          <v-list dense>
            <v-list-item>
              <v-list-item-content>
                <v-list-item-title>Email:</v-list-item-title>
                <v-list-item-subtitle>{{ Client.email }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>
                <v-list-item-title>Phone:</v-list-item-title>
                <v-list-item-subtitle>{{ Client.phone_number }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>
                <v-list-item-title>Pan Number:</v-list-item-title>
                <v-list-item-subtitle>{{ Client.pan_number }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>
                <v-list-item-title>Nominee Name:</v-list-item-title>
                <v-list-item-subtitle>{{ Client.nominee_name }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>
                <v-list-item-title>Bank Name:</v-list-item-title>
                <v-list-item-subtitle>{{ Client.bank_id }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>
                <v-list-item-title>Account Number:</v-list-item-title>
                <v-list-item-subtitle>{{ Client.bank_account }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>

            <v-list-item>
              <v-list-item-content>
                <v-list-item-title>Kyc Status:</v-list-item-title>
                <v-list-item-subtitle>{{ Client.kyc_is_completed }}</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
          </v-list>
        </v-card-text>
      </v-card>
    </template>
  </v-hover>
</template>

<script>
export default {
  data() {
    return {

    }
  },
  props: {
    Client: [],
  },
};
</script>