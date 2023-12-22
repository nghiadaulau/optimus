import { defineStore } from "pinia";
import { sent } from "../apis/request";

export const useContractStore = defineStore("contract", {
  state: () => ({}),
  actions: {
    listTodoList: async () => {
      const rep = await sent(
        "/query?channelid=mychannel&chaincodeid=basic&function=GetAllTodoItems",
        "GET"
      );
      return rep.data.data;
    },
    detailTodo: async (ID) => {
      const rep = await sent(
        `/query?channelid=mychannel&chaincodeid=basic&function=ReadTodoItem&args=${ID}`,
        "GET"
      );
      return rep;
    },
    createTodoList: async (Data) => {
      try {
        const response = await sent(
          "/invoke",
          "POST",
          Data,
          "'Content-Type': 'application/json'"
        );
        return response;
      } catch (error) {
        console.error("Lỗi khi thêm mới:", error);
        throw error;
      }
    },
    searchTodoList: async (Data) => {
      try {
        const response = await sent(
          "/search",
          "POST",
          Data,
          "'Content-Type': 'application/json'"
        );
        return response;
      } catch (error) {
        console.error("Lỗi khi thêm mới:", error);
        throw error;
      }
    },
    updateTodoList: async (Data) => {
      try {
        const response = await sent(
          "/invoke",
          "PUT",
          Data,
          "'Content-Type': 'application/json'"
        );
        return response;
      } catch (error) {
        console.error("Lỗi khi thêm mới:", error);
        throw error;
      }
    },
    deleteTodoList: async (Data) => {
      try {
        const response = await sent(
          "/delete",
          "DELETE",
          Data,
          "'Content-Type': 'application/json'"
        );
        return response;
      } catch (error) {
        console.error("Lỗi khi thêm mới:", error);
        throw error;
      }
    },
  },
});
