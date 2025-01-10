<template>
    <div>
      <v-card>
        <!-- Action Button to Open Dialog -->
        <div class="d-flex justify-space-between align-center">
          <v-card-title class="text-h5">Bank Details</v-card-title>
          <v-btn color="primary" class="mr-5" @click="dialog = true">Add Bank</v-btn>
        </div>
        <v-card-text>
          <v-data-table
            :headers="headers"
            :items="banks"
            class="elevation-1"
            fixed-header
          >
            <!-- Bank Details -->
          </v-data-table>
        </v-card-text>
      </v-card>
  
      <!-- Dialog to Upload Bank Details -->
      <v-dialog
        v-model="dialog"
        persistent
        max-width="600px"
      >
        <template v-slot:default>
          <v-card>
            <v-card-title class="text-h5">Upload Bank Details</v-card-title>
            <v-form @submit.prevent="submitBankDetails">
              <v-card-text>
                <v-text-field
                  label="Bank Name"
                  v-model="newBank.bankName"
                  required
                ></v-text-field>
                <v-text-field
                  label="Branch Name"
                  v-model="newBank.branchName"
                  required
                ></v-text-field>
                <v-text-field
                  label="IFSC Code"
                  v-model="newBank.ifscCode"
                  required
                ></v-text-field>
                <v-text-field
                  label="Address"
                  v-model="newBank.address"
                  required
                ></v-text-field>
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
        newBank: {
          bankName: '',
          branchName: '',
          ifscCode: '',
          address: '',
        },
        headers: [
          { text: 'Bank Name', value: 'bankName' },
          { text: 'Branch Name', value: 'branchName' },
          { text: 'IFSC Code', value: 'ifscCode' },
          { text: 'Address', value: 'address' },
        ],
        banks: [
          { id: 1, bankName: 'Bank A', branchName: 'Branch A1', ifscCode: 'IFSC001', address: 'Address A' },
          { id: 2, bankName: 'Bank B', branchName: 'Branch B1', ifscCode: 'IFSC002', address: 'Address B' },
        ],
      };
    },
    methods: {
      submitBankDetails() {
        this.banks.push({ ...this.newBank, id: Date.now() });
        this.dialog = false;
        this.newBank = { bankName: '', branchName: '', ifscCode: '', address: '' };
      },
    },
  };
  </script>
  