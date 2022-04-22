<template>
  <v-card class="overflow-hidden">
    <v-app-bar
      absolute
      color="#6A76AB"
      dark
    >
      <v-toolbar-title>Do I have COVID-19?</v-toolbar-title>
      <v-menu
        bottom
        left
      >
        <v-list>
          <v-list-item
            v-for="(item, i) in items"
            :key="i"
          >
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>

      <template v-slot:extension>
        <v-tabs align-with-title>
            <router-link to="/">
          <v-tab>Home</v-tab></router-link>
           <router-link to="/about">
          <v-tab>Contact Us</v-tab>
          </router-link>
        </v-tabs>
      </template>
    </v-app-bar>
    <v-sheet
      id="scrolling-techniques-4"
      class="overflow-y-auto"
      max-height="600"
    >
      <v-container class="container"></v-container>
    </v-sheet>
  </v-card>
</template>

<script>
  export default {
    data: () => ({
      items: [
        { title: 'Click Me' },
        { title: 'Click Me' },
        { title: 'Click Me' },
        { title: 'Click Me 2' },
      ],
    }),
  }
</script>

<style scoped>
a {
  text-decoration: none;
}
 
.router-link-active {
  text-decoration: none;
}
.container{
    height:100px
}
</style>