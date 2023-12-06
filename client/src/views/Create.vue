<template>
    <div>
        <h1>Create a short link</h1>
        <CreateShortLinkForm @created="appendToCollection"/>
        <ShortUrlTable :short-urls="shortUrls" />
    </div>
</template>

<script setup lang="ts">
import CreateShortLinkForm from '@/components/CreateShortLinkForm.vue'
import ShortUrlTable from '../components/ShortUrlTable.vue';
import { ref } from 'vue';
import { ShortUrl } from '../models/ShortUrl';

const shortUrls = ref<Map<number, ShortUrl>>(new Map());

function appendToCollection(newUrl: ShortUrl) {
    shortUrls.value.set(newUrl.id, newUrl);
}

</script>
