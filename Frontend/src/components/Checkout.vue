<template>
  <div class="wrap-cart-info">
    <p class="hide">Hide</p>
    <div>
      <!-- <div class="wrap-promo-code">
        <p>Nhập mã giảm giá</p>
        <input type="text" v-model="promoCode" />
        <button @click="applyPromoCode">Áp dụng</button>
        <div>
          <small>{{ alertInput }}</small>
        </div>
      </div> -->
      <div class="wrap-checkout">
        <div class="wrap-price-subtotal">
          <span>Trước thuế</span>
          <span>{{ formatPrice(subTotal) }}</span>
        </div>
        <div class="wrap-price-sale" v-if="percentSale">
          <span>Giảm giá {{ percentSale }}%</span>
          <span>-{{ formatPrice(priceSale) }}</span>
        </div>
        <div class="wrap-price-tax">
          <span>Thuế (VAT: 10%)</span>
          <span>{{ formatPrice(tax) }} </span>
        </div>
        <div class="wrap-price-total">
          <span>Tổng tiền</span>
          <span>{{ formatPrice(totalPrice) }}</span>
        </div>
        <button class="btn-checkout">THANH TOÁN</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Checkout",
  props: {
    products: Object,
  },
  data() {
    return {
      vat: 10,
      promoCodes: [
        { code: "vietnamvodich", percent: 50 },
        { code: "abcxyz", percent: 30 },
      ],
      percentSale: 0,
      promoCode: "",
      alertInput: "",
    };
  },
  computed: {
    subTotal() {
      // return this.products.reduce(
      //   (sumPrice, product) => sumPrice + product.price * product.num,
      //   0
      // );
      return this.$store.getters.getTotal;
    },
    priceSale() {
      return (this.subTotal * this.percentSale) / 100;
    },
    tax() {
      return (this.subTotal * this.vat) / 100;
    },
    totalPrice() {
      return this.subTotal + this.tax;
    },
  },
  watch: {
    promoCode(newState) {
      if (newState === "") {
        this.alertInput = "";
      }
    },
  },
  methods: {
    formatPrice(price) {
      return new Intl.NumberFormat("vi-VN", {
        style: "currency",
        currency: "VND",
      }).format(price);
    },
    applyPromoCode() {
      let objPromoCode = this.promoCodes.find(
        (element) => element.code.toLowerCase() === this.promoCode.toLowerCase()
      );
      this.percentSale = objPromoCode ? objPromoCode.percent : 0;
      if (this.percentSale) {
        this.alertInput = `Đơn hàng được giảm giá ${this.percentSale}%`;
      } else {
        this.alertInput = "Mã giảm giá không tồn tại";
      }
      if (!this.promoCode) {
        this.alertInput = "Mã không hợp lệ";
      }
    },
  },
};
</script>

<style scoped>
.wrap-cart-info {
  flex: 40%;
  padding-left: 50px;
  box-sizing: border-box;
  /* background-color: #f4f4f4; */
  position: sticky;
  top: 0;
}
.wrap-cart-info > div {
  /* background-color: #f4f4f4; */
  position: sticky;
  top: 0;
}
.hide {
  visibility: hidden;
}
.wrap-promo-code {
  border-bottom: 1px solid rgb(233, 233, 233);
  padding-bottom: 30px;
}
.wrap-promo-code p {
  padding-top: 15px;
}
.wrap-promo-code input {
  width: inherit;
  border-left: 1px solid rgb(145, 145, 145);
}

.wrap-checkout {
  padding-top: 20px;
}
.wrap-price-tax,
.wrap-price-subtotal,
.wrap-price-total,
.wrap-price-sale {
  display: flex;
  justify-content: space-between;
  padding: 15px 0;
}
.btn-checkout {
  width: 100%;
  padding: 15px 0;
  cursor: pointer;
  outline: none;
  border: none;
  background-color: #fcb800;
  color: white;
  font-weight: 700;
  font-size: 20px;
}
small {
  padding-top: 15px;
}
@media only screen and (max-width: 992px) {
  .wrap-cart-info {
    flex: 100%;
    padding-left: 0;
  }
}
</style>
