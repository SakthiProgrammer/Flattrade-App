<template>
    <v-container>
        <v-card class="pa-md-15 mx-lg-auto pa-sm-10" width="500px">
            <v-row>
                <h1 class="mx-auto mb-7">Sign In</h1>
            </v-row>
            <v-row>
                <!-- <v-select v-model="name" outlined item-text="name" item-value="name" label="User ID"></v-select> -->
                <v-text-field label="User Id" outlined></v-text-field>

            </v-row>
            <v-row>
                <v-text-field label="Password" outlined></v-text-field>
            </v-row>
            <v-row>
                <h1 class="body-2"> If Don't Have an Account
                    <router-link to="/register">
                        Create?
                    </router-link>
                </h1>
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
</script>