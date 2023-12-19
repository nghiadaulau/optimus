<script setup>
import { reactive } from "vue";
import { useUserStore } from "@/stores/user.js";
import { useMainStore } from "@/stores/main.js";
import { mdiAccount, mdiAsterisk } from "@mdi/js";
import SectionFullScreen from "@/components/SectionFullScreen.vue";
import CardBox from "@/components/CardBox.vue";
import FormCheckRadio from "@/components/FormCheckRadio.vue";
import FormField from "@/components/FormField.vue";
import FormControl from "@/components/FormControl.vue";
import BaseButton from "@/components/BaseButton.vue";
import BaseButtons from "@/components/BaseButtons.vue";
import LayoutGuest from "@/layouts/LayoutGuest.vue";
import router from "@/router/index.js";
import localS from "@/util/localS.js";
const form = reactive({
  login: "",
  pass: "",
  remember: false,
});

const userStore = useUserStore();
const submit = async () => {
  const response = await userStore.loginUser({
    username: form.login,
    password: form.pass,
  });
  if (response.message == "Đăng nhập thành công") {
    // store.setUser(response.user);
    // setItem("user", response.user);
    localS.setItem("user", JSON.stringify(response.user));
    localS.setItem("lastLogin", JSON.stringify(response.lastLogin));

    // Lấy dữ liệu từ localStorage
    // const userDataJSON = localS.getItem("user");

    // // Chuyển đổi chuỗi JSON thành đối tượng JavaScript
    // const userData = JSON.parse(userDataJSON);

    // // Truy cập thuộc tính "id" của đối tượng
    // if (userData && userData.id) {
    //   console.log(userData.id);
    // } else {
    //   console.log("Không tìm thấy thông tin người dùng hoặc thuộc tính id.");
    // }

    router.push("/dashboard");
  }
};
</script>

<template>
  <LayoutGuest>
    <SectionFullScreen v-slot="{ cardClass }" bg="purplePink">
      <CardBox :class="cardClass" is-form @submit.prevent="submit">
        <FormField label="Login" help="Please enter your login">
          <FormControl
            v-model="form.login"
            :icon="mdiAccount"
            name="login"
            autocomplete="username"
          />
        </FormField>

        <FormField label="Password" help="Please enter your password">
          <FormControl
            v-model="form.pass"
            :icon="mdiAsterisk"
            type="password"
            name="password"
            autocomplete="current-password"
          />
        </FormField>

        <FormCheckRadio
          v-model="form.remember"
          name="remember"
          label="Remember"
          :input-value="true"
        />

        <template #footer>
          <BaseButtons>
            <BaseButton type="submit" color="info" label="Login" />
            <BaseButton to="/dashboard" color="info" outline label="Back" />
          </BaseButtons>
        </template>
      </CardBox>
    </SectionFullScreen>
  </LayoutGuest>
</template>
