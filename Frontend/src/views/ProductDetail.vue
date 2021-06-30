<template>
  <div class="wrap-page">
    <div class="container">
      <div class="main">
        <div class="wrap-slide-image">
          <div class="wrap-image">
            <transition name="fade">
              <img
                :src="
                  domainName + '/' + product?.ProductMedia[imageSelected].Uri
                "
                alt="Image"
              />
            </transition>
          </div>
          <div class="list-image">
            <img
              v-for="(image, index) in product?.ProductMedia"
              :class="{ active: index === imageSelected }"
              :src="domainName + '/' + image.Uri"
              alt="image"
              :key="image.ID"
              @click="changeImageSelected(index)"
            />
          </div>
        </div>
        <div class="wrap-product">
          <div class="product-title">{{ product?.Name }}</div>
          <div class="product-price">{{ formatPrice(product?.Price) }}</div>
          <div class="product-manufacturer">
            Hãng sản xuất: <b>{{ product?.Manufacturer.Name }}</b>
          </div>
          <div class="product-country">
            Quốc gia: <b>{{ product?.Country.Name }}</b>
          </div>
          <div class="product-description">
            <h5>Mô tả</h5>
            <p>{{ product?.Description }}</p>
          </div>
          <div class="wrap-action">
            <div class="counter">
              <button @click="updateNum('down')">-</button>
              <input
                type="number"
                v-model.number="numProduct"
                @blur="handleChangeNum($event)"
              />
              <button @click="updateNum('up')">+</button>
            </div>
            <button class="btn-add" @click="addToCart">Thêm giỏ hàng</button>
          </div>
        </div>
        <div class="wrap-properties">
          <h5>Cấu hình, thuộc tính</h5>
          <div v-for="property in product?.ProductProperty" :key="property.ID">
            <span>{{ property.Key }}</span>
            <span>{{ property.Value }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
  <CompAlert
    v-show="isOpenAlert"
    alert="Sản phẩm đã được thêm vào giỏ hàng"
    @close-alert="closeAlert"
  />
</template>

<script>
import axios from "axios";
import CompAlert from "@/components/CompAlert.vue";
export default {
  name: "ProductDetail",
  components: {
    CompAlert,
  },
  data() {
    return {
      product: null,
      numProduct: 1,
      imageSelected: 0,
      domainName: "http://localhost:3000",
      isOpenAlert: false,
    };
  },
  methods: {
    formatPrice(price) {
      return new Intl.NumberFormat("vi-VN", {
        style: "currency",
        currency: "VND",
      }).format(price);
    },
    updateNum(type) {
      if (type === "up") {
        this.numProduct++;
      } else {
        if (this.numProduct > 1) {
          this.numProduct--;
        }
      }
    },
    handleChangeNum(event) {
      if (event.target.value < 1) {
        this.numProduct = 1;
      } else {
        this.numProduct = parseInt(event.target.value);
      }
    },
    changeImageSelected(index) {
      this.imageSelected = index;
    },
    addToCart() {
      this.$store.commit("addToCart", {
        product: this.product,
        num: this.numProduct,
      });
      this.isOpenAlert = true;
    },
    closeAlert() {
      this.isOpenAlert = false;
    },
  },
  created() {
    const productId = this.$route.params.productId;
    axios
      .get(`${this.domainName}/api/product/getone/${productId}`)
      .then((response) => {
        this.product = response.data;
      });
  },
};
</script>

<style lang="scss" scoped>
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
input[type="number"] {
  -moz-appearance: textfield;
}
.active {
  border: 1px solid #fcb800;
}
.main {
  margin-top: 20px;
  padding-top: 20px;
  margin-bottom: 20px;
  box-sizing: border-box;
  display: flex;
  flex-wrap: wrap;
  background-color: white;
  .wrap-slide-image {
    flex: 40%;
    .wrap-image {
      width: 100%;
      padding-bottom: 20px;
      img {
        width: 100%;
        height: auto;
      }
    }
    .list-image {
      width: 50px;
      display: flex;
      img {
        width: 100%;
        height: auto;
        margin-right: 10px;
        cursor: pointer;
      }
    }
  }
  .wrap-product {
    flex: 60%;
    box-sizing: border-box;
    padding-left: 20px;
    .product-title {
      font-size: 35px;
      font-weight: 400;
    }
    .product-price {
      font-size: 30px;
      color: tomato;
    }
    .product-description {
      margin-top: 20px;
      padding-bottom: 30px;
    }
    .wrap-action {
      .counter {
        padding-bottom: 30px;
        input {
          outline: none;
          width: 70px;
          height: 30px;
          border: 1px solid;
          border-right: none;
          border-left: none;
          text-align: center;
          // position: relative;
          border-color: #fcb800;
          // top: 1px;
        }
        button {
          height: 30px;
          width: 30px;
          background-color: #fcb800;
          color: white;
        }
      }
      button {
        outline: none;
        border: none;
        width: 300px;
        height: 50px;
        background-color: #fcb800;
        color: white;
        transition: all 0.4s ease-in-out;
        &:hover {
          opacity: 0.8;
        }
      }
    }
  }
  .wrap-properties {
    margin-top: 20px;
    width: 100%;
    h5 {
      padding: 0 10px;
    }
    div {
      padding: 10px;
      &:nth-child(2n + 1) {
        background-color: rgb(228, 228, 228);
      }
      span {
        display: inline-block;
        min-width: 50%;
      }
    }
  }
}
</style>
