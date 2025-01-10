<template>
        
        <v-container class="mt-8">
            <bank-table/>
        </v-container>
</template>

<script>
import BankTable from '../components/admin/bank/BankTable.vue';

export default {
    components:{
        BankTable
    }
}
</script>