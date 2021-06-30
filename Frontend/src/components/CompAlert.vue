<template>
  <div @click="closeAlert" class="wrap-modal">
    <div class="modal-custom">
      <div class="modal-custom-content">{{ alert }}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: "CompAlert",
  props: {
    alert: String,
  },
  methods: {
    closeAlert() {
      this.$emit("close-alert");
    },
  },
};
</script>

<style>
.wrap-modal {
  z-index: 9999999;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgba(0, 0, 0, 0.3);
}
.modal-custom {
  width: 400px;
  height: 100px;
  background-color: white;
  border-radius: 10px;
  padding: 0 20px;
}
.modal-custom-content {
  padding: 20px 0;
}
</style>
