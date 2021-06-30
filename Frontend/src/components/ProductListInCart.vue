<template>
  <div class="wrap-cart-product">
    <div class="product-list">
      <p class="info">
        <span class="break">Giỏ hàng</span>
        <span>Có {{ totalNumProduct }} sản phẩm trong túi</span>
      </p>
      <div class="product-item" v-for="product in products" :key="product.ID">
        <div class="product-image">
          <img
            :src="domainName + '/' + product.ProductMedia[0].Uri"
            alt="image"
          />
        </div>
        <div class="product-detail">
          <div class="product-name">
            <a href="#"> {{ product.Name }}</a>
          </div>
          <div class="product-description">{{ product.Description }}</div>
          <div class="product-price">{{ formatPrice(product.Price) }}</div>
          <div class="product-counter-mobile">
            <button
              class="btn-decrease"
              @click="downNumProduct(--product.num, product.id)"
            >
              -
            </button>
            <input
              type="number"
              :value="product.num"
              @input="changeNumProduct($event, product.id)"
              min="1"
            />
            <button
              class="btn-increase"
              @click="upNumProduct(++product.num, product.id)"
            >
              +
            </button>
          </div>
        </div>
        <div class="product-counter">
          <button
            class="btn-decrease"
            @click="downNumProduct(product.num - 1, product.ID)"
          >
            -
          </button>
          <input
            type="number"
            :value="product.num"
            @input="changeNumProduct($event, product.id)"
            @blur="handleOnBlur($event, product.id)"
            min="1"
          />
          <button
            class="btn-increase"
            @click="upNumProduct(product.num + 1, product.ID)"
          >
            +
          </button>
        </div>
        <div class="product-action">
          <button class="btn-delete" @click="openModal(product.ID)">Xóa</button>
        </div>
      </div>
    </div>
  </div>
  <transition name="fade">
    <CompModal v-if="isOpenModal" @confirm="confirmDeleteProduct" />
  </transition>
</template>

<script>
import CompModal from "./CompModal.vue";
export default {
  name: "ProductListInCart",
  components: {
    CompModal,
  },
  props: {
    products: Object,
  },
  data() {
    return {
      numProduct: 0,
      isOpenModal: false,
      idProduct: 0,
      domainName: "http://localhost:3000",
    };
  },
  computed: {
    totalNumProduct() {
      return this.products.reduce((num, product) => num + product.num, 0);
    },
  },
  methods: {
    formatPrice(price) {
      return new Intl.NumberFormat("vi-VN", {
        style: "currency",
        currency: "VND",
      }).format(price);
    },
    // changeNumProduct(event, idProduct) {
    //   let num = event.target.value;
    //   if (num !== "") {
    //     this.$emit("change-num-product", parseInt(num), idProduct);
    //   } else {
    //   }
    // },
    upNumProduct(num, idProduct) {
      // this.$emit("up-num-product", parseInt(num), idProduct);
      this.$store.commit("updateNumProduct", { id: idProduct, num: num });
    },
    downNumProduct(num, idProduct) {
      if (num === 0) {
        this.$store.commit("updateNumProduct", { id: idProduct, num: 1 });
        // this.$emit("down-num-product", 1, idProduct);
      } else {
        this.$store.commit("updateNumProduct", { id: idProduct, num: num });
        // this.$emit("down-num-product", parseInt(num), idProduct);
      }
    },
    openModal(idProduct) {
      this.isOpenModal = true;
      this.idProduct = idProduct;
    },
    confirmDeleteProduct(isConfirm) {
      // console.log(this.idProduct);
      if (isConfirm) {
        // this.$emit("delete-product", this.idProduct);
        this.$store.commit("deleteProduct", this.idProduct);
      }
      this.isOpenModal = false;
    },
    handleOnBlur(event, idProduct) {
      let num = event.target.value;
      if (num === "") {
        this.$emit("down-num-product", 1, idProduct);
      }
    },
  },
};
</script>

<style scoped lang="scss">
input,
button {
  padding: 5px 10px;
  border: none;
  outline: none;
  text-align: center;
}
input {
  border-top: 1px solid rgb(145, 145, 145);
  border-bottom: 1px solid rgb(145, 145, 145);
}
button {
  background-color: #fcb800;
  color: white;
  border-top: 1px solid #fcb800;
  border-bottom: 1px solid #fcb800;
  cursor: pointer;
  transition: all 0.2s ease-in-out;
}
button:hover {
  opacity: 0.8;
}
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
input[type="number"] {
  -moz-appearance: textfield;
}
.wrap-cart-product {
  flex: 60%;
}
.product-list {
  margin: 10px 0;
  border-bottom: 1px solid rgb(233, 233, 233);
}
.product-item {
  padding: 30px;
  display: flex;
  border-top: 1px solid rgb(233, 233, 233);
}
.product-image {
  width: 70px;
  height: 100px;
}
img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}
.product-counter-mobile {
  display: none;
}
.product-detail {
  flex: 50%;
  padding-left: 15px;
  padding-top: 10px;
}
.product-name {
  font-weight: 600;
  padding-bottom: 5px;
}
a {
  text-decoration: none;
}
.product-description {
  font-size: 13px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
.product-price {
  padding: 10px 0;
  font-weight: 600;
  font-weight: 25px;
}
.btn-delete:hover {
  opacity: 0.4;
}
.product-counter {
  flex: 20%;
  display: flex;
  align-items: center;
  padding: 0 20px;
  input {
    border-top: 1px solid #fcb800;
    border-bottom: 1px solid #fcb800;
  }
}
input {
  width: 50px;
}

.product-action {
  flex: 10%;
  display: flex;
  align-items: center;
  justify-content: flex-end;
}
i {
  cursor: pointer;
}
.btn-delete {
  background: rgb(255, 66, 66);
  border-top: 1px solid rgb(255, 66, 66);
  border-bottom: 1px solid rgb(255, 66, 66);
}
.info {
  display: flex;
  justify-content: space-between;
  margin: 0;
  padding: 15px 0;
  position: sticky;
  top: 0;
}
@media only screen and (max-width: 992px) {
  .wrap-cart-product {
    flex: 100%;
  }
}
@media only screen and (max-width: 576px) {
  .product-counter input {
    width: 20px;
  }
  .product-item {
    padding: 10px;
  }
  .product-description,
  .break {
    display: none;
  }
  .product-counter {
    display: none;
  }
  .product-counter-mobile {
    display: block;
  }
  .product-counter-mobile input {
    width: 15px;
  }
}
</style>
