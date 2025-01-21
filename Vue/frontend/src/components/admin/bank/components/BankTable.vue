<template>
  <div>
    <v-card>
      <div class="d-flex justify-space-between align-center">
        <v-card-title class="text-h5">Bank Details</v-card-title>
        <v-btn color="primary" class="mr-5" @click="dialog = true"
          >Add Bank</v-btn
        >
      </div>
      <v-card-text>
        <v-data-table
          :headers="headers"
          :items="banks"
          class="elevation-1"
          fixed-header
        >
          <template v-slot:body="{ items }">
            <tbody>
              <tr v-for="(bank, index) in items" :key="bank.id">
                <td>{{ index + 1 }}</td>
                <!-- S.No -->
                <td>{{ bank.bank_name }}</td>
                <td>{{ bank.branch_name }}</td>
                <td>{{ bank.ifsc_code }}</td>
                <td>{{ bank.address }}</td>
                <td>
                  <v-btn color="red" small @click="editBank(bank)">
                    <v-icon>mdi-pencil</v-icon>
                  </v-btn>
                </td>
              </tr>
            </tbody>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>

    <v-dialog v-model="dialog" persistent max-width="600px">
      <template v-slot:default>
        <v-card>
          <v-card-title class="text-h5">Upload Bank Details</v-card-title>
          <v-form>
            <v-card-text>
              <v-text-field
                label="Bank Name"
                v-model="newBank.bank_name"
                required
              ></v-text-field>
              <v-text-field
                label="Branch Name"
                v-model="newBank.branch_name"
                required
              ></v-text-field>
              <v-text-field
                label="IFSC Code"
                v-model="newBank.ifsc_code"
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
              <v-btn text @click="dialog = cancelDialog()">Cancel</v-btn>
              <v-btn @click="addBank" color="primary">Save</v-btn>
            </v-card-actions>
          </v-form>
        </v-card>
      </template>
    </v-dialog>
  </div>
</template>

<script>
import EventService from "../../../../Services/EventService";

export default {
  data() {
    return {
      edit: false,
      dialog: false,
      newBank: {
        id: "",
        bank_name: "",
        branch_name: "",
        ifsc_code: "",
        address: "",
      },
      headers: [
        { text: "S.No", value: "sno", sortable: false },
        { text: "Bank Name", value: "bank_name" },
        { text: "Branch Name", value: "branch_name" },
        { text: "IFSC Code", value: "ifsc_code" },
        { text: "Address", value: "address" },
        { text: "Edit", value: "edit", sortable: false, align: "center" },
      ],
      banks: [],
    };
  },
  methods: {
    addBank() {
      if (this.edit) {
        EventService.UpdateBank(this.newBank)
          .then((res) => {
            if (res.data.status == "S") {
              console.log("Succefully Edit bank");
            } else {
              console.log(res.data.errMsg);
            }
          })
          .catch((err) => console.log(err));
        this.edit = false;
        this.banks.find((item) => {
          if (this.newBank.id === item.id) {
            item.bank_name = this.newBank.bank_name;
            item.branch_name = this.newBank.branchName;
            item.ifsc_code = this.newBank.ifscCode;
            item.address = this.newBank.address;
          }
        });
      } else {
        EventService.CreateBank(this.newBank)
          .then((res) => {
            if (res.data.status == "S") {
              console.log("Bank Created");
              this.banks.push(this.newBank);
            } else {
              console.log(res.data.errMsg);
            }
          })
          .catch((err) => console.log(err));
      }
      this.newBank = {
        bankName: "",
        branchName: "",
        ifscCode: "",
        address: "",
      };
      this.dialog = false;
    },

    cancelDialog() {
      this.dialog = false;
      this.newBank.id = "";
      this.newBank.bank_name = "";
      this.newBank.branch_name = "";
      this.newBank.address = "";
      this.newBank.ifsc_code = "";
    },

    editBank(bank) {
      this.newBank.id = bank.id;
      this.newBank.bank_name = bank.bank_name;
      this.newBank.branch_name = bank.branch_name;
      this.newBank.address = bank.address;
      this.newBank.ifsc_code = bank.ifsc_code;
      this.dialog = true;
      this.edit = true;
    },
  },
  beforeMount() {
    EventService.GetBankDetails("/getbanks")
      .then((res) => {
        if (res.data.status == "S") {
          this.banks = res.data.bank_details;
        } else {
          console.log("err", res.data.errMsg);
        }
      })
      .catch((err) => console.log(err));
  },
};
</script>
