<template>
  <div class="container">
    <div class="cart">
      <div class="none-cart" v-if="!products.length">
        Giỏ hàng không có sản phẩm nảo
        <button @click="goHomePage">Quay lại mua hàng</button>
      </div>
      <ProductListInCart
        v-if="products.length"
        :products="products"
        @change-num-product="changeNumProduct"
        @up-num-product="upNumProduct"
        @down-num-product="downNumProduct"
        @delete-product="deleteProductById"
      />
      <Checkout :products="products" v-if="products.length" />
    </div>
  </div>
</template>

<script>
import Checkout from "@/components/Checkout.vue";
import ProductListInCart from "../components/ProductListInCart.vue";
export default {
  name: "Cart",
  components: {
    Checkout,
    ProductListInCart,
  },
  data() {
    return {};
  },
  computed: {
    products() {
      return this.$store.state.cart;
    },
  },
  methods: {
    changeNumProduct(num, idProduct) {
      this.products = this.products.map((product) => {
        if (product.id === idProduct) {
          product.num = num;
        }
        return product;
      });
    },
    upNumProduct(num, idProduct) {
      this.changeNumProduct(num, idProduct);
    },
    downNumProduct(num, idProduct) {
      this.changeNumProduct(num, idProduct);
    },
    deleteProductById(idProduct) {
      this.products = this.products.filter(
        (product) => product.id !== idProduct
      );
    },
    goHomePage() {
      this.$router.push({ name: "Home" });
    },
  },
  created() {
    this.productsClone = this.products.map((product) => ({ ...product }));
  },
};
</script>

<style scoped>
body {
  background-color: #f4f4f4;
}
.cart {
  display: flex;
  flex-wrap: wrap;
  position: relative;
  background-color: white;
  padding: 0 10px;
}

h1 {
  font-size: 20px;
  flex: 100%;
}
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
.none-cart {
  width: 100%;
  height: 80vh;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
.none-cart button {
  margin-top: 20px;
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
