<script setup>
import { computed, ref } from "vue";
import { mdiClose } from "@mdi/js";
import BaseButton from "@/components/BaseButton.vue";
import BaseButtons from "@/components/BaseButtons.vue";
import CardBox from "@/components/CardBox.vue";
import OverlayLayer from "@/components/OverlayLayer.vue";
import CardBoxComponentTitle from "@/components/CardBoxComponentTitle.vue";
import { useContractStore } from "@/stores/contract";

const branchStore = useContractStore();

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
  tmp: {
    type: Object,
    default: null,
  },
  hasCancel: Boolean,
  modelValue: {
    type: [String, Number, Boolean],
    default: null,
  },
});
const setBranchData = () => {
  return {
    channelid: "mychannel",
    chaincodeid: "basic",
    function: "DeleteTodoItem",
    args: [props.tmp.ID],
  };
};

const emit = defineEmits(["update:modelValue", "cancel", "confirm", "delete"]);

const value = computed({
  get: () => props.modelValue,
  set: (value) => emit("update:modelValue", value),
});

const confirmCancel = (mode) => {
  value.value = false;
  emit(mode);
};
const sleep = (milliseconds) => {
  return new Promise((resolve) => setTimeout(resolve, milliseconds));
};
const isSaveButtonClicked = ref(false);

const deleteItem = async () => {
  if (!isSaveButtonClicked.value) {
    isSaveButtonClicked.value = true;
    const jsonData = setBranchData();
    const rep = await branchStore.deleteTodoList(jsonData);
    if (rep.data.status == "success") {
      await sleep(3000);
      value.value = false;
      emit("delete");
    }
    isSaveButtonClicked.value = false;
  }
};

const cancel = () => confirmCancel("cancel");

window.addEventListener("keydown", (e) => {
  if (e.key === "Escape" && value.value) {
    cancel();
  }
});
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
      <strong v-if="props.tmp"> ID: {{ props.tmp.ID }}</strong>

      <div class="space-y-3">
        <table v-if="props.tmp">
          <tbody>
            <tr style="background-color: transparent">
              <td class="table-cell-no-border">
                Description: {{ props.tmp.Description }}
              </td>
              <td class="table-cell-no-border">
                Status: {{ props.tmp.Status }}
              </td>
            </tr>
            <tr style="background-color: transparent">
              <td class="table-cell-no-border">
                StartDate:
                {{ new Date(props.tmp.StartDate).toLocaleDateString("vi-VN") }}
              </td>
              <td class="table-cell-no-border">
                EndDate:
                {{ new Date(props.tmp.EndDate).toLocaleDateString("vi-VN") }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <template #footer>
        <BaseButtons>
          <BaseButton
            label="Delete"
            color="danger"
            outline
            @click="deleteItem"
          />
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
  width: 50%;
}
</style>
