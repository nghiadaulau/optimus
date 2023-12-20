<script setup>
import { computed, ref, onMounted } from "vue";
import LayoutAuthenticated from "@/layouts/LayoutAuthenticated.vue";
import VueQRCode from "vue-qrcode";
import BaseButton from "@/components/BaseButton.vue";
import { useUserStore } from "@/stores/user.js";
import { useMainStore } from "@/stores/main.js";

const qrCode = ref("Vui Lòng Đăng nhập để sử dụng chức năng này!!!");
const store = useMainStore();
const userStore = useUserStore();

const generateQrCodeForUser = async (userId) => {
  const response = await userStore.userQr(userId);
  const jsonString = JSON.stringify(response);
  const formattedJson = jsonString.replace(/"/g, " ").replace(/[{}[\]]/g, "");
  qrCode.value = formattedJson.replace(/,/g, "\n");
};
generateQrCodeForUser(store.$state.userid);
// const openLink = () => {
//   // window.location.href = this.qrValue;
//   alert(qrCode.value);
// };
</script>

<template>
  <LayoutAuthenticated>
    <div class="qrcode-container">
      <figure class="qrcode">
        <VueQRCode
          :value="qrCode"
          :type="'image/png'"
          :color="{ dark: '#FFFFFF', light: '#FFA500' }"
          tag="png"
          :options="{
            width: 500,
          }"
        ></VueQRCode>
        <img class="qrcode__image" src="../static/Glogo.png" alt="Silver" />
      </figure>
    </div>
    <!-- <div class="button">
      <BaseButton
        @click="openLink"
        target="_blank"
        :icon="mdiGithub"
        label="Infomation"
        color="contrast"
        rounded-full
        small
      />
    </div> -->
  </LayoutAuthenticated>
</template>

<style scoped>
.qrcode-container {
  display: flex;
  justify-content: center;
  align-items: center;
}
/* .button {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 10vh; 
} */
.qrcode {
  display: inline-block;
  font-size: 0;
  margin-bottom: 0;
  position: relative;
  text-align: center;

  white-space: pre-wrap;
}

.qrcode__image {
  background-color: #fff;
  border: 0.25rem solid #fff;
  border-radius: 0.25rem;
  box-shadow: 0 0.125rem 0.25rem rgba(0, 0, 0, 0.25);
  height: 20%;
  left: 50%;
  overflow: hidden;
  position: absolute;
  top: 50%;
  transform: translate(-50%, -50%);
  width: 20%;
}
</style>
