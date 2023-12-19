<script setup>
import { computed, reactive, watch, ref } from "vue";
import { mdiClose } from "@mdi/js";
import BaseButton from "@/components/BaseButton.vue";
import BaseButtons from "@/components/BaseButtons.vue";
import CardBox from "@/components/CardBox.vue";
import OverlayLayer from "@/components/OverlayLayer.vue";
import CardBoxComponentTitle from "@/components/CardBoxComponentTitle.vue";
import FormField from "./FormField.vue";
import FormControl from "./FormControl.vue";
import { useContractStore } from "@/stores/contract";

const branchStore = useContractStore();

const priority = [
  { id: 1, label: "Cao" },
  { id: 2, label: "Trung bình" },
  { id: 3, label: "Thấp" },
];
const status = [
  { id: 1, label: "Done" },
  { id: 2, label: "Pending" },
];
const props = defineProps({
  title: {
    type: String,
    required: true,
  },
  button: {
    type: String,
    default: "info",
  },
  buttonLabel: {
    type: String,
    default: "Done",
  },
  hasCancel: Boolean,
  modelValue: {
    type: [String, Number, Boolean],
    default: null,
  },
  tmp: {
    type: Object,
    default: null,
  },
});
const selectedValues = reactive({});
const setBranchData = (value) => {
  if (value.StartDate && value.EndDate) {
    const startDate = new Date(value.StartDate);
    const formattedStartDate = startDate.toISOString();
    const endDate = new Date(value.EndDate);
    const formattedEndDate = endDate.toISOString();
    return {
      channelid: "mychannel",
      chaincodeid: "basic",
      function: "UpdateTodoItem",
      args: {
        ID: value.ID,
        Owner: value.Owner,
        Description: value.Description,
        Status: value.Status.label,
        StartDate: formattedStartDate,
        EndDate: formattedEndDate,
        Priority: value.Priority.id,
      },
    };
  }
};

const emit = defineEmits(["update:modelValue", "cancel", "confirm"]);
const value = computed({
  get: () => props.modelValue,
  set: (value) => emit("update:modelValue", value),
});
const sleep = (milliseconds) => {
  return new Promise((resolve) => setTimeout(resolve, milliseconds));
};
const isSaveButtonClicked = ref(false);

const confirmCancel = async (mode) => {
  if (!isSaveButtonClicked.value) {
    isSaveButtonClicked.value = true;
    if (mode == "confirm") {
      const jsonData = setBranchData(selectedValues);
      const rep = await branchStore.updateTodoList(jsonData);
      if (rep.data.status == "success") {
        await sleep(3000);
        value.value = false;
        emit(mode);
      }
    } else {
      value.value = false;
      emit(mode);
    }
    isSaveButtonClicked.value = false;
  }
};

const confirm = () => confirmCancel("confirm");

const cancel = () => confirmCancel("cancel");

window.addEventListener("keydown", (e) => {
  if (e.key === "Escape" && value.value) {
    cancel();
  }
});
watch(
  () => props.tmp,
  async (newValue) => {
    const mapPriority = (priorityId) => {
      const priorityItem = priority.find((item) => item.id === priorityId);
      return priorityItem
        ? { id: priorityItem.id, label: priorityItem.label }
        : null;
    };
    const mapStatus = (statusId) => {
      const statusItem = status.find((item) => item.label == statusId);
      return statusItem ? { id: statusItem.id, label: statusItem.label } : null;
    };
    const fetchData = async () => {
      selectedValues.ID = newValue.ID || "";
      selectedValues.Status = mapStatus(newValue.Status);
      selectedValues.Owner = newValue.Owner || "";
      selectedValues.StartDate = (newValue.StartDate || "").split("T")[0] || "";
      selectedValues.EndDate = (newValue.EndDate || "").split("T")[0] || "";
      selectedValues.Description = newValue.Description || "";
      selectedValues.Priority = mapPriority(newValue.Priority);
    };
    await fetchData();
  }
);
</script>

<template>
  <OverlayLayer v-show="value" @overlay-click="cancel">
    <CardBox
      v-show="value"
      class="shadow-lg max-h-modal w-11/12 md:w-3/5 lg:w-2/5 xl:w-4/12 z-50"
      is-modal
    >
      <CardBoxComponentTitle :title="title">
        <BaseButton
          v-if="hasCancel"
          :icon="mdiClose"
          color="whiteDark"
          small
          rounded-full
          @click.prevent="cancel"
        />
      </CardBoxComponentTitle>
      <strong> ID: {{ selectedValues.ID }}</strong>
      <div class="space-y-3">
        <table>
          <tbody>
            <tr style="background-color: transparent">
              <td class="table-cell-no-border">
                <FormField label="Ngày bắt đầu">
                  <FormControl
                    v-model="selectedValues.StartDate"
                    type="date"
                    icon=""
                  />
                </FormField>
              </td>
              <td class="table-cell-no-border">
                <FormField label="Ngày kết thúc">
                  <FormControl
                    v-model="selectedValues.EndDate"
                    type="date"
                    icon=""
                  />
                </FormField>
              </td>
            </tr>
            <tr style="background-color: transparent">
              <td class="table-cell-no-border">
                <FormField label="Trạng thái">
                  <FormControl
                    v-model="selectedValues.Status"
                    type="select"
                    icon=""
                    :options="status"
                  />
                </FormField>
              </td>
              <td class="table-cell-no-border">
                <FormField label="Mức độ ưu tiên">
                  <FormControl
                    v-model="selectedValues.Priority"
                    type="select"
                    icon=""
                    :options="priority"
                  />
                </FormField>
              </td>
            </tr>
            <tr style="background-color: transparent">
              <td class="table-cell-no-border">
                <FormField label="Mô tả công việc">
                  <FormControl
                    v-model="selectedValues.Description"
                    type="textarea"
                    icon=""
                    placeholder="Vui lòng nhập mô tả"
                  />
                </FormField>
              </td>
              <td class="table-cell-no-border">
                <strong> Owner: {{ selectedValues.Owner }}</strong>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <template #footer>
        <BaseButtons>
          <BaseButton label="Update" :color="button" @click="confirm" />
          <BaseButton
            v-if="hasCancel"
            label="Cancel"
            :color="button"
            outline
            @click="cancel"
          />
        </BaseButtons>
      </template>
    </CardBox>
  </OverlayLayer>
</template>
<style scoped>
.table-cell-no-border {
  border: none;
  vertical-align: top;
}
</style>
