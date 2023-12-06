<template>
    <v-card class="short-url-card">
        <v-card-text>
            <p>Short: 
                <a :href="shortUrl.getShortUrl()" target="_blank">{{ shortUrl.getShortUrl() }}</a>
            </p>
            <p>Leads to: 
                <a :href="shortUrl.url" target="_blank">{{ shortUrl.url }}</a>
            </p>
        </v-card-text>
    </v-card>
</template>

<script lang="ts" setup>
import { ShortUrl } from '../models/ShortUrl';

defineProps({
    shortUrl: {
        type: ShortUrl,
        required: true,
    },
})
</script>

<style scoped>
.short-url-card {
    margin-bottom: 1rem;
}
</style>