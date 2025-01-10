<template>

<!-- Navigation Drawer for Mobile -->
<v-navigation-drawer
      v-model="drawer"
      app
      temporary
      left
    >
      <v-list>
        <v-list-item
          v-for="(link, index) in navLinks"
          :key="index"
          :link="true"
          :to="link.path"
        >
          <v-list-item-content>
            <v-list-item-title>
              {{ link.title }}
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
</template>

<script>

</script>