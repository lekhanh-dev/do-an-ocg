<template>
  <header>
    <div class="container">
      <button class="menu" @click="openNav">
        <i class="fas fa-bars"></i>
      </button>
      <div class="logo">
        <router-link :to="{ name: 'Home' }"
          ><span class="text-primary"> TECH </span
          ><span class="text-secondary">SHOP</span></router-link
        >
      </div>
      <nav :class="{ 'open-menu': openMenu }">
        <button class="close" @click="closeNav">
          <i class="fas fa-times"></i>
        </button>
        <ul>
          <li class="nav-lv-1">
            <span>Danh mục sản phẩm</span>
            <div class="nav-lv-2">
              <div class="wrap-nav-lv-2">
                <ul
                  class="wrap-categoryParent"
                  v-for="categoryParent in categories"
                  :key="categoryParent.id"
                >
                  <span class="categoryParent">{{ categoryParent.name }}</span>

                  <li
                    v-for="categoryChild in categoryParent.list_child"
                    :key="categoryChild.id"
                  >
                    <router-link
                      :to="{
                        name: 'Category',
                        params: {
                          categoryName: stringToSlug(categoryParent.name),
                          categoryId: categoryParent.id,
                          categoryChildName: stringToSlug(categoryChild.name),
                          categoryChildId: categoryChild.id,
                        },
                      }"
                      >{{ categoryChild.name }}</router-link
                    >
                  </li>
                </ul>
              </div>
            </div>
          </li>
          <!-- <li class="nav-lv-1">
            <router-link
              :to="{
                name: 'News',
              }"
              >Tin công nghệ</router-link
            >
          </li> -->
          <li class="nav-lv-1">
            <router-link
              :to="{
                name: 'About',
              }"
              >Liên hệ</router-link
            >
          </li>
        </ul>
      </nav>
      <div class="group-icon">
        <!-- <div class="search">
          <i class="fas fa-search"></i>
        </div> -->
        <div class="cart">
          <span class="sum-product">{{ sumProduct }}</span>
          <router-link :to="{ name: 'Cart' }">
            <i class="fas fa-shopping-cart"></i>
          </router-link>
        </div>
        <div class="avatar">
          <!-- <i class="fas fa-user-alt"></i> -->
          <!-- <div>
            <p>Đăng nhập</p>
            <p>Đăng ký</p>
          </div> -->
        </div>
      </div>
    </div>
  </header>
  <div class="cache"></div>
</template>

<script>
import axios from "axios";
export default {
  name: "HeaderPage",
  props: {},
  data() {
    return {
      listPhone: [
        ["Apple", "Samsung", "Xiaomi"],
        ["OPPO", "Vsmart", "Vivo"],
        ["OnePlus", "Nubia", "Vetur"],
        ["Điện thoại phổ thông"],
      ],
      openMenu: false,
      domainName: "http://localhost:3000",
      categories: [],
    };
  },
  computed: {
    sumProduct() {
      return this.$store.getters.getSumProduct;
    },
  },
  methods: {
    openNav() {
      this.openMenu = true;
    },
    closeNav() {
      this.openMenu = false;
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
  },
  created() {
    axios.get(`${this.domainName}/api/category`).then((response) => {
      this.categories = response.data;
    });
  },
};
</script>

<style scoped lang="scss">
.cache {
  padding-top: 70px;
}
.cart {
  position: relative;
  .sum-product {
    display: inline-block;
    position: absolute;
    top: -3px;
    right: 0px;
    background: rgb(44, 44, 252);
    border-radius: 50%;
    width: 15px;
    color: white;
    height: 15px;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 12px;
  }
}
i {
  color: black;
}
header {
  z-index: 300;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  background-color: #fcb800;
  .container {
    display: flex;
    align-items: center;
    .menu {
      outline: none;
      background-color: transparent;
      border: none;
      margin-right: 5px;
      margin-top: 1px;
      display: none;
      i {
        font-size: 20px;
        transition: all 0.3s ease-in-out;
        &:hover {
          color: white;
        }
      }
    }
    .logo {
      padding-right: 30px;
      cursor: pointer;
      a {
        text-decoration: none;
      }
      span {
        font-size: 25px;
        font-weight: 700;
      }
    }
    nav {
      .close {
        display: none;
      }
      ul {
        display: flex;
        margin: 0;
        li {
          color: black;
          cursor: pointer;
          padding: 20px 0;
          margin-right: 15px;
          span {
            &:hover {
              color: white;
            }
          }
          a {
            &:hover {
              color: white;
            }
          }
          &:nth-child(1):hover .nav-lv-2 {
            max-height: inherit;
            visibility: visible;
            top: 100%;
            opacity: 1;
          }
          a {
            text-decoration: none;
            color: black;
          }
        }
        .nav-lv-2 {
          background: white;
          position: absolute;
          left: 0;
          right: 0;
          top: 90%;
          padding: 0 20px;
          box-shadow: 0px 1px 2px #c5c5c5;
          max-height: 0;
          visibility: hidden;
          opacity: 0;
          transition: all 0.2s ease-out;
          .wrap-nav-lv-2 {
            display: flex;
            width: 1160px;
            margin: 0 auto;
            padding: 20px 0;
            ul {
              padding-right: 50px;
              display: block;
              span {
                color: black;
                &:hover {
                  color: black;
                  cursor: auto;
                }
              }
              li {
                padding-bottom: 5px;
                padding-top: 0;
                transition: all 0.3s ease-in-out;
                margin-right: 0;
                a {
                  &:hover {
                    color: #fcb800 !important;
                  }
                }
              }
            }
          }
        }
      }
    }
    .group-icon {
      flex: 1 1 auto;
      text-align: right;
      display: flex;
      align-items: center;
      justify-content: flex-end;
      > div:not(.avatar) {
        margin-left: 20px;
        i {
          font-size: 30px;
          transition: all 0.3s ease-in-out;
          cursor: pointer;
          &:hover {
            color: white;
          }
        }
      }
      .avatar {
        display: flex;
        align-items: center;
        i {
          font-size: 20px;
          margin-left: 20px;
        }
        div {
          margin-left: 5px;
          p {
            margin: 0;
            font-size: 12px;
            text-align: left;
            font-weight: 700;
            transition: all 0.3s ease-in-out;
            cursor: pointer;
            &:hover {
              color: white;
            }
          }
        }
      }
    }
  }
}
.categoryParent {
  font-weight: 700;
  font-size: 18px;
  border-bottom: 1px solid rgb(92, 92, 92);
  display: inline-block;
  margin-bottom: 15px;
}
.wrap-categoryParent {
  width: 200px;
}
@media (max-width: 768px) {
  header {
    .container {
      padding: 20px;
      .menu {
        display: flex;
        padding: 0;
        cursor: pointer;
      }
      .logo {
        span {
          font-size: 22px;
        }
      }
      nav {
        position: absolute;
        width: 100vw;
        height: 100vh;
        top: 0;
        left: -120%;
        bottom: 0;
        background-color: white;
        transition: all 0.5s ease-in;
        overflow-y: scroll;
        > ul {
          padding-left: 20px;
          flex-direction: column;
          margin-top: 20px;
          li {
            margin-right: 0;
          }
          .nav-lv-2 {
            position: static;
            padding: 0 20px;
            box-shadow: 0px 0px 0px transparent;
            max-height: inherit;
            visibility: visible;
            opacity: 1;
            .wrap-nav-lv-2 {
              display: block;
              padding-bottom: 0;
              width: auto;
            }
          }
        }
        .close {
          display: block;
          position: absolute;
          top: 20px;
          right: 20px;
          outline: none;
          font-size: 20px;
          border: none;
          cursor: pointer;
          background-color: transparent;
          transition: all 0.5s ease-in;
        }
      }
      .open-menu {
        display: block;
        left: 0;
      }
    }
  }
}
@media (max-width: 576px) {
  header {
    .container {
      padding: 10px;
      .logo {
        padding-right: 0;
        display: none;
        span {
          font-size: 14px;
        }
      }
    }
  }
}
</style>
