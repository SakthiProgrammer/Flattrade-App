<template>

    <v-container class="row-gap-4" >
        <v-row align="center" justify="space-between">
          <!-- Bank Balance Section -->
          <v-col cols="12" md="6" class="d-flex align-center">
            <v-card class="pa-3" outlined color="blue lighten-5">
              <v-card-text class="text-h6 font-weight-bold">
                <a href="#" class="text-decoration-none text-primary">
                  Bank Balance: ₹ {{balance }}
                </a>
              </v-card-text>
            </v-card>
          </v-col>
    
          <!-- User Details Section -->
          <v-col cols="12" md="6">
            <UserDetails />
          </v-col>
        </v-row>
        <v-row>
            <v-col>
            <PendingAndHistoryTable/>
            </v-col>
        </v-row>
    
      </v-container>
    </template>
    
    
    <script>
    import PendingAndHistoryTable from './components/PendingAndHistoryTable.vue';
    import UserDetails from './components/ClientDetails.vue';
    
    export default {
        data(){
            return {
            balance : 133231,
            };
        },
        components : {
            UserDetails,
            PendingAndHistoryTable,
        }
    }
    </script>