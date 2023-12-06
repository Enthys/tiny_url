<template>
    <v-form @submit.prevent="create">
        <v-text-field
            v-model="url"
            label="URL"
            :rules="urlRules"
            placeholder="https://example.com"
            required
        ></v-text-field>
        <v-btn type="submit">Create</v-btn>
    </v-form>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { createShortUrl } from '../models/ShortUrl';

const emit = defineEmits(['created']);

const urlRules = ref([
    (v: string) => !!v || 'URL is required',
    (v: string) => {
        try {
            new URL(v);
            return true;
        } catch (e) {
            return 'URL is invalid';
        }
    },
]);

const url = ref('');

async function create() {
    const newUrl = await createShortUrl(url.value);
    url.value = '';
    emit('created', newUrl);
}
</script>