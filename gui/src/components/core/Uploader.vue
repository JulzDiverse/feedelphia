<template>
  <v-layout row justify-center>
    <v-dialog v-model="dialog" persistent max-width="600px">
      <template v-slot:activator="{ on }">
        <v-btn color="primary" dark v-on="on">Upload Photo</v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="headline">Upload Photo</span>
        </v-card-title>
        <v-card-text>
          <v-container grid-list-md>
            <v-layout wrap>
              <v-flex xs12 sm12 md12>
                <v-text-field label="Photo Title" v-model="title" required></v-text-field>
              </v-flex>
              <v-flex xs12 sm12 md12>
                <v-text-field label="Author" hint="name of the photographer" v-model="author"></v-text-field>
              </v-flex>
              <v-flex xs12 sm12 md12>
                <input
                class="ml-0 hidden-sm-and-down"
                type="file"
                ref="uploadFile"
                accept="image/*"
                capture="user"
                @change="setPhoto($event)"
                >
              </v-flex>
              </v-layout>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" flat @click="dialog = false">Cancel</v-btn>
          <v-btn color="blue darken-1" flat @click="upload()">Upload</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-layout>
</template>
<script>
  import * as axios from 'axios';

  const BASE_URL = 'http://localhost:8081';

  export default {
    data: () => ({
      dialog: false,
      author: "",
      title: "",
      photo: ""
    }),
    methods: {
       setPhoto(event){
          this.photo = event.target.files[0]
       },
       upload(){
         const url = "http://localhost:8081/photos";

         this.dialog = false;
         var formData = new FormData();

         formData.append("photo", this.photo)
         formData.append("title", this.title)
         formData.append("author", this.author)

         axios.post(url, formData)
         clearForm(this);
       }
    }
  }

function clearForm(that){
   that.author = "";
   that.title = "";
}
</script>
