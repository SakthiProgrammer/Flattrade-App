<!-- <template>
    <v-container class="d-flex justify-center align-center fill-height">
        <app-userlogin-form />
        <app-adminlogin-form /> </v-container>

</template>

<script>

import AdminLoginForm from "../components/login/AdminLoginForm.vue";
import UserLogin from "../components/login/components/LoginForm.vue"

export default {
    data() {
        return {

        }
    },
    components: {
        "app-userlogin-form": UserLogin,
        "app-adminlogin-form": AdminLoginForm

    }
}

</script>  -->

<template>
    <div class="my-auto fill-height">
    <v-container>
        <v-row class="fill-height align-center justify-center text-center">
            <!-- Left Column -->
            <v-col cols="12" md="6" sm="12" class="d-flex flex-column align-center justify-center">
                <h1>Welcome to Flattrade</h1>
                <div>
                <v-btn @click="selectRole('admin')" class="ma-2 my-5" outlined color="indigo" rounded>
                    Admin
                </v-btn>
                <v-btn @click="selectRole('user')" class="ma-2 my-5" outlined color="indigo" rounded>
                    User
                </v-btn>
            </div>
                {{ selectedRole }}
            </v-col>

            <v-col cols="12" md="6" sm="12" class="d-flex align-center justify-center">
                <component :is="selectedRole === 'admin' ? 'AdminLogin' : 'UserLogin'"></component>
            </v-col>
        </v-row>
    </v-container>
</div>

</template>

<script>
import AdminLogin from '../components/login/AdminLoginForm.vue';
import UserLogin from '../components/login/UserLoginForm.vue';

export default {
    components: {
        AdminLogin,
        UserLogin,
    },
    data() {
        return {
            selectedRole: 'user', // To track whether Admin or User is selected
        };
    },
    methods: {
        selectRole(role) {
            this.selectedRole = role;
        },
    },
    
};
</script>
