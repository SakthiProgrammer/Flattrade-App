<template>
    <div>
        <Header/>
    <user-register />

</div>
</template>


<script>
import Register from '../components/register/Register.vue';
import Header from '../components/header/header.vue';

export default {
    data() {
        return {

        }
    },
    components: {
        "user-register": Register,
        Header,
    },

};
</script>
