<template>
  <v-container>
    <v-card elevation="6" class="pa-md-15 mx-lg-auto pa-sm-10" width="100%">
      <v-form>
        <v-row>
          <h1 class="mx-auto mb-7">Sign Up</h1>
        </v-row>

        <v-row>
          <v-col cols="12" sm="4" md="4" lg="4">
            <v-text-field
              v-model="registerClient.first_name"
              label="First Name"
              outlined
            ></v-text-field>
          </v-col>

          <v-col cols="12" sm="4" md="4" lg="4">
            <v-text-field
              v-model="registerClient.last_name"
              label="Last Name"
              outlined
            ></v-text-field>
          </v-col>

          <v-col cols="12" sm="4" md="4" lg="4">
            <v-text-field
              v-model="registerClient.email"
              label="Email"
              outlined
              required
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12" sm="4" md="4" lg="4">
            <v-text-field
              v-model="registerClient.phone_number"
              label="Phone Number"
              outlined
              required
            ></v-text-field>
          </v-col>

          <v-col cols="12" sm="4" md="4" lg="4">
            <v-text-field
              v-model="registerClient.nominee_name"
              label="Nominee Name"
              outlined
            ></v-text-field>
          </v-col>
          <v-col cols="12" sm="4" md="4" lg="4">
            <v-text-field
              v-model="registerClient.bank_account"
              label="Bank Account Number"
              required
              outlined
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12" sm="4" md="4" lg="4">
            <v-select
              v-model="SelectedBank"
              :items="bankDetails.bank_details"
              item-text="bank_name"
              item-value="id"
              outlined
              label="Bank Name"
              required
              @change="updateBankDetails"
            >
            </v-select>
          </v-col>

          <v-col cols="12" sm="4" md="4" lg="4">
            <v-text-field
              v-model="branchName"
              label="Branch Name"
              required
              outlined
              readonly
            ></v-text-field>
          </v-col>

          <v-col cols="12" sm="4" md="4" lg="4">
            <v-text-field
              v-model="ifscCode"
              label="IFSC Code"
              required
              outlined
              readonly
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" sm="4" md="4" lg="4">
            <v-text-field
              v-model="registerClient.pan_number"
              label="Pan Number"
              required
              outlined
            ></v-text-field>
          </v-col>
          <v-col cols="12" sm="4" md="4" lg="4">
            <v-text-field
              v-model="registerClient.password"
              :type="showPassword ? 'text' : 'password'"
              label="Password"
              required
              outlined
              append-icon="mdi-eye"
              @click:append="showPassword = !showPassword"
            ></v-text-field>
          </v-col>
          <v-col cols="12" sm="4" md="4" lg="4">
            <v-text-field
              :type="showConfirmPassword ? 'text' : 'password'"
              v-model="ConformPassword"
              @click:append="showConfirmPassword = !showConfirmPassword"
              append-icon="mdi-eye"
              label="Conform Password"
              required
              outlined
            ></v-text-field>
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12">
            <div class="d-flex justify-center mx-8">
              <v-btn class="green" @click="RegisterClient" dark>Register</v-btn>
            </div>
          </v-col>
        </v-row>
      </v-form>
    </v-card>
    <v-dialog
      v-model="dialog"
      transition="dialog-top-transition"
      max-width="500"
    >
      <template v-slot:default>
        <v-card flat>
          <v-card-title class="text-h6 text-center">User ID</v-card-title>
          <v-card-text>
            <v-text-field
              v-model="textToCopy"
              outlined
              readonly
              :append-icon="icon"
              :class="{ 'success-icon': icon === 'mdi-check-circle' }"
              @click:append="copyToClipboard"
            ></v-text-field>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" @click="navigateToLogin()">Login</v-btn>
          </v-card-actions>
        </v-card>
      </template>
    </v-dialog>
    <!-- {{ data }}
        {{ registerClient }}
        {{ registerClient.bank_id }} -->
  </v-container>
</template>

<script>
import EventService from "../../Services/EventService";

export default {
  data() {
    return {
      data: [],
      SelectedBank: 0,
      bankDetails: [],
      textToCopy: "FT0001",
      branchName: "",
      ifscCode: "",
      dialog: true,
      icon: "mdi-content-copy",
      // password: "",
      ConformPassword: "",
      showPassword: false,
      showConfirmPassword: false,
      registerClient: {
        first_name: "",
        last_name: "",
        phone_number: "",
        nominee_name: "",
        bank_account: "",
        pan_number: "",
        email: "",
        bank_id: 0,
        password: "",
      },
    };
  },

  mounted() {
    // alert("hi");
    EventService.GetBankDetails()
      .then((res) => {
        this.bankDetails = res.data;
      })
      .catch((err) => console.log(err));
  },
  methods: {
    navigateToLogin() {
      this.dialog = false;
      this.$router.push("/login");
    },
    updateBankDetails() {
      const selectedItem = this.bankDetails.bank_details.find(
        (bank) => bank.id === this.SelectedBank
      );
      console.log(selectedItem);

      if (selectedItem) {
        this.branchName = selectedItem.branch_name;
        this.ifscCode = selectedItem.ifsc_code;

        // Update registerClient.bank_id when a bank is selected
        this.registerClient.bank_id = this.SelectedBank;
      }
    },

    copyToClipboard() {
      navigator.clipboard.writeText(this.textToCopy).then(
        () => {
          this.icon = "mdi-check-circle"; // Change to check icon on success
          this.$toast?.success("Text copied to clipboard!");
          setTimeout(() => {
            this.icon = "mdi-content-copy"; // Reset back to copy icon after 2 seconds
          }, 2000);
        },
        (err) => {
          this.$toast?.error("Failed to copy text: " + err);
        }
      );
    },
    RegisterClient() {
      EventService.CreateClient(this.registerClient)
        .then((res) => {
          if (res.data.errMsg != "") {
            console.log(res.data.errMsg);
            this.data = res.data;
            this.dialog = true;
          }
        })
        .catch((err) => console.log(err));
    },
  },
};
</script>
