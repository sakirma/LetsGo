<template>
  <v-app>
    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'App',
  mounted() {
    this.axios.get('http://localhost:9000').then(e => {
      console.log(e)
    })
  },
})
</script>
