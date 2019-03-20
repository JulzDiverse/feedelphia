<template>
  <v-container
    grid-list-xl
  >
    <v-layout wrap>
      <v-flex xs12>
        <slot />
      </v-flex>

      <feed-card
        v-for="(article, i) in paginatedArticles"
        :key="article.title"
        :size="layout[i]"
        :value="article"
      />
    </v-layout>

    <v-layout align-center>
      <v-flex xs3>
        <base-btn
          v-if="page !== 1"
          class="ml-0"
          title="Previous page"
          square
          @click="page--"
        >
          <v-icon>mdi-chevron-left</v-icon>
        </base-btn>
      </v-flex>

      <v-flex
        xs6
        text-xs-center
        subheading
      >
        PAGE {{ page }} OF {{ pages }}
      </v-flex>

      <v-flex
        xs3
        text-xs-right
      >
        <base-btn
          v-if="pages > 1 && page < pages"
          class="mr-0"
          title="Next page"
          square
          @click="page++"
        >
          <v-icon>mdi-chevron-right</v-icon>
        </base-btn>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
  // Utilities
  import {
    mapState
  } from 'vuex'
  import * as axios from 'axios';

  export default {
    name: 'Feed',

    components: {
      FeedCard: () => import('@/components/FeedCard')
    },

    data: () => ({
      layout: [2, 2, 1, 2, 2, 3, 3, 3, 3, 3, 3],
      page: 1,
      contents: []
    }),

    methods: {
      downloadContents(){
        const url = "http://localhost:8081/photos"
       axios.get(url)
       .then(function (response) {
         console.log(response);
         var contents = response.data;
         contents.forEach(function(obj){
          console.log(obj.data);
          const fs = require('browserify-fs');
          const reader = new FileReader();
          reader.readAsDataURL(obj.data)
          fs.writeFile(obj.hero, reader.result, function(err) {
              if(err) {
                console.log(err);
              } else {
                console.log("The file was saved!");
              }
          });
        })
       })
       .catch(function (error) {
         console.log(error);
       })
      }
    },


    computed: {
      ...mapState(['articles']),

      pages () {
        return Math.ceil(this.articles.length / 11)
      },
      paginatedArticles () {
        const start = (this.page - 1) * 11
        const stop = this.page * 11

        return this.articles.slice(start, stop)
      }
    },

    beforeMount() {
       this.downloadContents();
    },

    watch: {
      page () {
        window.scrollTo(0, 0)
      }
    }
  }
</script>
