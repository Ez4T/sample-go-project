<template>
  <div>
    <main class="cover bg-cover min-h-screen">
      <nav class="flex justify-between items-center px-12 py-4">
        <div
          class="flex items-center justify-center bg-gray-200 rounded-full w-10 h-10"
        >
          <img class="w-8 h-8" src="@/assets/logo.png" alt="" />
        </div>
        <div>
          <button
            class="bg-gray-200 p-2 rounded-lg hover:bg-gray-100 transition text-black font-medium"
            type="button"
            @click="() => $router.push('/logout')"
          >
            Logout
          </button>
        </div>
      </nav>
      <section class="flex flex-col space-y-4 border-b-2 pb-6 mx-12">
        <div class="text-center">
          <h2 class="gallery">Upload Your Images</h2>
        </div>
        <div class="text-center">
          <input
            type="file"
            accept="image/*"
            @change="handleFiles($event)"
            multiple
          />
        </div>
        <div class="text-center">
          <button
            @click="uploadImage"
            class="bg-gray-200 hover:bg-gray-100 transition py-2 px-4 rounded-lg"
            type="button"
          >
            Submit
          </button>
        </div>
      </section>
      <div class="flex flex-col py-4">
        <div class="text-center">
          <h2 class="gallery">Gallery</h2>
        </div>
        <div
          class="grid grid-cols-3 justify-items-center space-x-4 space-y-4 image-box px-12 py-4"
        >
          <ImageComponentVue
            v-for="(item, index) in orderImagesList"
            :key="index"
            :img="displayImage(item.Img_path)"
          />
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { fetchGet, postUpload } from "../lib/api";
import ImageComponentVue from "../components/ImageComponent.vue";
const token = window.localStorage.getItem("token");
const username = window.localStorage.getItem("username");
const images = ref([]);
const orderImagesList = ref([]);
const formData = new FormData();
formData.append("username", username);
onMounted(async () => {
  fetchImages();
});

async function fetchImages() {
  const url = process.env.VUE_APP_URL + "api/v1/image-list";
  const response = await fetchGet(url, token);
  const results = response.data;
  if (response.status === 200) {
    images.value = results;
    orderImagesList.value = images.value.Images.slice().reverse();
  }
}

function displayImage(imgPath) {
  return process.env.VUE_APP_URL + imgPath;
}

function handleFiles(event) {
  let fileList = event.target.files;
  for (let file of fileList) {
    formData.append("file", file);
  }
}

async function uploadImage() {
  const url = process.env.VUE_APP_URL + "api/v1/multi-upload";
  const response = await postUpload(url, formData, token);
  if (response.status === 200) {
    fetchImages();
  }
}
</script>

<style scoped>
.cover {
  background-image: url("https://img.freepik.com/free-photo/white-cloud-blue-sky_74190-7709.jpg?w=1380&t=st=1659862288~exp=1659862888~hmac=b68f9be642e97fd7bbf6cf84a6651700ae63b4c0f0ab2b8b3407c33677dc4aa5");
  background-repeat: no-repeat;
  background-position: 100%;
}

.gallery {
  font-weight: 600;
  font-size: 30px;
  line-height: 50px;
  /* or 167% */

  text-align: center;

  color: #000000;
}
</style>
