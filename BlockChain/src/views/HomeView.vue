<script setup>
import { ref } from "vue";
import {
  mdiPlusCircle,
  mdiFileDownload,
  mdiSquareEditOutline,
  mdiDelete,
} from "@mdi/js";
import LayoutAuthenticated from "@/layouts/LayoutAuthenticated.vue";
import TableTodoList from "@/components/TableTodoList.vue";
import FormLegend from "@/components/FormLegend.vue";
import CardBox from "@/components/CardBox.vue";
import Search from "@/components/Search.vue";
import { useContractStore } from "@/stores/contract";
const branchStore = useContractStore();
// const items = ref([]);
// const branchStore = useContractStore();
// onMounted(async () => {
//   items.value = await branchStore.listTodoList();
//   console.log(items.value);
// });
const priority = [
  { id: 1, label: "Cao" },
  { id: 2, label: "Trung bình" },
  { id: 3, label: "Thấp" },
];
const status = [
  { id: 1, label: "Done" },
  { id: 2, label: "Pending" },
];
const buttonTable = ref([
  {
    color: "info",
    icon: mdiPlusCircle,
    roundedFull: true,
    small: true,
    outline: true,
    label: "Thêm mới",
  },
  {
    color: "info",
    icon: mdiFileDownload,
    roundedFull: true,
    small: true,
    outline: true,
    label: "Export",
  },
]);
const buttonOperation = ref([
  {
    color: "info",
    icon: mdiSquareEditOutline,
    name: "edit",
    small: true,
  },
  {
    color: "danger",
    icon: mdiDelete,
    name: "delete",
    small: true,
  },
]);
const coll1s = [
  {
    icon: "",
    label: "Mã công việc",
    placehoder: "Nhập Mã công việc",
    index: "ID",
  },
  {
    icon: "",
    label: "Tên người sỡ hữu",
    placehoder: "Nhập Tên người sỡ hữu",
    index: "Owner",
  },
];
const coll2s = [
  {
    icon: "",
    label: "Mức độ ưu tiên",
    options: priority,
    type: "select",
    index: "Priority",
  },
  {
    icon: "",
    label: "Trạng thái",
    options: status,
    type: "select",
    index: "Status",
  },
];
const coll3s = [
  {
    icon: "",
    label: "Ngày bắt đầu",
    type: "date",
    index: "StartDate",
  },
  {
    icon: "",
    label: "Ngày kết thúc",
    type: "date",
    index: "EndDate",
  },
];
const opts = ref([
  {
    col: coll1s,
  },
  {
    col: coll2s,
  },
  {
    col: coll3s,
  },
]);
const setBranchData = (data) => {
  return {
    channelid: "mychannel",
    chaincodeid: "basic",
    status: data.Status,
  };
};
const search = ref();
const handleSearchSubmit = async (data) => {
  data.Status = data.Status ? data.Status.label : "";
  const jsonData = setBranchData(data);
  if (data.Status) {
    const rep = await branchStore.searchTodoList(jsonData);
    search.value = rep.data.data;
  }
};
</script>

<template>
  <LayoutAuthenticated>
    <div class="col-md-4 column">
      <Search
        :opts="opts"
        :show-search-svip="true"
        @search-submit="handleSearchSubmit"
      />
    </div>
    <br />
    <div class="col-md-6 column">
      <CardBox>
        <FormLegend label="Danh sách công việc">
          <TableTodoList
            checkable
            :button-table="buttonTable"
            :button-operation="buttonOperation"
            :search="search"
          ></TableTodoList></FormLegend
      ></CardBox>
    </div>
  </LayoutAuthenticated>
</template>
