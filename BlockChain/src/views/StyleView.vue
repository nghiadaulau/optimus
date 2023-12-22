<script setup>
import { useRouter } from "vue-router";
import { useStyleStore } from "@/stores/style.js";
import { gradientBgPurplePink } from "@/colors.js";
import SectionMain from "@/components/SectionMain.vue";
import CardBox from "@/components/CardBox.vue";
import LayoutGuest from "@/layouts/LayoutGuest.vue";
import localS from "@/util/localS";
const styles = ["white", "basic", "custom"];

const styleStore = useStyleStore();

styleStore.setDarkMode(false);

const router = useRouter();

const click = (slug) => {
  styleStore.setStyle(slug);
  const userData = localS.getItem("user");
  if (userData) {
    // User data exists, you might want to redirect to another page
    // or perform any other action
    router.push("/dashboard");
  } else {
    router.push("/login");
  }
};
</script>

<template>
  <LayoutGuest>
    <div
      :class="gradientBgPurplePink"
      class="flex min-h-screen items-center justify-center"
    >
      <SectionMain>
        <h1
          class="text-4xl md:text-5xl text-center text-white font-bold mt-12 mb-3 lg:mt-0"
        >
          Choose a style&hellip;
        </h1>
        <h2 class="text-xl md:text-xl text-center text-white mb-12">
          Design By
          <code class="px-1.5 py-0.5 rounded bg-white bg-opacity-20"
            >Silver</code
          >
        </h2>
        <div
          class="grid gap-6 grid-cols-1 lg:grid-cols-3 px-6 max-w-6xl mx-auto"
        >
          <CardBox
            v-for="style in styles"
            :key="style"
            class="cursor-pointer bg-gray-50"
            is-hoverable
            @click="click(style)"
          >
            <div class="mb-3 md:mb-6">
              <img
                v-if="style !== 'custom'"
                :src="`https://static.justboil.me/templates/one/small/${style}-v3.png`"
                width="1280"
                height="720"
              />
              <img
                v-else
                :src="`https://cdn.sforum.vn/sforum/wp-content/uploads/2023/07/hinh-nen-ai-98.jpg`"
                width="1280"
                height="720"
              />
            </div>

            <h1 class="text-xl md:text-2xl font-black capitalize">
              {{ style }}
            </h1>
            <h2 class="text-lg md:text-xl">& Dark mode</h2>
          </CardBox>
        </div>
      </SectionMain>
    </div>
  </LayoutGuest>
</template>
