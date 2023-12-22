<script setup>
import UserAvatar from "@/components/UserAvatar.vue";
</script>

<template>
  <UserAvatar
    username="Admin"
    :avatar="'https://inkythuatso.com/uploads/thumbnails/800/2022/03/anh-dai-dien-zalo-51-29-15-13-22.jpg'"
  >
    <slot />
  </UserAvatar>
</template>
