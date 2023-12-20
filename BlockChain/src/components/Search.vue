<script setup>
import { computed, ref } from "vue";

import {
  mdiArrowDownDropCircleOutline,
  mdiArrowUpDropCircleOutline,
  mdiMagnify,
} from "@mdi/js";
import CardBox from "@/components/CardBox.vue";
import FormField from "@/components/FormField.vue";
import FormControl from "@/components/FormControl.vue";
import GymButton from "./BaseButton.vue";
import BaseIcon from "./BaseIcon.vue";

const props = defineProps({
  opts: {
    type: Object,
    default: () => ({}),
  },
  opts2: {
    type: Object,
    default: () => ({}),
  },
  showSearchSvip: {
    type: Boolean,
    default: false,
  },
  placeholder: {
    type: String,
    default: "@nhập thông tin cần tìm kiếm",
  },
});

const emit = defineEmits(["search-submit"]);

const selectedValues = ref({
  search: "",
});
// const showSearch = () => {
//   alert(JSON.stringify(filteredSelectedValues.value));
// };
const showSearch = () => {
  const eventData = filteredSelectedValues.value;
  emit("search-submit", eventData);
};
const filteredSelectedValues = computed(() => {
  const filtered = {};
  for (const key in selectedValues.value) {
    if (selectedValues.value[key].id !== null) {
      filtered[key] = selectedValues.value[key];
    }
  }
  return filtered;
});
const showTable = ref(true);
if (props.showSearchSvip === true) {
  showTable.value = false;
}

const toggle = () => {
  showTable.value = !showTable.value;
};
</script>
<template>
  <CardBox form @submit.prevent="">
    <div v-if="props.showSearchSvip">
      <button @click="toggle">
        Tìm kiếm nâng cao
        <BaseIcon v-if="!showTable" :path="mdiArrowUpDropCircleOutline" />
        <BaseIcon v-else :path="mdiArrowDownDropCircleOutline" />
      </button>
      <FormField v-if="!showTable">
        <FormControl
          v-model="selectedValues.search"
          icon=""
          :placeholder="placeholder"
        />
      </FormField>
    </div>
    <table v-if="showTable">
      <tbody>
        <tr
          v-for="(row, rowIndex) in opts"
          :key="rowIndex"
          style="background-color: transparent"
        >
          <td
            v-for="(col, colIndex) in row.col"
            :key="colIndex"
            class="table-cell-no-border"
          >
            <!-- <template v-if="row.col.length === 2"> -->
            <FormField
              :label="
                col.type !== 'radio' && col.type !== 'checkbox' ? col.label : ''
              "
            >
              <FormControl
                v-model="selectedValues[col.index]"
                :icon="col.icon"
                :placeholder="col.placehoder"
                :type="col.type"
                :options="col.options"
                :maxlength="col.maxlength"
              />
            </FormField>
          </td>
        </tr>
      </tbody>
    </table>
    <table v-if="showTable" style="width: 90%">
      <tbody>
        <tr
          v-for="(row, rowIndex) in opts2"
          :key="rowIndex"
          style="background-color: transparent"
        >
          <td
            v-for="(col, colIndex) in row.col"
            :key="colIndex"
            class="table-cell-no-border"
          >
            <FormField
              :label="
                col.type !== 'radio' && col.type !== 'checkbox' ? col.label : ''
              "
            >
              <FormControl
                v-model="selectedValues[col.index]"
                :icon="col.icon"
                :placeholder="col.placehoder"
                :type="col.type"
                :options="col.options"
                :maxlength="col.maxlength"
              />
            </FormField>
          </td>
        </tr>
      </tbody>
    </table>
    <!-- <div>{{ filteredSelectedValues }}</div> -->
    <div class="mt-5">
      <div v-if="showTable" style="margin-left: 1%">
        <GymButton
          :icon="mdiMagnify"
          type="submit"
          color="orange"
          label="Search"
          @click="showSearch"
        />
      </div>
      <div v-else>
        <GymButton
          :icon="mdiMagnify"
          type="submit"
          color="orange"
          label="Search"
          @click="showSearch"
        />
      </div>
    </div>
  </CardBox>
</template>

<style scoped>
.table-cell-no-border {
  border: none;
}

/* .transparent-bg {
  background-color: transparent;
} */
</style>
