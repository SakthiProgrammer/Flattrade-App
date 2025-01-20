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
    <div>
        <Header />
        <div class="my-auto fill-height">
            <v-container>
                <v-row class="fill-height align-center justify-center text-center">
                    <!-- Left Column -->
                    <v-col cols="12" md="6" sm="12" class="d-flex flex-column align-center justify-center">
                        <h1>Welcome to Flattrade</h1>
                        <div>
                            <v-btn @click="selectRole('admin')" class="ma-2 my-5" :outlined="isBtnSelected" color="cyan"
                                rounded :dark="!isBtnSelected">
                                Admin
                            </v-btn>
                            <v-btn @click="selectRole('client')" class="ma-2 my-5" color="cyan"
                                :outlined="!isBtnSelected" rounded :dark="isBtnSelected">
                                User
                            </v-btn>
                        </div>
                        <!-- {{ selectedRole }} -->
                    </v-col>

                    <v-col cols="12" md="6" sm="12" class="d-flex align-center justify-center">
                        <!-- <component
              :is="selectedRole === 'admin' ? 'AdminLogin' : 'UserLogin'"
            ></component> -->
                        <div v-show="selectedRole === 'admin'">
                            <admin-login />
                        </div>
                        <div v-show="selectedRole === 'client'">
                            <user-login />
                        </div>
                    </v-col>
                </v-row>
            </v-container>
        </div>
    </div>

</template>

<script>
import Header from '../components/header/header.vue';
import AdminLogin from '../components/login/AdminLoginForm.vue';
import UserLogin from '../components/login/UserLoginForm.vue';

export default {
    components: {
        AdminLogin,
        UserLogin,
        Header
    },
    data() {
        return {
            selectedRole: 'client', // To track whether Admin or User is selected
            isBtnSelected: true,
        };
    },
    methods: {
        selectRole(role) {
            this.selectedRole = role;
            this.isBtnSelected = !this.isBtnSelected;
        },
    },

};
</script>
