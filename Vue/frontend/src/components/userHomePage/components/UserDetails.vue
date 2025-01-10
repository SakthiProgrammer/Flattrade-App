<template>
    
    <v-hover>
<template v-slot:default="{ hover }">

          <v-card class="pa-4" :elevation="hover ? 3 : 0" outlined width="450px">
            <v-card-title class="text-h5">Profile</v-card-title>
            <v-card-subtitle>{{ user.name }}</v-card-subtitle>
            <v-card-text>
              <v-list dense>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>Email:</v-list-item-title>
                    <v-list-item-subtitle>{{ user.email }}</v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>Phone:</v-list-item-title>
                    <v-list-item-subtitle>{{ user.phone }}</v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>Pan Number:</v-list-item-title>
                    <v-list-item-subtitle>{{ user.pan_card }}</v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>Nominee Name:</v-list-item-title>
                    <v-list-item-subtitle>{{ user.nominee_name }}</v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>Bank Name:</v-list-item-title>
                    <v-list-item-subtitle>{{ user.bank_name }}</v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>Account Number:</v-list-item-title>
                    <v-list-item-subtitle>{{ user.bank_account }}</v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>Kyc Status:</v-list-item-title>
                    <v-list-item-subtitle>{{ user.kyc_completed }}</v-list-item-subtitle>
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
        user: {
          name: "John Doe",
          email: "john.doe@example.com",
          phone: "+123456789",
          pan_card : "SEF1312KJ",
          nominee_name: "sakthivel",
          bank_name: "UNION BANK",
          bank_account :"123kj131k2313",
          kyc_completed :"Pending"
        },
      };
    },
  };
  </script>
  