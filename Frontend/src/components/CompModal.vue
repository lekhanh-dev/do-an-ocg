<template>
  <div class="wrap-modal">
    <div class="modal-custom">
      <div class="modal-content-custom">Bạn có muốn xóa sản phẩm này ?</div>
      <div class="modal-group-action">
        <button class="btn-yes" @click="confirmYes">Đồng ý</button>
        <button class="btn-no" @click="confirmNo">Hủy bỏ</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "CompModal",
  methods: {
    confirmYes() {
      this.$emit("confirm", true);
    },
    confirmNo() {
      this.$emit("confirm", false);
    },
  },
};
</script>

<style>
.wrap-modal {
  z-index: 2;
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
  width: 300px;
  height: 100px;
  background-color: white;
  border-radius: 10px;
  padding: 0 20px;
}
.btn-yes {
  background-color: #fcb800;
  border: none;
  color: white;
}
.btn-no {
  border: none;
  background-color: tomato;
  color: white;
}
.modal-content-custom {
  padding: 20px 0;
}
.modal-group-action {
  width: 100%;
  display: flex;
  justify-content: space-between;
}
.modal-group-action button {
  width: 40%;
}
</style>
