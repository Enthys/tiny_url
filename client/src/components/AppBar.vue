<template>
    <v-app-bar density="compact">
        <v-app-bar-title>
            <router-link to="/">tinyURL</router-link>
        </v-app-bar-title>
        <template v-slot:append>
            <router-link to="/">
                <v-btn border="outlined">Home</v-btn>
            </router-link>
            <router-link to="/create">
                <v-btn border="outlined">Create</v-btn>
            </router-link>
        </template>
    </v-app-bar>
</template>