<template>
  <v-main>
    <AppBar />
    <v-container>
      <v-row class="mb-6">
        <v-col offset="3" cols="6">
          <router-view />
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script lang="ts" setup>
  import AppBar from '@/components/AppBar.vue'
  //
</script>
