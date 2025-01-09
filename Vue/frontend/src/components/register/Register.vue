<template>

    <v-container>
        <v-card elevation="6" class="pa-md-15 mx-lg-auto pa-sm-10" width="100%">
            <v-form>
                <v-row>
                    <h1 class="mx-auto mb-7">Sign Up</h1>
                </v-row>

                <v-row>
                    <!-- First Name -->
                    <v-col cols="12" sm="4" md="4" lg="4">
                        <v-text-field label="First Name" outlined></v-text-field>
                    </v-col>

                    <!-- Last Name -->
                    <v-col cols="12" sm="4" md="4" lg="4">
                        <v-text-field label="Last Name" outlined></v-text-field>
                    </v-col>

                    <!-- Email -->
                    <v-col cols="12" sm="4" md="4" lg="4">
                        <v-text-field label="Email" outlined required></v-text-field>
                    </v-col>
                </v-row>

                <v-row>
                    <!-- Phone Number -->
                    <v-col cols="12" sm="4" md="4" lg="4">
                        <v-text-field label="Phone Number" outlined required></v-text-field>
                    </v-col>

                    <!-- Nominee Name -->
                    <v-col cols="12" sm="4" md="4" lg="4">
                        <v-text-field label="Nominee Name" outlined></v-text-field>
                    </v-col>
                    <!-- Bank Account No -->
                    <v-col cols="12" sm="4" md="4" lg="4">
                        <v-text-field label="Bank Account Number" required outlined></v-text-field>
                    </v-col>


                </v-row>

                <v-row>
                    <v-col cols="12" sm="4" md="4" lg="4">
                        <v-select v-model="SelectedBank" :items="bankDetails.bank_details" item-text="bank_name"
                            item-value="id" outlined label="Bank Name" required @change="updateBankDetails">
                        </v-select>
                    </v-col>

                    <v-col cols="12" sm="4" md="4" lg="4">
                        <v-text-field v-model="branchName" label="Branch Name" required outlined
                            readonly></v-text-field>
                    </v-col>

                    <v-col cols="12" sm="4" md="4" lg="4">
                        <v-text-field v-model="ifscCode" label="IFSC Code" required outlined readonly></v-text-field>
                    </v-col>

                </v-row>

                <v-row>
                    <v-col cols="12">
                        <div class="d-flex justify-center mx-8">
                            <v-btn class="green" dark>Register</v-btn>
                        </div>
                    </v-col>
                </v-row>
            </v-form>
        </v-card>
        
    </v-container>
</template>



<script>
import EventService from '../../EventServices/EventService';


export default {
    data() {
        return {
            SelectedBank: "",
            bankDetails: [],
            branchName: "",
            ifscCode: ""
        }
    },

    mounted() {
        alert("hi")
        EventService.GetBankDetails()
            .then((res) => {
                this.bankDetails = res.data
            })
            .catch((err) => alert(err))
    },
    methods: {
        updateBankDetails() {
            const selectedItem = this.bankDetails.bank_details.find((bank) => bank.id === this.SelectedBank);
            console.log(selectedItem);
            if (selectedItem) {
                this.branchName = selectedItem.branch_name;
                this.ifscCode = selectedItem.ifsc_code
            }
        }
    }

}

</script>