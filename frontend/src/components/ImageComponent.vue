<template>
  <div>
    <img
      class="w-[448px] max-h-[383px] object-cover"
      :src="props.img"
      alt="images"
    />
  </div>
</template>

<script setup>
import { defineProps } from "vue";
const props = defineProps({
  img: String,
});
</script>

<style></style>
