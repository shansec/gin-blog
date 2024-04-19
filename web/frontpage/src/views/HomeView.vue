<template>
  <v-app app>
    <Header></Header>
    <v-main class="grey lighten-3">
      <v-container>
        <v-row>
          <v-col cols="3">
            <NavView></NavView>
          </v-col>

          <v-col cols="8">
            <v-sheet  min-height="80vh" rounded="lg">
              <router-view :key="$route.path"></router-view>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
    <Footer></Footer>
  </v-app>
</template>

<script>
import Footer from '@/components/FooterView'
import Header from '@/components/HeaderView'
import NavView from '@/components/NavView'

export default {
  name: 'HomeView',
  components: {
    Footer,
    Header,
    NavView,
  },
}
</script>
