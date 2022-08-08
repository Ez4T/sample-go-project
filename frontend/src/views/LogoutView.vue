<template>
  <div></div>
</template>

<script setup>
import { useRouter } from "vue-router";
const router = useRouter();
window.localStorage.removeItem("username");
window.localStorage.removeItem("token");
router.push("/login");
</script>

<style></style>
