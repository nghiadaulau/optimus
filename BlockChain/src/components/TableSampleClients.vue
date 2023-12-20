<script setup>
import { computed, ref, onMounted } from "vue";
import { mdiEye, mdiTrashCan } from "@mdi/js";
import CardBoxModal from "@/components/CardBoxModal.vue";
import TableCheckboxCell from "@/components/TableCheckboxCell.vue";
import BaseLevel from "@/components/BaseLevel.vue";
import BaseButtons from "@/components/BaseButtons.vue";
import BaseButton from "@/components/BaseButton.vue";
import UserAvatar from "@/components/UserAvatar.vue";
import { useUserStore } from "@/stores/user.js";
import VueQRCode from "vue-qrcode";
defineProps({
  checkable: Boolean,
});

const items = ref([]);
const userStore = useUserStore();
onMounted(async () => {
  items.value = await userStore.listUser();
  // console.log(items.value);
});

const isModalActive = ref(false);

const isModalDangerActive = ref(false);

const perPage = ref(5);

const currentPage = ref(0);

const checkedRows = ref([]);

const itemsPaginated = computed(() =>
  items.value.slice(
    perPage.value * currentPage.value,
    perPage.value * (currentPage.value + 1)
  )
);

const numPages = computed(() => Math.ceil(items.value.length / perPage.value));

const currentPageHuman = computed(() => currentPage.value + 1);

const pagesList = computed(() => {
  const pagesList = [];

  for (let i = 0; i < numPages.value; i++) {
    pagesList.push(i);
  }

  return pagesList;
});

const remove = (arr, cb) => {
  const newArr = [];

  arr.forEach((item) => {
    if (!cb(item)) {
      newArr.push(item);
    }
  });

  return newArr;
};

const checked = (isChecked, client) => {
  if (isChecked) {
    checkedRows.value.push(client);
  } else {
    checkedRows.value = remove(
      checkedRows.value,
      (row) => row.id === client.id
    );
  }
};
const selectedUser = ref(null);

const showUserModal = (user) => {
  selectedUser.value = user;
  isModalActive.value = true;
};

const idUser = ref("");
const qrCode = ref("");

const showQrUser = (user) => {
  idUser.value = user.id;
  // console.log(idUser.value);
  isModalDangerActive.value = true;
  generateQrCodeForUser(idUser.value);
};
const generateQrCodeForUser = async (userId) => {
  const response = await userStore.userQr(userId);

  const jsonString = JSON.stringify(response);

  qrCode.value = jsonString;
};
</script>

<template>
  <CardBoxModal v-model="isModalActive" title="User Information">
    <div v-if="selectedUser">
      <p><strong>Name:</strong> {{ selectedUser.name }}</p>
      <p><strong>Email:</strong> {{ selectedUser.email }}</p>
      <p><strong>Username:</strong> {{ selectedUser.username }}</p>
      <p><strong>Department:</strong> {{ selectedUser.departments }}</p>
      <p><strong>Status:</strong> {{ selectedUser.status }}</p>
    </div>
  </CardBoxModal>

  <CardBoxModal v-model="isModalDangerActive" title="QR">
    <figure class="qrcode">
      <VueQRCode
        :value="qrCode"
        :type="'image/png'"
        :color="{ dark: '#FFFFFF', light: '#FFA500' }"
        tag="png"
        :options="{
          width: 200,
        }"
      ></VueQRCode>
      <img class="qrcode__image" src="../static/Glogo.png" alt="Silver" />
    </figure>
  </CardBoxModal>

  <div v-if="checkedRows.length" class="p-3 bg-gray-100/50 dark:bg-slate-800">
    <span
      v-for="checkedRow in checkedRows"
      :key="checkedRow.id"
      class="inline-block px-2 py-1 rounded-sm mr-2 text-sm bg-gray-100 dark:bg-slate-700"
    >
      {{ checkedRow.name }}
    </span>
  </div>

  <table>
    <thead>
      <tr>
        <th v-if="checkable" />
        <th>STT</th>
        <th>Avatar</th>
        <th>Họ và Tên</th>
        <th>Email</th>
        <th>Trạng Thái</th>
        <th />
      </tr>
    </thead>
    <tbody>
      <tr v-for="(client, index) in itemsPaginated" :key="client.id">
        <TableCheckboxCell
          v-if="checkable"
          @checked="checked($event, client)"
        />
        <td class="border-b-0 lg:w-6 before:hidden">
          {{ index + 1 }}
        </td>
        <td class="border-b-0 lg:w-6 before:hidden">
          <UserAvatar
            :username="client.name"
            class="w-24 h-24 mx-auto lg:w-6 lg:h-6"
          />
        </td>

        <td data-label="Name">
          {{ client.name }}
        </td>
        <td data-label="Email">
          {{ client.email }}
        </td>

        <td data-label="Status">
          {{ client.status }}
        </td>

        <td class="before:hidden lg:w-1 whitespace-nowrap">
          <BaseButtons type="justify-start lg:justify-end" no-wrap>
            <BaseButton
              color="info"
              :icon="mdiEye"
              small
              @click="showUserModal(client)"
            />
            <BaseButton
              color="danger"
              :icon="mdiTrashCan"
              small
              @click="showQrUser(client)"
            />
          </BaseButtons>
        </td>
      </tr>
    </tbody>
  </table>
  <div class="p-3 lg:px-6 border-t border-gray-100 dark:border-slate-800">
    <BaseLevel>
      <BaseButtons>
        <BaseButton
          v-for="page in pagesList"
          :key="page"
          :active="page === currentPage"
          :label="page + 1"
          :color="page === currentPage ? 'lightDark' : 'whiteDark'"
          small
          @click="currentPage = page"
        />
      </BaseButtons>
      <small>Page {{ currentPageHuman }} of {{ numPages }}</small>
    </BaseLevel>
  </div>
</template>
<style scoped>
.qrcode {
  display: inline-block;
  font-size: 0;
  margin-bottom: 0;
  position: relative;
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
