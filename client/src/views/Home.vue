<template>
  <v-container>
    <v-row class="page-nav" align-content="center" align="center">
      <v-btn @click="previousPage">Previous</v-btn>
      <v-spacer></v-spacer>
      <v-btn @click="nextPage">Next</v-btn>
    </v-row>
    <ShortUrlTable :short-urls="shortUrls" />
  </v-container>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue';
import ShortUrlTable from '../components/ShortUrlTable.vue';
import { ShortUrl, getShortUrls } from '../models/ShortUrl';

const shortUrls = ref<Map<number, ShortUrl>>(new Map());
const page = ref(1);

getShortUrls(page.value).then((urls) => {
  shortUrls.value = urls;
});

function previousPage() {
  if (page.value > 1) {
    page.value--;
    getShortUrls(page.value).then((urls) => {
      shortUrls.value = urls;
    });
  }
}

function nextPage() {
  page.value++;
  getShortUrls(page.value).then((urls) => {
    shortUrls.value = urls;
  });
}
</script>

<style>
.page-nav {
  margin-bottom: 2rem;
}
</style>