<template>
    <v-container>
        <div v-if="shortUrls.size !== 0">
            <ShortUrlCard v-for="url in shortUrls.values()" :key="url.id" :short-url="url"/>
        </div>
        <div v-else>
            <v-card>
                <v-card-text>
                    <h3>No short new URLs yet.</h3>
                </v-card-text>
            </v-card>
        </div>
    </v-container>
</template>

<script lang="ts" setup>
import ShortUrlCard from '@/components/ShortUrlCard.vue';
import { ShortUrl } from '@/models/ShortUrl';
import { PropType } from 'vue';

defineProps({
    shortUrls: {
        type: Map as PropType<Map<number,ShortUrl>>,
        required: true,
        default: () => [],
    }
})
</script>
