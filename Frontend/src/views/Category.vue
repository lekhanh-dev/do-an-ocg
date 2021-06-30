<template>
  <div class="category">
    <div class="wrap-banner">
      <img :src="getLinkImage('banner_phone_2.jpg')" alt="img" />
    </div>
    <div class="container main">
      <div class="wrap-option">
        <div class="item-category">
          <div class="title">
            <h5>{{ categoryParent.Name }}</h5>
          </div>
          <div class="body">
            <h6 v-for="category in categoriesChild" :key="category.ID">
              <router-link
                :class="{
                  'link-active': category.ID == $route.params.categoryChildId,
                }"
                :to="{
                  name: 'Category',
                  params: {
                    categoryName: stringToSlug(categoryParent.Name),
                    categoryId: categoryParent.ID,
                    categoryChildName: stringToSlug(category.Name),
                    categoryChildId: category.ID,
                  },
                }"
                >{{ category.Name }}</router-link
              >
            </h6>
          </div>
        </div>
        <div class="item-category">
          <div class="title">
            <h5>Giá</h5>
          </div>
          <div class="body">
            <div class="price price-min">
              <span>Từ</span
              ><input
                type="number"
                :value="priceMin"
                @change="setPrice($event, 'min')"
                @blur="setPrice($event, 'min')"
              /><span>đồng</span>
            </div>
            <div class="price price-max">
              <span>Đến</span
              ><input
                type="number"
                :value="priceMax"
                @change="setPrice($event, 'max')"
                @blur="setPrice($event, 'max')"
              /><span>đồng</span>
            </div>
          </div>
        </div>
      </div>
      <div class="wrap-list-product">
        <transition name="fade">
          <div v-show="isOpenBoxNull" class="none-product">
            Không có sản phẩm nào thỏa mãn
          </div>
        </transition>
        <!-- <transition name="fade">
          <div v-show="loading" class="loading"><span>Loading...</span></div>
        </transition> -->
        <div class="header" v-if="!isOpenBoxNull">
          <select @change="handleSortProduct" :value="statusSort">
            <option value="">--Sắp xếp sản phẩm--</option>
            <option value="lowToHigh">Giá sản phảm từ thấp đến cao</option>
            <option value="highToLow">Giá sản phẩm từ cao đến thấp</option>
          </select>
          <span>Có {{ products.length }} sản phẩm</span>
        </div>
        <!-- <transition name="fade"> -->
        <div class="list-product">
          <!-- <transition-group name="fade"> -->
          <ProductCard
            v-for="product in productList"
            :key="product.ID"
            :product="product"
          ></ProductCard>
          <!-- </transition-group> -->

          <div
            class="align-right align-center-xs"
            v-if="products.length < sumProduct"
          >
            <h5 class="upp" @click="getMoreProducts">
              <span>Xem thêm</span>
            </h5>
          </div>
        </div>
        <!-- </transition> -->
      </div>
    </div>
  </div>
</template>

<script>
import ProductCard from "@/components/ProductCard.vue";
import axios from "axios";
export default {
  name: "Category",
  components: {
    ProductCard,
  },
  data() {
    return {
      categoriesChild: [],
      categoryParent: [],
      domainName: "http://localhost:3000",
      limit: 4,
      offset: 0,
      manufacturers: [],
      products: [],
      statusSort: "",
      isOpenBoxNull: false,
      priceMin: 0,
      priceMax: 100000000,
      loading: false,
      sumProduct: 0,
    };
  },
  computed: {
    productList() {
      switch (this.statusSort) {
        case "lowToHigh":
          return this.products
            .slice()
            .sort((product1, product2) => product1.Price - product2.Price);
        case "highToLow":
          return this.products
            .slice()
            .sort((product1, product2) => product2.Price - product1.Price);
        default:
          return [...this.products];
      }
    },
  },
  methods: {
    setPrice(event, type) {
      this.loading = true;
      let value = event.target.value;
      if (value === "") {
        value = 0;
      }
      const categoryChildId = this.$route.params.categoryChildId;
      if (type === "min") {
        this.priceMin = value;
      }
      if (type === "max") {
        this.priceMax = value;
      }
      axios
        .get(
          `${this.domainName}/api/product/query/v2?limit=4&offset=0&categoryId=${categoryChildId}&priceMin=${this.priceMin}&priceMax=${this.priceMax}`
        )
        .then((response) => {
          this.products = response.data;
          if (this.products.length === 0) {
            this.isOpenBoxNull = true;
          } else {
            this.isOpenBoxNull = false;
          }
          this.loading = false;
        });
    },
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
    getLinkImage(imageName) {
      return require(`@/assets/img/${imageName}`);
    },
    getMoreProducts() {
      this.offset += this.limit;
      const categoryChildId = this.$route.params.categoryChildId;
      this.loading = true;
      axios
        .get(
          `${this.domainName}/api/product/query/v2?limit=${this.limit}&offset=${this.offset}&categoryId=${categoryChildId}&priceMin=${this.priceMin}&priceMax=${this.priceMax}`
        )
        .then((response) => {
          this.products = this.products.concat(response.data);
          this.loading = false;
        });
    },
    handleSortProduct(event) {
      this.statusSort = event.target.value;
    },
  },
  created() {
    const categoryId = this.$route.params.categoryId;
    const categoryChildId = this.$route.params.categoryChildId;
    this.loading = true;
    axios
      .get(`${this.domainName}/api/category/${categoryId}`)
      .then((response) => {
        this.categoryParent = response.data[response.data.length - 1];
        this.categoriesChild = response.data.slice(0, response.data.length - 1);
      });
    axios.get(`${this.domainName}/api/manufacturer`).then((response) => {
      this.manufacturers = response.data;
    });
    axios
      .get(
        `${this.domainName}/api/product/query/v2?limit=${this.limit}&offset=${this.offset}&categoryId=${categoryChildId}&priceMin=${this.priceMin}&priceMax=${this.priceMax}`
      )
      .then((response) => {
        this.products = response.data;
        this.loading = false;
        if (this.products.length === 0) {
          this.isOpenBoxNull = true;
        } else {
          this.isOpenBoxNull = false;
        }
      });
    axios
      .get(
        `${this.domainName}/api/product/count?categoryId=${categoryChildId}&priceMin=${this.priceMin}&priceMax=${this.priceMax}`
      )
      .then((response) => {
        this.sumProduct = response.data;
      });
  },
  beforeRouteUpdate(to, from, next) {
    const categoryId = to.params.categoryId;
    const categoryChildId = to.params.categoryChildId;
    this.limit = 4;
    this.offset = 0;
    this.loading = true;
    this.statusSort = "";
    this.priceMin = 0;
    this.priceMax = 100000000;
    axios
      .get(`${this.domainName}/api/category/${categoryId}`)
      .then((response) => {
        this.categoryParent = response.data[response.data.length - 1];
        this.categoriesChild = response.data.slice(0, response.data.length - 1);
      });
    axios
      .get(
        `${this.domainName}/api/product/query/v2?limit=${this.limit}&offset=${this.offset}&categoryId=${categoryChildId}&priceMin=${this.priceMin}&priceMax=${this.priceMax}`
      )
      .then((response) => {
        this.products = response.data;
        this.loading = false;
        if (this.products.length === 0) {
          this.isOpenBoxNull = true;
        } else {
          this.isOpenBoxNull = false;
        }
        next();
      });
    axios
      .get(
        `${this.domainName}/api/product/count?categoryId=${categoryChildId}&priceMin=${this.priceMin}&priceMax=${this.priceMax}`
      )
      .then((response) => {
        this.sumProduct = response.data;
      });
  },
};
</script>

<style scoped lang="scss">
.loading {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: rgba(0, 0, 0, 0.1);
  position: absolute;
  top: 40px;
  left: 0;
  right: 0;
  z-index: 3;
}
.link-active {
  color: #fcb800 !important;
  cursor: auto;
  &:hover {
    border-bottom: none !important;
  }
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
.align-center-xs {
  flex: 100%;
  text-align: center;
  padding: 30px 0;
}
.none-product {
  width: 100%;
  height: 80vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: white;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  z-index: 2;
}
.wrap-banner {
  position: relative;
  top: -6px;
  width: 100%;
  // height: 300px;
  padding-bottom: 50px;
  img {
    object-fit: cover;
    width: 100%;
    height: auto;
  }
}
.main {
  display: flex;
  flex-wrap: wrap;
  min-height: 80vh;
  .wrap-option {
    flex: 20%;
    .item-category {
      padding-bottom: 50px;
      .title {
        padding-bottom: 25px;
        h5 {
          font-weight: 700;
        }
      }
      .body {
        h6 {
          margin-bottom: 20px;
          a {
            text-decoration: none;
            color: black;
            &:hover {
              border-bottom: 1px solid black;
            }
          }
        }
        .price {
          padding-bottom: 15px;
          span {
            display: inline-block;
            width: 40px;
          }
          input {
            margin-right: 5px;
            width: 120px;
            padding-left: 10px;
            height: 25px;
            outline: none;
            border: 1px solid rgb(211, 211, 211);
          }
        }
      }
    }
  }
  .wrap-list-product {
    position: relative;
    flex: 80%;
    min-height: 100vh;
    .header {
      flex: 100%;
      display: flex;
      justify-content: space-between;
      align-items: flex-end;
      padding: 6px;
      select {
        height: 35px;
        border-radius: 5px;
        padding: 5px;
        outline: none;
        border: 1px solid #bababa;
        border-radius: 0;
      }
    }
    .list-product {
      display: flex;
      flex-wrap: wrap;
      min-height: 400px;
    }
  }
}
</style>
