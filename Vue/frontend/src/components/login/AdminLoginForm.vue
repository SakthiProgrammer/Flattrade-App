<template>
    <v-container>
        <v-card class="pa-md-15 mx-lg-auto pa-sm-10" width="500px">
            <v-row>
                <h1 class="mx-auto mb-7">Sign In</h1>
            </v-row>
            <v-row>
                <!-- <v-select v-model="name" outlined item-text="name" item-value="name" label="User ID"></v-select> -->
                <v-text-field label="Admin Id" outlined></v-text-field>

            </v-row>
            <v-row>
                <v-text-field label="Password" outlined></v-text-field>
            </v-row>
            <v-row>
                <v-select v-model="name" :items="roles" item-text="name" item-value="id" outlined
                    label="Role"></v-select>
            </v-row>

            <v-row>
                <v-col cols="12">
                    <div class="d-flex justify-center mx-8">
                        <v-btn class="green " dark>Login</v-btn>
                    </div>
                </v-col>
            </v-row>
        </v-card>
    </v-container>
</template>

<script>
export default {

    data() {
        return {
            name: null,
            roles: [
                { id: 1, name: 'Admin' },
                { id: 2, name: 'Back Officer' },
                { id: 3, name: 'Biller' },
                { id: 4, name: 'Approver' },
            ],
        };
    }

}
</script>