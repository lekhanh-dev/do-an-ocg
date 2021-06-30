<template>
  <section class="products">
    <div class="container">
      <h2 class="h2 upp align-center">Sản phẩm nổi bật nhất</h2>

      <div class="row">
        <transition-group name="fade">
          <div
            class="col-sm-6 col-md-3 product"
            v-for="product in products"
            :key="product.ID"
          >
            <div class="body">
              <a href="#favorites" class="favorites" data-favorite="inactive"
                ><i class="far fa-heart"></i
              ></a>
              <router-link
                :to="{
                  name: 'ProductDetail',
                  params: {
                    productName: stringToSlug(product.Name),
                    productId: product.ID,
                  },
                }"
              >
                <img
                  :src="domainName + '/' + product.ProductMedia[0].Uri"
                  alt="Image"
              /></router-link>
              <div class="content align-center">
                <p class="price">{{ formatPrice(product.Price) }}</p>
                <h2 class="h3">{{ product.Name }}</h2>
                <div class="btn-group">
                  <button class="btn btn-link">
                    <router-link
                      :to="{
                        name: 'ProductDetail',
                        params: {
                          productName: stringToSlug(product.Name),
                          productId: product.ID,
                        },
                      }"
                    >
                      <i class="fas fa-info"></i> Chi tiết</router-link
                    >
                  </button>
                  <button class="btn btn-primary btn-sm rounded">
                    <i class="ion-bag"></i> Thêm vào giỏ
                  </button>
                </div>
              </div>
            </div>
          </div>
        </transition-group>
      </div>

      <div
        v-show="products.length < sumProduct"
        class="align-right align-center-xs"
      >
        <hr class="offset-sm" />
        <h5 class="upp" @click="getMoreProducts"><span>Xem thêm</span></h5>
      </div>
    </div>
  </section>
</template>

<script>
import axios from "axios";

export default {
  name: "ProductList",
  props: {},
  data() {
    return {
      products: [],
      domainName: "http://localhost:3000",
      limit: 4,
      offset: 0,
      sumProduct: 0,
    };
  },
  methods: {
    stringToSlug(str) {
      str = str.toLowerCase();
      str = str.normalize("NFD").replace(/[\u0300-\u036f]/g, "");
      str = str.replace(/[đĐ]/g, "d");
      str = str.replace(/([^0-9a-z-\s])/g, "");
      str = str.replace(/(\s+)/g, "-");
      str = str.replace(/-+/g, "-");
      str = str.replace(/^-+|-+$/g, "");
      return str;
    },
    formatPrice(price) {
      return new Intl.NumberFormat("vi-VN", {
        style: "currency",
        currency: "VND",
      }).format(price);
    },
    getMoreProducts() {
      this.offset += this.limit;
      axios
        .get(
          `${this.domainName}/api/product/query?limit=${this.limit}&offset=${this.offset}`
        )
        .then((response) => {
          this.products = this.products.concat(response.data);
        });
    },
  },
  created() {
    axios
      .get(
        `${this.domainName}/api/product/query?limit=${this.limit}&offset=${this.offset}`
      )
      .then((response) => {
        this.products = response.data;
      });
    axios.get(`${this.domainName}/api/product/countAll`).then((response) => {
      this.sumProduct = response.data;
    });
  },
};
</script>

<style scoped lang="scss">
.container {
  padding: 30px;
}
h2.h2 {
  margin-left: -10px;
}
h5.upp {
  text-align: center;
  cursor: pointer;
}
h5.upp > span {
  display: inline-block;
  padding: 10px 30px;
  border-radius: 30px;
  color: white;
  background-color: #fcb800;
  transition: all 0.3s ease-in-out;
  &:hover {
    opacity: 0.7;
  }
}
hr.offset-sm {
  border: none;
  background: transparent;
  margin: 0;
  padding: 0;
  width: 100%;
  height: 15px;
}
.btn-group {
  display: flex;
  justify-content: center;
}
.products .product .btn {
  text-transform: none;
}
.btn.rounded {
  border-radius: 50px !important;
}
.btn.rounded:hover {
  background-color: #fcb800;
  border: 1px solid #fcb800;
}
.h3 {
  font-size: 17px;
  height: 40px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
.h3:hover {
  color: #00a1b4;
}
.btn-sm,
.btn-group-sm > .btn {
  padding: 5px 10px;
  font-size: 12px;
  line-height: 1.5;
  border-radius: 3px;
}
.btn-primary {
  color: #fff;
  background-color: rgba(30, 30, 30, 1);
  border-color: rgba(30, 30, 30, 1);
}
.btn-link {
  font-weight: normal;
  color: rgba(30, 30, 30, 1);
  border-radius: 0;
}
.products {
  display: block;
  width: 100%;
  height: auto;
  box-sizing: border-box;
  padding-top: 0px;
}

.products .product {
  box-sizing: border-box;
  margin-bottom: 5px;
  padding: 6px;
}

.products .product .body {
  box-sizing: border-box;
  padding: 15px 0 0px;
  background-color: #fff;
  border: 1px solid #fff;

  box-shadow: 0px 4px 2px rgba(31, 31, 31, 0.02);
  -webkit-box-shadow: 0px 4px 2px rgba(31, 31, 31, 0.02);

  -webkit-transition: all 0.4s;
  -moz-transition: all 0.4s;
  -ms-transition: all 0.4s;
  -o-transition: all 0.4s;
  transition: all 0.4s;

  -webkit-transform: translateY(0px);
  transform: translateY(0px);
}
.products .product:hover .body {
  -webkit-transform: translateY(-4px);
  transform: translateY(-4px);

  box-shadow: 0px 4px 10px rgba(31, 31, 31, 0.15);
  -webkit-box-shadow: 0px 4px 10px rgba(31, 31, 31, 0.15);
}

.products .product img {
  display: block;
  width: 90%;
  height: auto;
  padding: 20px;
  margin: 0 auto;
}

.products .product .content {
  display: block;
  width: 85%;
  height: auto;
  min-height: 145px;
  margin: 0 auto;
}

.products .product .content h1 {
  margin: 0;
}

.products .product p.price {
  display: inline-block;
  font-size: 17px;
  // line-height: 28px;
  color: #00a1b4;
  padding: 3px 0px;
  margin: 0;
}
.products .product p.sale {
  display: inline-block;
  font-size: 24px;
  line-height: 28px;
  color: #00a1b4;
}

.products .product p.through {
  text-decoration: line-through;
  font-size: 16px;
  line-height: 28px;
  color: #d2441c;
}

.products .product label {
  display: block;
  font-size: 14px;
  font-weight: 100;
  line-height: 28px;
}

.products .product .btn {
  text-transform: none;
}

.products .product .btn.btn-link {
  padding-left: 0px;
  padding-right: 0px;
  text-align: left;
}
.products .product .btn.btn-link:hover {
  text-decoration: none;
  color: #fcb800;
}

/* Favorites */
.products .product a.favorites {
  position: absolute;
  top: 55px;
  left: 0;
  right: 0;
  margin-left: auto;
  margin-right: auto;
  width: 50px;
  height: 50px;
  border-radius: 50px;
  background-color: rgba(0, 168, 180, 0.8);
  text-align: center;

  opacity: 0;

  -webkit-transition: all 0.3s;
  -moz-transition: all 0.3s;
  -ms-transition: all 0.3s;
  -o-transition: all 0.3s;
  transition: all 0.3s;

  -webkit-transform: scale(0, 0);
  -moz-transform: scale(0, 0);
  -ms-transform: scale(0, 0);
  -o-transform: scale(0, 0);
  transform: scale(0, 0);
}
.products .product a.favorites[data-favorite="active"] {
  opacity: 1;

  -webkit-transform: scale(1, 1);
  -moz-transform: scale(1, 1);
  -ms-transform: scale(1, 1);
  -o-transform: scale(1, 1);
  transform: scale(1, 1);
}

.products .product:hover a.favorites {
  opacity: 1;

  -webkit-transform: scale(1, 1);
  -moz-transform: scale(1, 1);
  -ms-transform: scale(1, 1);
  -o-transform: scale(1, 1);
  transform: scale(1, 1);
}

.products .product a.favorites > i {
  font-size: 40px;
  line-height: 60px;
  color: #fff;
  font-weight: 300;
}
</style>
